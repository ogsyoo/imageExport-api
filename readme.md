# 发布
## go vendor
```bash
# install govendor
go get -u -v github.com/kardianos/govendor

# 初始化vendor目录
govendor init

# 将本地环境中的依赖Copy到vendor目录中
govendor add +external
```