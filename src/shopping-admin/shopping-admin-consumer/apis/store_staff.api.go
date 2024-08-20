package apis

import (
	_ "core/req"
	_ "core/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-admin/shopping-admin-consumer/service"
)

type storeStaff struct{}

var StoreStaff = &storeStaff{}

// Create
// @Tags 员工管理
// @Summary 创建员工信息
// @Description 主要用于员工新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param storeStaff body shopping_admin.StoreStaff true "新增员工参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /storeStaff [post]
func (ss *storeStaff) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.StoreStaff.Create(ctx))
}

// Update
// @Tags 员工管理
// @Summary 修改员工信息
// @Description 主要用于员工新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param storeStaff body shopping_admin.StoreStaff true "新增员工参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /storeStaff [put]
func (ss *storeStaff) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.StoreStaff.Update(ctx))
}

// Delete
// @Tags 员工管理
// @Summary 删除员工
// @Description 主要用于员工删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /storeStaff [delete]
func (ss *storeStaff) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.StoreStaff.Delete(ctx))
}

// FindPageList
// @Tags 员工管理
// @Summary 员工分页列表
// @Description 主要用于查询员工分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body shopping_admin.StoreStaffPageAuthQuery true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /storeStaff/findPageList [post]
func (ss *storeStaff) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.StoreStaff.FindPageList(ctx))
}

// FindById
// @Tags 员工管理
// @Summary 根据id查询员工内容详情
// @Description 根据id查询员工内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "员工id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /storeStaff/findById/{id} [get]
func (ss *storeStaff) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.StoreStaff.FindById(ctx))
}
