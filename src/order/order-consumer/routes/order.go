package routes

import (
	"core/middleware"
	"fmt"
	"order/order-consumer/apis"

	"github.com/gin-gonic/gin"
)

var Order = &order{}

type order struct{}

func (o *order) InitRoutes(routes *gin.Engine, basePath string) {

	order := routes.Group(fmt.Sprintf("%s/order", basePath))
	// 统一单
	order.POST("", middleware.Authorization("order:c:manage"), apis.Order.Create)
	// 订单支付
	order.POST("/pay", middleware.Authorization("order:c:manage"), apis.Order.Pay)
	// 订单删除
	order.DELETE("", middleware.Authorization("order:c:manage"), apis.Order.Delete)
	// 订单列表
	order.POST("/findPageList", middleware.Authorization("order:c:manage"), apis.Order.FindPageList)
	// 订单详情
	order.GET("/findById/:id", middleware.Authorization("order:c:manage"), apis.Order.FindById)
	// 根据订单参数查询订单信息
	order.POST("/findByParams", middleware.Authorization("order:c:manage"), apis.Order.FindByParams)
	// 订单统计
	order.GET("/statistics", middleware.Authorization("order:c:manage"), apis.Order.OrderStatistics)
}
