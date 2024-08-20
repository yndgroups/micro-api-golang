package apis

import (
	"net/http"
	"order/order-consumer/service"

	"github.com/gin-gonic/gin"

	_ "core/req"
	_ "core/resp"
)

var Wechat = &wechat{}

type wechat struct{}

// FindOrderByTradeNo
// @Tags 微信支付
// @Summary 订单支付查询
// @Description 主要用于订单查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param tradeNo path string  true "订单号"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /wechat/findOrderByTradeNo/{tradeNo} [get]
func (w *wechat) FindOrderByTradeNo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Wechat.FindOrderByTradeNo(ctx))
}

// CloseOrder
// @Tags 微信支付
// @Summary 关闭订单
// @Description 主要用于关闭订单
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param tradeNo path string  true "订单号"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /wechat/closeOrder/{tradeNo} [get]
func (w *wechat) CloseOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Wechat.CloseOrder(ctx))
}

// Refunds
// @Tags 微信支付
// @Summary 申请退款
// @Description 主要用于申请退款
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /wechat/weChatRefunds [get]
func (w *wechat) Refunds(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Wechat.Refunds(ctx))
}

// PayNotify
// @Tags 微信支付
// @Summary 微信支付成功回调
// @Description 主要用于微信订单支付成功回调
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /wechat/notify/{confId} [post]
func (w *wechat) Notify(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Wechat.Notify(ctx))
}
