package apis

import (
	"net/http"
	"shopping-client/shopping-client-consumer/service"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var ProductCategory = &productCategory{}

type productCategory struct{}

// @Tags 商品接口
// @Summary 根据商品父id查询下级分类
// @Description 主要用于根据商品父id查询下级分类
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param parentId path string  true "分类父id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /product/categoryByParentId/{parentId} [get]
func (pc *productCategory) FindCategoryByParentId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductCategory.FindCategoryByParentId(ctx))
}

// @Tags 商品接口
// @Summary 根据商品父id查询所有子分类【树形结构】
// @Description 主要用于根据商品父id查询所有子分类
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param parentId path string  true "分类父id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /product/categoryTreeByParentId/{parentId} [get]
func (pc *productCategory) FindCategoryTreeByParentId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductCategory.FindCategoryTreeByParentId(ctx))
}
