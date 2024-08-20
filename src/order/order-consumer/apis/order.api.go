package apis

import (
	"net/http"
	"order/order-consumer/service"

	"github.com/gin-gonic/gin"

	_ "core/req"
	_ "core/resp"
)

var Order = &order{}

type order struct{}

// Create
// @Tags 订单管理
// @Summary 统一下单
// @Description 主要用于统一下单
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param createOder body order.CreateOrderParams true "应用参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /order [post]
func (o *order) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Order.Create(ctx))
}

// Pay
// @Tags 订单管理
// @Summary 订单支付
// @Description 主要用于订单支付
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body order.PayOrderParam true "订单支付参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /order/pay [post]
func (o *order) Pay(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Order.Pay(ctx))
}

// Delete
// @Tags 订单管理
// @Summary 删除订单
// @Description 主要用于订单删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /order [delete]
func (o *order) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Order.Delete(ctx))
}

// FindPageList
// @Tags 订单管理
// @Summary 订单分页列表
// @Description 主要用于查询订单分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[order.OrderParams] true "查询参数"
// @Success 200 {object} resp.Response{data=resp.Pager{list=[]order.OrderListVo}}  "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /order/findPageList [post]
func (o *order) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Order.FindPageList(ctx))
}

// FindById
// @Tags 订单管理
// @Summary 根据id查询订单内容详情
// @Description 根据id查询订单内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "订单id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /order/findById/{id} [get]
func (o *order) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Order.FindById(ctx))
}

// OrderStatistics
// @Tags 订单管理
// @Summary 订单统计
// @Description 订单统计
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /order/statistics [get]
func (o *order) OrderStatistics(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Order.OrderStatistics(ctx))
}

// FindByParams
// @Tags 订单管理
// @Summary 根据订单参数查询订单信息
// @Description 根据订单参数查询订单信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body order.OrderParams true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /order/findByParams [post]
func (o *order) FindByParams(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Order.FindByParams(ctx))
}
