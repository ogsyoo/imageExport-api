package service

import (
	"errors"
	"ogsyoo/imageExport-api/src/common/client"
	"ogsyoo/imageExport-api/src/dao"
	"strconv"
)

type Project struct {
}

func (p *Project) GetListPage(limit string, offset string) (list []dao.Project, err error) {
	l, _ := strconv.Atoi(limit)
	o, _ := strconv.Atoi(offset)
	db, err := client.GetConnect()
	if err != nil {
		return
	}
	err = db.Where("status > 0").Limit(l, o).Find(&list)
	return
}

func (p *Project) InsertPorject(project *dao.Project) (err error) {

	db, err := client.GetConnect()
	if err != nil {
		return
	}
	affect, err := db.Insert(project)
	if affect == 0 || err != nil {
		err = errors.New("插入失败！")
	}
	return
}

func (p *Project) DelPorject(id int) (err error) {
	db, err := client.GetConnect()
	if err != nil {
		return
	}
	var project dao.Project
	affect, err := db.ID(id).Delete(&project)
	if affect == 0 || err != nil {
		err = errors.New("删除失败！")
	}
	return
}

func (p *Project) UpdatePorject(project *dao.Project) (err error) {
	db, err := client.GetConnect()
	if err != nil {
		return
	}
	affect, err := db.ID(project.Id).Update(project)
	if affect == 0 || err != nil {
		err = errors.New("更新失败！")
	}
	return
}
