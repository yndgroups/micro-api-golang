package apis

import (
	_ "core/req"
	_ "core/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-admin/shopping-admin-consumer/service"
)

var Ad = &ad{}

type ad struct{}

// Create
// @Tags 广告管理
// @Summary 创建广告
// @Description 主要用于广告新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param advertisement body shopping_admin.Ad true "新增广告参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /ad [post]
func (a *ad) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Ad.Create(ctx))
}

// Update
// @Tags 广告管理
// @Summary 修改广告
// @Description 主要用于广告修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param advertisement body shopping_admin.Ad true "修改广告参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /ad [put]
func (a *ad) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Ad.Update(ctx))
}

// Delete
// @Tags 广告管理
// @Summary 删除广告
// @Description 主要用于广告删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "广告是删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /ad [delete]
func (a *ad) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Ad.Delete(ctx))
}

// FindPageList
// @Tags 广告管理
// @Summary 广告分页列表
// @Description 主要用于查询广告分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[shopping_admin.AdParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /ad/findPageList [post]
func (a *ad) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Ad.FindPageList(ctx))
}

// FindById
// @Tags 广告管理
// @Summary 根据id查询广告内容详情
// @Description 根据id查询广告内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "广告id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /ad/findById/{id} [get]
func (a *ad) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Ad.FindById(ctx))
}
