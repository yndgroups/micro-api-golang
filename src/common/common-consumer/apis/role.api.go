package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Role = &role{}

type role struct{}

// Create
// @Tags 角色管理
// @Summary 角色添加
// @Description 主要用于角色添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Role true "角色参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /role [post]
func (r *role) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Role.Create(ctx))
}

// Update
// @Tags 角色管理
// @Summary 角色修改
// @Description 主要用于角色修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Role true "角色参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /role [put]
func (r *role) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Role.Update(ctx))
}

// Delete
// @Tags 角色管理
// @Summary 角色删除
// @Description 主要用于角色删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "角色id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /role [delete]
func (r *role) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Role.Delete(ctx))
}

// FindPageList
// @Tags 角色管理
// @Summary 角色列表查询
// @Description 主要用于角色列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[common.RoleParam] true "查询参数"
// @Success 200 {object} resp.Response{data=resp.Pager{list=[]common.Role}} "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /role/findPageList [post]
func (r *role) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Role.FindPageList(ctx))
}

// FindAuthListByUserId
// @Tags 角色管理
// @Summary 根据用户id查询其在应用下的角色及授权信息
// @Description 根据用户id查询其在应用下的角色及授权信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param userId path string true "用户id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /role/findAuthRoleListByUserId/{userId} [get]
func (r *role) FindAuthListByUserId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Role.FindAuthListByUserId(ctx))
}
