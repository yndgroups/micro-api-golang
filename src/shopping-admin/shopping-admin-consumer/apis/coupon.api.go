package apis

import (
	_ "core/req"
	_ "core/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-admin/shopping-admin-consumer/service"
)

var Coupon = &coupon{}

type coupon struct{}

// Create
// @Tags 优惠券管理
// @Summary 创建优惠券
// @Description 主要用于优惠券新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param coupon body shopping_admin.Coupon true "新增优惠券参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /coupon [post]
func (c *coupon) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Coupon.Create(ctx))
}

// Update
// @Tags 优惠券管理
// @Summary 修改优惠券
// @Description 主要用于优惠券修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param coupon body shopping_admin.Coupon true "修改优惠券参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /coupon [put]
func (c *coupon) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Coupon.Update(ctx))
}

// Delete
// @Tags 优惠券管理
// @Summary 删除优惠券
// @Description 主要用于优惠券删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /coupon [delete]
func (c *coupon) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Coupon.Delete(ctx))
}

// FindPageList
// @Tags 优惠券管理
// @Summary 优惠券分页列表
// @Description 主要用于查询优惠券分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[shopping_admin.CouponParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /coupon/findPageList [post]
func (c *coupon) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Coupon.FindPageList(ctx))
}

// FindById
// @Tags 优惠券管理
// @Summary 根据id查询优惠券内容详情
// @Description 根据id查询优惠券内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "优惠券id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /coupon/findById/{id} [get]
func (c *coupon) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Coupon.FindById(ctx))
}
