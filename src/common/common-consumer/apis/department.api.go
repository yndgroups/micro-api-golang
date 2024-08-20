package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Department = &department{}

type department struct{}

// Create
// @Tags 部门管理
// @Summary 部门添加
// @Description 主要用于用户进行部门添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Department true "部门参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /dept [post]
func (dept *department) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Department.Create(ctx))
}

// Update
// @Tags 部门管理
// @Summary 部门修改
// @Description 主要用于用户进行部门修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Department true "部门修改参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /dept [put]
func (dept *department) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Department.Update(ctx))
}

// Delete
// @Tags 部门管理
// @Summary 部门删除
// @Description 主要用于部门删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "部门id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /dept [delete]
func (dept *department) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Department.Delete(ctx))
}

// FindTree
// @Tags 部门管理
// @Summary 部门信息查询
// @Description 主要用于部门信息查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /dept/findTree [get]
func (dept *department) FindTree(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Department.FindTree(ctx))
}
