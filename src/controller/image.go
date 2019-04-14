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
	_, err := isv.InsertImageList(listImage)
	if len(listImage) == 0 {
		c.JSON(200, gin.H{"success": 0, "errMsg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": 1, "data": "插入成功"})
}

//更新镜像信息
func UpdateImageJob(c *gin.Context) {
	image := new(dao.ImageJob)
	c.Bind(image)
	isv := service.Image{}
	_, err := isv.UpdateImageJob(image)
	if err != nil {
		c.JSON(200, gin.H{"success": 0, "errMsg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": 1, "data": "更新成功"})
}

//删除镜像信息
func DeleteImageJob(c *gin.Context) {
	id := c.Param("id")
	isv := service.Image{}
	_, err := isv.DeleteImageJob(id)
	if err != nil {
		c.JSON(200, gin.H{"success": 0, "errMsg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": 1, "data": "更新成功"})
}

//打包镜像
func PackageImage(c *gin.Context) {
	ids := []dao.ImageJob{}
	c.Bind(&ids)
	fmt.Println(ids)
	isv := service.Image{}
	err := isv.PackageImage(ids)
	if err != nil {
		c.JSON(200, gin.H{"success": 0, "errMsg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": 1, "data": "打包成功"})
}
