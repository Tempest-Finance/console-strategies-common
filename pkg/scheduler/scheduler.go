package scheduler

import (
	asynq2 "github.com/Tempest-Finance/console-strategies-common/pkg/asynq"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type TaskConfigGenerator func() []*asynq.PeriodicTaskConfig

type periodicTaskCfgProvider struct {
	generators []TaskConfigGenerator
}

func (p *periodicTaskCfgProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	var configs []*asynq.PeriodicTaskConfig
	for _, g := range p.generators {
		configs = append(configs, g()...)
	}
	return configs, nil
}

type Scheduler struct {
	taskCfgProvider *periodicTaskCfgProvider
	taskManager     *asynq.PeriodicTaskManager
}

func (m *Scheduler) Run(address string, mode string) error {
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
		return m.taskManager.Run()
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func (m *Scheduler) RegisterTaskConfigGenerator(generator TaskConfigGenerator) {
	m.taskCfgProvider.generators = append(m.taskCfgProvider.generators, generator)
}

var scheduler *Scheduler

func Init(config Config) error {
	if scheduler != nil {
		return nil
	}

	redisConnOpt, err := asynq2.GetAsynqRedisConnectionOption(asynq2.Config(config))
	if err != nil {
		return err
	}

	provider := &periodicTaskCfgProvider{}
	manager, err := asynq.NewPeriodicTaskManager(asynq.PeriodicTaskManagerOpts{
		RedisConnOpt:               redisConnOpt,
		PeriodicTaskConfigProvider: provider,
	})
	if err != nil {
		return err
	}

	scheduler = &Scheduler{
		taskCfgProvider: provider,
		taskManager:     manager,
	}
	return nil
}

func Instance() *Scheduler {
	return scheduler
}
