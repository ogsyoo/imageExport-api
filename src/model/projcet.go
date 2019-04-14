package model

import "ogsyoo/imageExport-api/src/dao"

type ProjectImages struct {
	Project dao.Project    `json:"project"`
	Images  []dao.ImageJob `json:"images"`
}

type ProImages struct {
	dao.Project `xorm:"extends" json:"project"`
	Images      dao.ImageJob `xorm:"extends" json:"images"`
}

func (ProImages) TableName() string {
	return "project"
}
