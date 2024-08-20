package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Config = &config{}

type config struct{}

// Create
// @Tags 配置管理
// @Summary 配置添加
// @Description 主要用于进行配置添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Config true "配置参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /config [post]
func (cfg *config) Create(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, "开发中...")
}

// Update
// @Tags 配置管理
// @Summary 配置修改
// @Description 主要用于进行配置修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Config  true "配置参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /config [put]
func (cfg *config) Update(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, "开发中...")
}

// Delete
// @Tags 配置管理
// @Summary 配置删除
// @Description 主要用于进行配置删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "配置id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /config [delete]
func (cfg *config) Delete(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, "开发中...")
}

// FindById
// @Tags 配置管理
// @Summary 配置删除
// @Description 主要用于进行配置删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id query string  true "配置id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /config/findConfigById/{id} [get]
func (cfg *config) FindById(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, "开发中...")
}

// FindByConfigExpandParam
// @Tags 配置管理
// @Summary 根据参数条件查询配置文件
// @Description 主要用于根据参数条件查询配置文件
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param params body common.ConfigIds true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /config/findByConfigExpandParam [post]
func (cfg *config) FindByConfigExpandParam(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, service.Config.FindByConfigExpandParam(ctx))
}
