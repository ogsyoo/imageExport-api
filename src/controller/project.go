package controller

import (
	"github.com/gin-gonic/gin"
	"ogsyoo/imageExport-api/src/dao"
	"ogsyoo/imageExport-api/src/service"
	"strconv"
)

//添加导出项目任务
func AddProject(c *gin.Context) {
	project := new(dao.Project)
	c.Bind(project)
	pvs := service.Project{}
	err := pvs.InsertPorject(project)
	if err != nil {
		c.JSON(200, gin.H{"success": 0, "errMsg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": 1, "data": "新增成功"})
}

//删除导出项目任务
func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil || int_id == 0 {
		c.JSON(200, gin.H{"success": 0, "errMsg": "get id error"})
		return
	}
	pvs := service.Project{}
	err = pvs.DelPorject(int_id)
	if err != nil {
		c.JSON(200, gin.H{"success": 0, "errMsg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": 1, "data": "删除成功"})
}

//更新导出项目任务
func UpdateProject(c *gin.Context) {
	var pro *dao.Project
	c.Bind(pro)
	pvs := service.Project{}
	err := pvs.UpdatePorject(pro)
	if err != nil {
		c.JSON(200, gin.H{"success": 0, "errMsg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": 1, "data": "更新成功"})
}

//获取列表任务
func GetListPage(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	pvs := service.Project{}
	ls, err := pvs.GetListPage(limit, offset)
	if err != nil {
		c.JSON(200, gin.H{"success": 0, "errMsg": err.Error()})
		return
	}
	c.JSON(200, ls)
}
