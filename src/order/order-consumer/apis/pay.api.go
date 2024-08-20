package apis

import (
	"net/http"
	"order/order-consumer/service"

	"github.com/gin-gonic/gin"

	_ "core/req"
	_ "core/resp"
)

var Pay = &pay{}

type pay struct{}

// UpdateUserPayPwd
// @Tags 支付相关
// @Summary 设置支付密码
// @Description 主要用于设置支付密码
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pwdParam body order.SetPayPwdParam  true "支付参数设置"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /pay/updatePayPwd [post]
func (p *pay) UpdateUserPayPwd(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Pay.UpdateUserPayPwd(ctx))
}
