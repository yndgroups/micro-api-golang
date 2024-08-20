package apis

import (
	"net/http"
	"shopping-client/shopping-client-consumer/service"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Product = &product{}

type product struct{}

// FindById
// @Tags 商品接口
// @Summary 商品详情
// @Description 主要用于商品详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param productId path string  true "商品id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /product/findById/{productId} [get]
func (p *product) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Product.FindById(ctx))
}

// FindPageList
// @Tags 商品接口
// @Summary 商品分页列表
// @Description 主要用于商品分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body shopping_client.ProductPageAuthQuery true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /product/findPageList [post]
func (p *product) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Product.FindPageList(ctx))
}
