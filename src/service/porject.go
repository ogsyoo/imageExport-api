package service

import (
	"errors"
	"github.com/satori/go.uuid"
	"ogsyoo/imageExport-api/src/common/client"
	"ogsyoo/imageExport-api/src/dao"
	"ogsyoo/imageExport-api/src/model"
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
	err = db.Limit(l, o).Find(&list)
	return
}

func (p *Project) GetProImage(pid string) (res map[string]interface{}, err error) {
	db, err := client.GetConnect()
	if err != nil {
		return
	}
	project := make([]model.ProImages, 0)
	res = make(map[string]interface{})
	err = db.Join("INNER", "image_job", "project.id = image_job.project_id").Where("project.id = ?", pid).Find(&project)
	if len(project) > 0 {
		res["project"] = project[0].Project
		res["images"] = []dao.ImageJob{}
		for _, v := range project {
			res["images"] = append(res["images"].([]dao.ImageJob), v.Images)
		}
	}
	return
}

func (p *Project) InsertPorject(pi *model.ProjectImages) (err error) {
	db, err := client.GetConnect()
	//打开事务
	trans := db.NewSession()
	defer trans.Close()
	if err := trans.Begin(); err != nil {
		return err
	}
	//插入project
	id, _ := uuid.NewV4()
	pi.Project.Id = id.String()
	affect, err := trans.Insert(pi.Project)
	if affect == 0 || err != nil {
		trans.Rollback()
		err = errors.New("插入失败！")
		return
	}
	//插入images
	for _, v := range pi.Images {
		id, _ := uuid.NewV4()
		v.Id = id.String()
		v.ProjectId = pi.Project.Id
		aff, err := trans.Insert(v)
		if aff == 0 || err != nil {
			trans.Rollback()
			err = errors.New("插入失败！")
			return err
		}
	}
	trans.Commit()
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
