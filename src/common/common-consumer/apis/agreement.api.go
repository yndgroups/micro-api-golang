package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Agreement = &agreement{}

type agreement struct{}

// Create
// @Tags 协议管理
// @Summary 创建协议
// @Description 主要用于协议新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param agreement body common.Agreement true "新增协议参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /agreement [post]
func (a *agreement) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Agreement.Create(ctx))
}

// Update
// @Tags 协议管理
// @Summary 修改协议
// @Description 主要用于协议修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param agreement body common.Agreement true "修改协议参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /agreement [put]
func (a *agreement) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Agreement.Update(ctx))
}

// Delete
// @Tags 协议管理
// @Summary 删除协议
// @Description 主要用于协议删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /agreement [delete]
func (a *agreement) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Agreement.Delete(ctx))
}

// FindPageList
// @Tags 协议管理
// @Summary 协议分页列表
// @Description 主要用于查询协议分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[common.AgreementParam]{params=common.AgreementParam{}} true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /agreement/findPageList [post]
func (a *agreement) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Agreement.FindPageList(ctx))
}

// FindById
// @Tags 协议管理
// @Summary 根据id查询协议内容详情
// @Description 根据id查询协议内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "协议id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /agreement/findById/{id} [get]
func (a *agreement) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Agreement.FindById(ctx))
}
