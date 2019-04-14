package dao

import "time"

type ImageJob struct {
	Id       string `xorm:"pk" json:"id"`
	Name     string `json:"name"`
	Describe string `json:"describe"`
	Version  string `json:"version"`
	Path     string `json:"path"`
	//To          string `json:"to"`
	//IsPackage   int8   `json:"is_package"`
	ProjectId   string    `json:"project_id"`
	CreatedTime time.Time `xorm:"created_time created" json:"created_time"` //创建时间
	Status      string    `xorm:"status" json:"status"`
}
