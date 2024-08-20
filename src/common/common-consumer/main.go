package main

import (
	"common/common-consumer/routes"
	"core/app"

	_ "github.com/go-sql-driver/mysql"
)

// @Title 公共服务消费者API文档
// @Description Swagger Go API文档 将登录信息获取的token替换 "Bearer token" 里的token字符串，注意中间的空格
// @Contact.name 杨大琼
// @Contact.url http://localhost
// @Contact.email 826422592@qq.com
// @License.name Apache 2.0
// @Version 3.0.0
// @BasePath /api/common
func main() {
	app := app.AppLaunch{}
	app.StartHttpServer(routes.InitRoutes(app.Active))
}
