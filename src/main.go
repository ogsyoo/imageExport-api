package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"ogsyoo/imageExport-api/src/common/conf"
	"ogsyoo/imageExport-api/src/router"
	"os"

	"net/http"
	"time"
)

var (
	port     = pflag.Int("port", 8081, "")
	prefix   = pflag.String("prefix", "/export", "")
	docPath  = pflag.String("docPath", "/app/document", "")
	dbURL    = pflag.String("dbURL", "host=localhost port=5432 user=postgres password=passwd123 dbname=postgres sslmode=disable", "")
	redisURL = pflag.String("redisURL", "redis://localhost:6379", "")
)

func main() {
	pflag.Parse()
	initEnv()
	initConfig()
	server()
}

func fileServer() {
	fmt.Println("path:", fmt.Sprintf(`%s/image`, *prefix), *docPath+"/images")
	http.Handle("/trans/api", http.FileServer(http.Dir(*docPath+"/images")))
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
}

//注册公共config参数信息
func initConfig() {
	conf.BaseInfo.Prefix = *prefix
	conf.DatabaseURL = *dbURL
	conf.DocPath = *docPath
	conf.RedisURL = *redisURL
}
