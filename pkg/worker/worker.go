package worker

import (
	"context"
	asynq2 "github.com/Tempest-Finance/console-strategies-common/pkg/asynq"
	"github.com/Tempest-Finance/console-strategies-common/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type Worker struct {
	server *asynq.Server
	mux    *asynq.ServeMux
}

func (w *Worker) Run(address string, mode string) error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		gin.SetMode(mode)
		engine := gin.New()
		engine.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/health"))
		engine.Use(gin.Recovery())
		engine.GET("/health", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, "OK")
		})
		return engine.Run(address)
	})

	eg.Go(func() error {
		err := w.server.Run(w.mux)
		if err != nil {
			return err
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func (w *Worker) RegisterHandler(taskType string, handler func(ctx context.Context, t *asynq.Task) error) {
	w.mux.HandleFunc(taskType, handler)
}

var worker *Worker

func Init(config Config) error {
	if worker != nil {
		return nil
	}

	redisConnOpt, err := asynq2.GetAsynqRedisConnectionOption(asynq2.Config(config))
	if err != nil {
		return err
	}

	server := asynq.NewServer(
		redisConnOpt,
		asynq.Config{
			Logger: logger.L().Sugar(),
			HealthCheckFunc: func(err error) {
				if err != nil {
					logger.L().Sugar().Warnf("healthcheck error: %s", err)
				}
			},
			Queues: map[string]int{
				QueueIDCritical: QueueWeightCritical,
				QueueIDDefault:  QueueWeightDefault,
				QueueIDLow:      QueueWeightLow,
			},
			StrictPriority: true,
		})
	mux := asynq.NewServeMux()
	worker = &Worker{
		server: server,
		mux:    mux,
	}

	return nil
}

func Instance() *Worker {
	return worker
}
