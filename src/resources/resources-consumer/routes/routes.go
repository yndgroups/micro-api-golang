package routes

import (
	"core/middleware"
	"fmt"
	"resources/resources-consumer/apis"
	_ "resources/resources-consumer/docs"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(env string) (route *gin.Engine, baseApiPath string) {
	routes := gin.Default()

	// 环境模式设置
	switch true {
	case strings.Contains(env, "test"):
		gin.SetMode(gin.TestMode)
	case strings.Contains(env, "prod"):
		gin.SetMode(gin.ReleaseMode)
		// 生产环境捕获异常并让成功不要挂掉
		routes.Use(middleware.Recover)
	default:
		gin.SetMode(gin.TestMode)
	}

	// api基础路径
	basePath := "/api/resources"
	// 接口文档
	routes.GET(basePath+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(basePath+"/swagger/doc.json")))

	// 文件管理管理
	app := routes.Group(fmt.Sprintf("%s/file", basePath))
	// 单个文件上传 resources:file:singleUpload
	app.POST("/singleUpload/:sizeType", middleware.Authorization("anonymous"), apis.File.SingleUpload)
	// 单个文件上传 resources:file:multipleUpload
	app.POST("/multipleUpload/:sizeType", middleware.Authorization("anonymous"), apis.File.MultipleUpload)
	// 文件列表
	app.POST("/getList", middleware.Authorization("anonymous"), apis.File.GetList)
	// 文件删除
	app.POST("/delete", middleware.Authorization("anonymous"), apis.File.Delete)

	return routes, basePath
}
