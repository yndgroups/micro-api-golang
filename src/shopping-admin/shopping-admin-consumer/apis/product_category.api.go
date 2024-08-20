package apis

import (
	_ "core/req"
	_ "core/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-admin/shopping-admin-consumer/service"
)

type productCategory struct{}

var ProductCategory = &productCategory{}

// Create
// @Tags 商品分类
// @Summary 创建商品分类
// @Description 主要用于商品分类新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param productCategory body shopping_admin.ProductCategory true "新增商品分类参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /category [post]
func (pc *productCategory) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductCategory.Create(ctx))
}

// Update
// @Tags 商品分类
// @Summary 修改商品分类
// @Description 主要用于商品分类修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param productCategory body shopping_admin.ProductCategory true "修改商品分类参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /category [put]
func (pc *productCategory) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductCategory.Update(ctx))
}

// Delete
// @Tags 商品分类
// @Summary 删除商品分类
// @Description 主要用于商品分类删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /category [delete]
func (pc *productCategory) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductCategory.Delete(ctx))
}

// FindListByParentId
// @Tags 商品分类
// @Summary 根据商品父id查询下级分类
// @Description 主要用于根据商品父id查询下级分类
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param parentId path string  true "分类父id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /category/findListByParentId/{parentId} [get]
func (pc *productCategory) FindListByParentId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductCategory.FindListByParentId(ctx))
}

// FindTreeByParentId
// @Tags 商品分类
// @Summary 根据商品父id查询所有子分类【树形结构】
// @Description 主要用于根据商品父id查询所有子分类
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param parentId path string  true "分类父id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /category/findTreeByParentId/{parentId} [get]
func (pc *productCategory) FindTreeByParentId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductCategory.FindTreeByParentId(ctx))
}
