package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var App = &app{}

type app struct{}

// Create
// @Tags 应用管理
// @Summary 应用添加
// @Description 主要用于用户进行应用添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.App true "应用参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /app [post]
func (a *app) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.App.Create(ctx))
}

// Update
// @Tags 应用管理
// @Summary 应用修改
// @Description 主要用于用户进行应用修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.App true "应用修改参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /app [put]
func (a *app) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.App.Update(ctx))
}

// Delete
// @Tags 应用管理
// @Summary 应用删除
// @Description 主要用于应用删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "应用id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /app [delete]
func (a *app) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.App.Delete(ctx))
}

// FindList
// @Tags 应用管理
// @Summary 应用查询
// @Description 主要用于应用删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /app/findAppList [get]
func (a *app) FindList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.App.FindList(ctx))
}

// FindPageList
// @Tags 应用管理
// @Summary 应用分页列表查询
// @Description 主要用于应用分页列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[common.AppParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /app/findAppPageList [post]
func (a *app) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.App.FindPageList(ctx))
}
