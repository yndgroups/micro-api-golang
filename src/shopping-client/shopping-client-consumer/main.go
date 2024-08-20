package main

import (
	"core/app"
	"shopping-client/shopping-client-consumer/routes"

	_ "github.com/go-sql-driver/mysql"
)

// @Title 客户端接口文档
// @Description swagger go API文档
// @Contact.name 杨大琼
// @Contact.url http://localhost
// @Contact.email 826422592@qq.com
// @License.name Apache 2.0
// @BasePath /api/shopping/client
func main() {
	app := app.AppLaunch{}
	app.StartHttpServer(routes.InitRoutes(app.Active))
}
