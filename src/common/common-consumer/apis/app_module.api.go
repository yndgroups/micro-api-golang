package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var AppModule = &appModule{}

type appModule struct {
}

// Create
// @Tags 应用模块管理
// @Summary 应用模块添加
// @Description 主要用于应用模块添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.AppModule true "应用参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /appModule [post]
func (am *appModule) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.AppModule.Create(ctx))
}

// Update
// @Tags 应用模块管理
// @Summary 应用模块修改
// @Description 主要用于应用模块修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.AppModule true "应用修改参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /appModule [put]
func (am *appModule) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.AppModule.Update(ctx))
}

// Delete
// @Tags 应用模块管理
// @Summary 应用模块删除
// @Description 主要用于应用模块删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "应用模块id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /appModule [delete]
func (am *appModule) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.AppModule.Delete(ctx))
}

// FindById
// @Tags 应用模块管理
// @Summary 应用模块详情查询
// @Description 主要用于应用模块详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id query string  true "地区id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /appModule/findAreaById/{id} [get]
func (am *appModule) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.AppModule.FindById(ctx))
}

// FindList
// @Tags 应用模块管理
// @Summary 应用模块查询
// @Description 主要用于应用模块删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param appId query string  true "应用id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /appModule/findAppModuleListByAppId [get]
func (am *appModule) FindList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.AppModule.FindList(ctx))
}
