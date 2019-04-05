package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ogsyoo/imageExport-api/src/common/conf"
	"ogsyoo/imageExport-api/src/controller"
	"ogsyoo/imageExport-api/src/router/middleware/header"

	"net/http"
)

// Load loads the router
func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.Default()
	e.Use(gin.Recovery())
	e.Use(header.NoCache)
	e.Use(header.Options)
	e.Use(header.Secure)
	e.Use(middleware...)
	base := e.Group(fmt.Sprintf(`%s/`, conf.BaseInfo.Prefix))
	{
		project := base.Group("/project")
		{
			project.GET("/list", controller.GetListPage)
			project.POST("", controller.AddProject)
			project.DELETE("", controller.DeleteProject)

		}
		base.GET("/health", controller.Health)
	}
	return e
}
