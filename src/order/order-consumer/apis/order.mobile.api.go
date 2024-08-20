package apis

import (
	"net/http"
	"order/order-consumer/service"

	"github.com/gin-gonic/gin"

	_ "core/req"
	_ "core/resp"
)

var OrderMobile = &orderMobile{}

type orderMobile struct{}

// Delete
// @Tags 订单管理-移动端
// @Summary 删除订单
// @Description 主要用于订单删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /order/m [delete]
func (om *orderMobile) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Order.Delete(ctx))
}

// FindPageList
// @Tags 订单管理-移动端
// @Summary 订单分页列表-移动端
// @Description 主要用于查询订单分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[order.OrderMobileParams] true "查询参数"
// @Success 200 {object} resp.Response{data=resp.Pager{list=[]order.OrderListVo}}  "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /order/m/findPageList [post]
func (om *orderMobile) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Order.FindPageList(ctx))
}
