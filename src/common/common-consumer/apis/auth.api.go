package apis

import (
	"common/common-consumer/service"
	"net/http"

	"core/model"
	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Auth = &auth{}

type auth struct{}

// @Tags 授权管理
// @Summary 测试发送消息
// @Description 测试发送消息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /auth/sendMsg [get]
func (auth *auth) SendMsg(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.SendMsg(ctx))
}

// @Tags 授权管理
// @Summary 创建匿名token
// @Description 主要用于创建匿名token
// @Param appId path string  true "appId"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /auth/createAnonymityToken/{appId} [get]
func (auth *auth) CreateAnonymityToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.CreateAnonymityToken(ctx))
}

// @Tags 授权管理
// @Summary 用户登录
// @Description 主要用于用户登录
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body model.LoginDTO true "用户登录"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /auth/login [post]
func (auth *auth) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.Login(ctx))
}

// @Tags 授权管理
// @Summary 退出登陆
// @Description 用户退出系统登录
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /auth/logout [get]
func (auth *auth) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.Logout(ctx))
}

// @Tags 授权管理
// @Summary 刷新权限
// @Description 用户刷新权限
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /auth/refreshPermissions [get]
func (auth *auth) RefreshPermissions(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.RefreshPermissions(ctx))
}

// WechatLogin
// @Tags 授权管理
// @Summary 微信小程序登录
// @Description 主要用于微信小程序登录
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body model.WechatLoginParams true "微信小程序登录参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /auth/wechatLogin [post]
func (auth *auth) WechatLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.WechatLogin(ctx))
}

// WechatWebLogin
// @Tags 授权管理
// @Summary 微信公众号登录
// @Description 主要用于微信公众号登录
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body model.WechatLoginParams true "微信公众号登录参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /auth/wechatWebLogin [post]
func (auth *auth) WechatWebLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.WechatWebLogin(ctx))
}

// GetWeChatJSSDK
// @Tags 授权管理
// @Summary 获取微信 jssdk 授权
// @Description 主要用于获取微信 jssdk 授权
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body model.JsSdkDTO true "用户登录"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /auth/getWeChatJSSDK [post]
func (auth *auth) GetWeChatJSSDK(ctx *gin.Context) {
	jsSdk := model.JsSdkDTO{}
	ctx.BindJSON(&jsSdk)
	ctx.JSON(200, "开发中...")
}

// CreateRoleMenu
// @Tags 授权管理
// @Summary 新增角色关联菜单数据
// @Description 主要用于新增角色关联菜单数据
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param agreement body []common.RoleMenu true "角色菜单参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /auth/createRoleMenu [post]
func (auth *auth) CreateRoleMenu(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.CreateRoleMenu(ctx))
}

// CreateUserRole
// @Tags 授权管理
// @Summary 新增角色与用户关联
// @Description 新增角色与用户关联
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param agreement body []common.UserRole true "用户角色参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /auth/createUserRole [post]
func (auth *auth) CreateUserRole(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.CreateUserRole(ctx))
}

// CreateRoleFunc
// @Tags 授权管理
// @Summary 新增角色与功能关联
// @Description 新增角色与功能关联
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param agreement body []common.RoleFunc true "角色功能参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /auth/createRoleFunc [post]
func (auth *auth) CreateRoleFunc(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.CreateRoleFunc(ctx))
}

// GetToken
// @Tags 授权管理
// @Summary 根据appId获取token
// @Description 主要用于appId获取token,用于新应用数据获取或这切换应用之间进行切换
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param appId path string  true "appId"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /auth/getToken/{appId} [get]
func (auth *auth) GetToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Auth.GetToken(ctx))
}
