package apis

import (
	"net/http"
	"order/order-consumer/service"

	"github.com/gin-gonic/gin"

	_ "core/req"
	_ "core/resp"
)

var PocketMoney = &pocketMoney{}

type pocketMoney struct{}

// GetBalance
// @Tags 零钱管理
// @Summary 查询零钱余额
// @Description 主要用于查询零钱余额
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /pocketMoney/getBalance [get]
func (pt *pocketMoney) GetBalance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.PocketMoney.GetBalance(ctx))
}

// FindPageList
// @Tags 零钱管理
// @Summary 充值记录
// @Description 主要用于查询充值记录
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body order.PocketMoneyPageAuthQuery true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /pocketMoney/findPageList [post]
func (pt *pocketMoney) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.PocketMoney.FindPageList(ctx))
}

// OnlineRecharge returns
// @Tags 零钱管理
// @Summary 余额充值
// @Description 主要用于余额充值
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param amount path string  true "金额"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /pocketMoney/create/{amount} [get]
func (pt *pocketMoney) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.PocketMoney.Create(ctx))
}

// FindById
// @Tags 模板管理-运费模板
// @Summary 根据id查询详情
// @Description 主要用于根据id查询详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string  true "模板id"
// @Success 200 {object} resp.Response{data=order.FreightTemplates} "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /pocketMoney/findById/{id} [get]
func (ft *pocketMoney) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.PocketMoney.FindById(ctx))
}
