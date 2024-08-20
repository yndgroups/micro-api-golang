package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var SysFunc = &sysFunc{}

type sysFunc struct{}

// Create
// @Tags 功能管理
// @Summary 功能添加
// @Description 主要用于用户进行功能添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.SysFunc true "功能新增参数"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /func [post]
func (fn *sysFunc) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.SysFunc.Create(ctx))
}

// Update
// @Tags 功能管理
// @Summary 功能修改
// @Description 主要用于用户进行功能修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.SysFunc true "功能参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /func [put]
func (fn *sysFunc) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.SysFunc.Update(ctx))
}

// Delete
// @Tags 功能管理
// @Summary 功能删除
// @Description 主要用于用户进行功能删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "功能id集合"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /func [delete]
func (fn *sysFunc) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.SysFunc.Delete(ctx))
}

// FindList
// @Tags 功能管理
// @Summary 功能分页列表查询
// @Description 主要用于功能分页列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[common.SysFuncParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /func/findFuncList [post]
func (fn *sysFunc) FindList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.SysFunc.FindList(ctx))
}

// FindById
// @Tags 功能管理
// @Summary 根据id功能详情查询
// @Description 主要用于功能详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string  true "功能id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /func/findSysFuncById/{id} [get]
func (fn *sysFunc) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.SysFunc.FindById(ctx))
}

// FindInfoByRoleIds
// @Tags 功能管理
// @Summary 根据角色id查询当前应用的功能及角色已经用有的功能信息
// @Description 主要用于根据角色id查询当前应用的功能及角色已经用有的功能信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param roleId path string  true "应用id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /func/findFuncInfoByRoleIds/{roleId} [get]
func (fn *sysFunc) FindInfoByRoleIds(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.SysFunc.FindInfoByRoleIds(ctx))
}
