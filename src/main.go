package main

import (
	"fmt"
	"ogsyoo/imageExport-api/src/common/client"
	"ogsyoo/imageExport-api/src/common/conf"
	"ogsyoo/imageExport-api/src/dao"
	"ogsyoo/imageExport-api/src/router"
	"ogsyoo/imageExport-api/src/sse"
	"os"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"net/http"
	"time"
)

var (
	port     = pflag.Int("port", 8081, "")
	prefix   = pflag.String("prefix", "/export", "")
	docPath  = pflag.String("docPath", "/app/document", "")
	dbURL    = pflag.String("dbURL", "host=localhost port=54321 user=postgres password=password dbname=postgres sslmode=disable", "")
	redisURL = pflag.String("redisURL", "redis://localhost:6379", "")
	packeDoc = pflag.String("packeDoc", "d:\\images", "")
	uiDoc    = pflag.String("uiDoc", "/app/dist", "")
)

func main() {
	pflag.Parse()
	initEnv()
	initConfig()
	initDB()
	conf.SseClient = sse.NewServer(nil)
	server()
}

func server() error {
	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true),
	)
	fmt.Println(fmt.Sprintf(":%d", *port))
	return http.ListenAndServe(
		fmt.Sprintf(":%d", *port),
		handler,
	)
}

//注册环境变量
func initEnv() {
	if os.Getenv("TRANS_REDIS_URL") != "" {
		*redisURL = os.Getenv("TRANS_REDIS_URL")
	}
	if os.Getenv("TRANS_DB_URL") != "" {
		*dbURL = os.Getenv("TRANS_DB_URL")
	}
	if os.Getenv("PACKAGE_DOC") != "" {
		*packeDoc = os.Getenv("PACKAGE_DOC")
	}
}

//注册公共config参数信息
func initConfig() {
	conf.BaseInfo.Prefix = *prefix
	conf.DatabaseURL = *dbURL
	conf.DocPath = *docPath
	conf.RedisURL = *redisURL
	conf.PackeDoc = *packeDoc
	conf.UiDoc = *uiDoc
}

//初始化数据库
func initDB() {
	db, err := client.GetConnect()
	if err != nil {
		fmt.Println("get db error:", err.Error())
		return
	}
	db.Sync2(dao.Project{}, dao.ImageJob{})
}
