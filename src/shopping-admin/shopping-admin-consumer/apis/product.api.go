package apis

import (
	_ "core/model"
	_ "core/req"
	_ "core/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-admin/shopping-admin-consumer/service"
)

type product struct{}

var Product = &product{}

// Create
// @Tags 商品管理
// @Summary 创建商品
// @Description 主要用于商品新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param product body shopping_admin.Product true "新增商品参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /product [post]
func (p *product) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Product.Create(ctx))
}

// Update
// @Tags 商品管理
// @Summary 修改商品
// @Description 主要用于商品修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param product body shopping_admin.Product true "修改商品参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /product [put]
func (p *product) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Product.Update(ctx))
}

// Delete
// @Tags 商品管理
// @Summary 删除商品
// @Description 主要用于商品删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /product [delete]
func (p *product) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Product.Delete(ctx))
}

// FindById
// @Tags 商品管理
// @Summary 商品详情
// @Description 主要用于商品详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param productId path string  true "商品id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /product/findById/{productId} [get]
func (p *product) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Product.FindById(ctx))
}

// FindPageList
// @Tags 商品管理
// @Summary 商品分页列表
// @Description 主要用于商品分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body req.Request[any] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /product/findPageList [post]
func (p *product) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Product.FindPageList(ctx))
}

// FindListByProductIds [提供给订单服务]
// @Tags 商品管理
// @Summary 根据商品id查询商品列表[提供给订单服务]
// @Description 主要用于根据商品id查询商品列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param params body []model.ProductParam true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /product/findListByIds [post]
func (p *product) FindListByProductIds(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Product.FindListByProductIds(ctx))
}
