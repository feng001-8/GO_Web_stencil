package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_web/logger"
	"net/http"
)

func Setup() *gin.Engine {
	engine := gin.New()
	engine.Use(logger.GinLogger(zap.L()), logger.GinRecovery(zap.L(), true))

	engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	return engine
}
