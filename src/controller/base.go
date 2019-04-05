package controller

import (
	"github.com/gin-gonic/gin"
	"ogsyoo/imageExport-api/src/common/tools"
	"time"
)

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"host":       c.Request.Host,
		"header":     c.Request.Header,
		"serverTime": time.Now(),
		"ip":         tools.RemoteIp(c.Request),
	})
}
