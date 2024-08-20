package apis

import (
	"net/http"
	"shopping-client/shopping-client-consumer/service"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Cart = &cart{}

type cart struct{}

// @Tags 购物车
// @Summary 购物车列表
// @Description 主要用于查看购物车列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param addCartParam body shopping_client.Cart true "添加参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /cart/add [post]
func (c *cart) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Cart.Create(ctx))
}

// @Tags 购物车
// @Summary 购物车列表
// @Description 主要用于查看购物车列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body shopping_client.CartPageAuthQuery true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /cart/findPageList [post]
func (c *cart) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Cart.FindPageList(ctx))
}

// @Tags 购物车
// @Summary 根据商品id购删除购物车商品
// @Description 主要用于购物车商品删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /cart/delete [delete]
func (c *cart) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Cart.Delete(ctx))
}

// @Tags 购物车
// @Summary 查询用户的购物车数量
// @Description 主要用于查询用户的购物车数量
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /cart/findCount [get]
func (c *cart) FindCount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Cart.FindCount(ctx))
}

// @Tags 购物车
// @Summary 修改购物车数量
// @Description 主要用于修改购物车数量
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param cartId path string  true "购物车id"
// @Param quantity path string  true "商品数量id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /cart/changeQuantity/{cartId}/{quantity} [get]
func (c *cart) ChangeQuantity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Cart.ChangeQuantity(ctx))
}
