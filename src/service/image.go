package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"ogsyoo/imageExport-api/src/common/client"
	"ogsyoo/imageExport-api/src/common/conf"
	"ogsyoo/imageExport-api/src/dao"
	"ogsyoo/imageExport-api/src/sse"
)

type Image struct {
}

func (i *Image) InsertImageList(ls []*dao.ImageJob) (affect int64, err error) {
	db, err := client.GetConnect()
	if err != nil {
		return
	}
	affect, err = db.Insert(ls)
	if err != nil || affect == 0 {
		err = errors.New("批量插入失败")
	}
	return
}

func (i *Image) UpdateImageJob(im *dao.ImageJob) (affect int64, err error) {
	if im.Id == "" {
		return 0, errors.New("获取数据失败")
	}
	db, err := client.GetConnect()
	if err != nil {
		return
	}
	affect, err = db.Id(im.Id).Update(im)
	if err != nil || affect == 0 {
		err = errors.New("更新失败")
	}
	return
}

func (i *Image) DeleteImageJob(id string) (affect int64, err error) {
	if id == "" {
		return 0, errors.New("获取数据失败")
	}
	db, err := client.GetConnect()
	if err != nil {
		return
	}
	image := new(dao.ImageJob)
	affect, err = db.Id(id).Delete(image)
	if err != nil || affect == 0 {
		err = errors.New("删除失败")
	}
	return
}

func (i *Image) PackageImage(ids []dao.ImageJob) (err error) {
	if len(ids) == 0 {
		return errors.New("请选择要打包的镜像")
	}
	//修改数据库状态
	var path string
	err = Login("admin", "spaceIN511", "hub.wodcloud.com")
	if err != nil {
		fmt.Println("登陆失败", err)
		return err
	}
	db, _ := client.GetConnect()
	for _, v := range ids {
		v.Status = "running"
		db.Update(&v)
		sendMessage(v)
		path = fmt.Sprintf("%s/%s:%s", v.Path, v.Name, v.Version)
		err = Pull(path)
		if err != nil {
			fmt.Println("下载镜像失败", err)
			v.Status = "err"
			db.Update(&v)
			sendMessage(v)
			continue
		}
		tagname := fmt.Sprintf("%s/%s:%s", "reg.local:5000", v.Name, v.Version)
		err = Tag(path, tagname)
		if err != nil {
			fmt.Println("镜像打标签失败", err)
			v.Status = "err"
			db.Update(&v)
			sendMessage(v)
			continue
		}
		err = Save(tagname, conf.PackeDoc, v.Name+"-"+v.Version+".tar")
		if err != nil {
			fmt.Println("镜像保存失败", err)
			v.Status = "err"
			sendMessage(v)
			db.Update(&v)
			continue
		}
		v.Status = "success"
		db.Update(&v)
		sendMessage(v)
	}
	return
}

func sendMessage(image dao.ImageJob) {
	b, _ := json.Marshal(image)
	mes := sse.NewMessage(image.Id, string(b), "")
	fmt.Println(mes)
	conf.SseClient.SendMessage("images_package", mes)
}
