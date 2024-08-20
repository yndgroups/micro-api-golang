package routes

import (
	"core/middleware"
	"fmt"
	"order/order-consumer/apis"

	"github.com/gin-gonic/gin"
)

var Wechat = &wechat{}

type wechat struct{}

func (w *wechat) InitRoutes(routes *gin.Engine, basePath string) {

	wechat := routes.Group(fmt.Sprintf("%s/wechat", basePath))
	// 支付回调地址
	wechat.POST("/notify/:confId", middleware.Authorization("anonymous"), apis.Wechat.Notify)
	// 关闭订单
	wechat.GET("/closeOrder/:tradeNo", middleware.Authorization("order:c:pay"), apis.Wechat.CloseOrder)
	// 退款申请
	wechat.POST("/refunds", middleware.Authorization("order:c:pay"), apis.Wechat.Refunds)
	// 根据订单号查询微信充值记录
	wechat.GET("/findOrderByTradeNo/:tradeNo", middleware.Authorization("order:c:pay"), apis.Wechat.FindOrderByTradeNo)
	// 关闭订单
	wechat.POST("/close", middleware.Authorization("order:c:manage"), apis.Wechat.CloseOrder)
}
