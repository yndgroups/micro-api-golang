package apis

import (
	_ "core/req"
	_ "core/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-admin/shopping-admin-consumer/service"
)

var Store = &store{}

type store struct{}

// Create
// @Tags 店铺管理
// @Summary 创建店铺
// @Description 主要用于店铺新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param store body shopping_admin.Store true "新增店铺参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /store [post]
func (s *store) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Store.Create(ctx))
}

// Update
// @Tags 店铺管理
// @Summary 修改店铺
// @Description 主要用于店铺修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param store body shopping_admin.Store true "修改店铺参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /store [put]
func (s *store) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Store.Update(ctx))
}

// Delete
// @Tags 店铺管理
// @Summary 删除店铺
// @Description 主要用于店铺删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /store [delete]
func (s *store) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Store.Delete(ctx))
}

// FindPageList
// @Tags 店铺管理
// @Summary 店铺分页列表
// @Description 主要用于查询店铺分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[any] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /store/findPageList [post]
func (s *store) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Store.FindPageList(ctx))
}

// FindById
// @Tags 店铺管理
// @Summary 根据id查询店铺内容详情
// @Description 根据id查询店铺内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "店铺id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /store/findById/{id} [get]
func (s *store) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Store.FindById(ctx))
}
