package router

import (
	"fmt"
	"ogsyoo/imageExport-api/src/common/conf"
	"ogsyoo/imageExport-api/src/controller"
	"ogsyoo/imageExport-api/src/router/middleware/header"

	"github.com/gin-gonic/gin"

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
		base.StaticFS("/ui", http.Dir(conf.UiDoc))
		sse := base.Group("/sse")
		{
			sse.GET("/event/:name", controller.Subcribe)
			sse.GET("/publish", controller.Publish)
		}
		project := base.Group("/project")
		{
			project.GET("/list", controller.GetListPage)
			project.POST("", controller.AddProject)
			project.DELETE("", controller.DeleteProject)
			project.GET("/images/:pid", controller.GetImageListPage)
		}
		image := base.Group("/image")
		{
			image.PUT("", controller.UpdateImageJob)
			image.DELETE("/:id", controller.DeleteImageJob)
			image.POST("/package", controller.PackageImage)
		}
		base.GET("/health", controller.Health)
	}
	return e
}
