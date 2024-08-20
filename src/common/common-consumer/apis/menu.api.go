package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Menu = &menu{}

type menu struct{}

// @Tags 菜单管理
// @Summary 菜单添加
// @Description 主要用于菜单添加
// @Param accessToken header string true "授权令牌"  default(Bearer token)
// @Param object body common.Menu true "菜单参数"
// @Success 200 {object} resp.Response 成功时返回数据结构
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /menu [post]
func (m *menu) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Menu.Create(ctx))
}

// @Tags 菜单管理
// @Summary 菜单修改
// @Description 主要用于菜单修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Menu true "登录对象参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /menu [put]
func (m *menu) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Menu.Update(ctx))
}

// @Tags 菜单管理
// @Summary 菜单删除
// @Description 主要用于菜单删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "菜单id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /menu [delete]
func (m *menu) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Menu.Delete(ctx))
}

// @Tags 菜单管理
// @Summary 根据用户权限进行菜单查询
// @Description 主要用于根据用户权限进行树状菜单查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /menu/findTreeList [get]
func (m *menu) FindTreeList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Menu.FindTreeList(ctx))
}

// @Tags 菜单管理
// @Summary 查询应用下的所有菜单
// @Description 主要用于根据用户权限进行树状菜单查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /menu/findAll [get]
func (m *menu) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Menu.FindAll(ctx))
}

// @Tags 菜单管理
// @Summary 根据角色id查询当前应用的菜单及角色已经用有的菜单信息
// @Description 主要用于根据角色id查询当前应用的菜单及角色已经用有的菜单信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param roleId path string  true "应用id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /menu/findByRoleIds/{roleId} [get]
func (m *menu) FindInfoByRoleIds(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Menu.FindInfoByRoleIds(ctx))
}
