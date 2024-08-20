package routes

import (
	"core/middleware"
	_ "order/order-consumer/docs"
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
	basePath := "/api/order"

	// 接口文档
	routes.GET(basePath+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(basePath+"/swagger/doc.json")))

	// 微信支付
	Wechat.InitRoutes(routes, basePath)

	PocketMoney.InitRoutes(routes, basePath)

	// 订单
	Order.InitRoutes(routes, basePath)

	// 订单移动端版
	OrderMobile.InitRoutes(routes, basePath)

	// 支付相关信息
	Pay.InitRoutes(routes, basePath)

	// 运费-模板管理
	Templates.InitRoutes(routes, basePath)

	// 免邮-模板管理
	TemplatesFree.InitRoutes(routes, baseApiPath)

	// 区域运费-模板管理
	TemplatesRegion.InitRoutes(routes, basePath)

	// 不送达-模板管理
	TemplatesUndelivered.InitRoutes(routes, basePath)

	return routes, basePath
}
