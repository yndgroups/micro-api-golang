package apis

import (
	_ "core/req"
	_ "core/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-admin/shopping-admin-consumer/service"
)

var Express = &express{}

type express struct{}

// Create
// @Tags 快递管理
// @Summary 创建快递
// @Description 主要用于快递新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param express body shopping_admin.Express true "新增快递参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /express [post]
func (e *express) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Express.Create(ctx))
}

// Update
// @Tags 快递管理
// @Summary 修改快递
// @Description 主要用于快递修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param express body shopping_admin.Express true "修改快递参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /express [put]
func (e *express) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Express.Update(ctx))
}

// Delete
// @Tags 快递管理
// @Summary 删除快递
// @Description 主要用于快递删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /express [delete]
func (e *express) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Express.Delete(ctx))
}

// FindPageList
// @Tags 快递管理
// @Summary 快递分页列表
// @Description 主要用于查询快递分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[shopping_admin.ExpressParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /express/findPageList [post]
func (e *express) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Express.FindPageList(ctx))
}

// FindById
// @Tags 快递管理
// @Summary 根据id查询快递内容详情
// @Description 根据id查询快递内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "快递id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /express/findById/{id} [get]
func (e *express) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Express.FindById(ctx))
}
