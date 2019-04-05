package client

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"ogsyoo/imageExport-api/src/common/conf"
)

var DB *xorm.Engine

func GetConnect() (*xorm.Engine, error) {
	if DB != nil && DB.Ping() == nil {
		return DB, nil
	}
	xormDB, err := xorm.NewEngine("postgres", conf.DatabaseURL)
	if err != nil {
		logrus.Errorf("error creating db instance to %s: %s", conf.DatabaseURL, err)
		return nil, err
	}
	if err = xormDB.Ping(); err != nil {
		logrus.Errorf("error creating db connection to %s: %s", conf.DatabaseURL, err)
	}
	xormDB.SetMapper(core.SnakeMapper{})
	xormDB.ShowSQL(true)
	return xormDB, err
}
