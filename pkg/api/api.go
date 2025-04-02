package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	engine *gin.Engine
}

func Init(mode string, registerFn func(engine *gin.Engine)) *Api {
	engine := newEngine(mode)
	registerFn(engine)
	return &Api{engine: engine}
}

func (s *Api) Run(address string) error {
	return s.engine.Run(address)
}

func newEngine(mode string) *gin.Engine {
	gin.SetMode(mode)
	engine := gin.New()
	engine.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/health"))
	engine.Use(gin.Recovery())
	engine.Use(Logger)
	setCORS(engine)
	return engine
}

func setCORS(engine *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowMethods(http.MethodOptions)
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowHeaders("X-API-Key")
	corsConfig.AllowAllOrigins = true
	engine.Use(cors.New(corsConfig))
}
