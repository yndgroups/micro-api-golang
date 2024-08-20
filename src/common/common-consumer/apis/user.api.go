package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var User = &user{}

type user struct{}

// Create
// @Tags 用户管理
// @Summary 添加用户信息
// @Description 主要用于添加用户信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.User true "用户信息数据"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /user [post] 添加用户
func (u *user) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.User.Create(ctx))
}

// Update
// @Tags 用户管理
// @Summary 修改用户信息
// @Description 主要用于修改用户信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.User true "用户信息"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /user [put]
func (u *user) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.User.Update(ctx))
}

// Delete
// @Tags 用户管理
// @Summary 删除用户信息
// @Description 主要用于删除用户信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "用户id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /user [delete]
func (u *user) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.User.Delete(ctx))
}

// FindWeChatWebUserInfo
// @Tags 用户管理
// @Summary 查询微信公众号用户信息
// @Description 主要用于查询微信公众号用户信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param confId query string  true "配置id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /user/findWeChatWebUserInfo [get]
func (u *user) FindWeChatWebUserInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.User.FindWeChatWebUserInfo(ctx))
}

// FindById
// @Tags 用户管理
// @Summary 查询用户详情
// @Description 主要查询用户详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id query string  true "用户id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /user/findById/{id} [get]
func (u *user) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.User.FindById(ctx))
}

// FindPageList
// @Tags 用户管理
// @Summary 用户分页列表
// @Description 主要用于查询用户分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body common.UserPageAuthQuery true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /user/findPageList [post]
func (u *user) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.User.FindPageList(ctx))
}

// UpdateDetail
// @Tags 用户管理
// @Summary 修改用户本人基本信息
// @Description 主要用于修改用户本人基本信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /user/updateDetail [post]
func (u *user) UpdateDetail(ctx *gin.Context) {
	// @Param userBaseInfo body common.BaseUserInfo  true "查询参数"
	ctx.JSON(http.StatusOK, service.UserDetail.UpdateDetail(ctx))
}

// FindDetail
// @Tags 用户管理
// @Summary 获取用户本人基本信息
// @Description 主要用于获取用户本人基本信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /user/findDetail [get]
func (u *user) FindDetail(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.UserDetail.FindDetail(ctx))
}

// FindBalance
// @Tags 用户管理
// @Summary 查询用户本人余额信息
// @Description 主要用于查询用户本人余额信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /user/findBalance [get]
func (u *user) FindBalance(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.User.FindBalance(ctx))
}

// Logout
// @Tags 用户管理
// @Summary 退出登录
// @Description 这是一个用户退出系统登录的接口
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /user/logout [get]
func (u *user) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.User.Logout(ctx))
}

// FindUserIdByUserParam
// @Tags 用户管理
// @Summary 根据指定条件查询用户id
// @Description 主要用于根据指定条件查询用户id
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param params body common.UserParam  true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /user/findUserIdByUserParam [post]
func (u *user) FindUserIdByUserParam(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.User.FindUserIdByUserParam(ctx))
}
