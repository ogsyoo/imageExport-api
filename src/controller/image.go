package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ogsyoo/imageExport-api/src/dao"
	"ogsyoo/imageExport-api/src/service"
)

//新增镜像任务---批量
func AddImageJob(c *gin.Context) {
	var listImage []*dao.ImageJob
	c.Bind(listImage)
	if len(listImage) == 0 {
		c.JSON(200, gin.H{"success": 0, "errMsg": "请完善要添加的数据"})
		return
	}
	isv := service.Image{}
	a, err := isv.InsertImageList(listImage)
	if len(listImage) == 0 {
		c.JSON(200, gin.H{"success": 0, "errMsg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": 1, "data": fmt.Sprint(`成功插入%d条数据`, a)})
}
