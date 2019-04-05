package service

import (
	"errors"
	"ogsyoo/imageExport-api/src/common/client"
	"ogsyoo/imageExport-api/src/dao"
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
