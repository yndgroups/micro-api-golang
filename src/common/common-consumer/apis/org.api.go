package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Org = &org{}

type org struct{}

// Create
// @Tags 企业或组织管理
// @Summary 创建企业或组织
// @Description 主要用于企业或组织新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param org body common.Org true "新增企业或组织参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /org [post]
func (o *org) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Org.Create(ctx))
}

// Update
// @Tags 企业或组织管理
// @Summary 修改企业或组织
// @Description 主要用于企业或组织修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param org body common.Org true "修改企业或组织参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /org [put]
func (o *org) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Org.Update(ctx))
}

// Delete
// @Tags 企业或组织管理
// @Summary 删除企业或组织
// @Description 主要用于企业或组织删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /org [delete]
func (o *org) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Org.Delete(ctx))
}

// FindPageList
// @Tags 企业或组织管理
// @Summary 企业或组织分页列表
// @Description 主要用于查询企业或组织分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[common.OrgParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /org/findPageList [post]
func (o *org) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Org.FindPageList(ctx))
}

// FindById
// @Tags 企业或组织管理
// @Summary 根据id查询企业或组织内容详情
// @Description 根据id查询企业或组织内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "企业或组织id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /org/findById/{id} [get]
func (o *org) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Org.FindById(ctx))
}
