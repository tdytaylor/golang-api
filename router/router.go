package router

import (
	"github.com/gin-gonic/gin"
	"goweb/handler/checks"
	"goweb/router/middleware"
	"net/http"
)

func Init(engine *gin.Engine) {
	engine.Use(gin.Logger())
	engine.Use(middleware.NoCache)
	engine.Use(middleware.Options)
	engine.Use(middleware.Secure)

	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The incorrect API route")
	})

	group := engine.Group("/check")
	group.GET("/health", checks.HealthCheck)
	group.GET("/disk", checks.DiskCheck)
	group.GET("/cpu", checks.CPUCheck)
	group.GET("/ram", checks.RAMCheck)
}
