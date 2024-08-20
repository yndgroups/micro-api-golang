package apis

import (
	_ "core/req"
	_ "core/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-admin/shopping-admin-consumer/service"
)

var Brand = &brand{}

type brand struct{}

// Create
// @Tags 品牌管理
// @Summary 创建品牌
// @Description 主要用于品牌新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param brand body shopping_admin.Brand true "新增品牌参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /brand [post]
func (b *brand) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Brand.Create(ctx))
}

// Update
// @Tags 品牌管理
// @Summary 修改品牌
// @Description 主要用于品牌修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param brand body shopping_admin.Brand true "修改品牌参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /brand [put]
func (b *brand) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Brand.Update(ctx))
}

// Delete
// @Tags 品牌管理
// @Summary 删除品牌
// @Description 主要用于品牌删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /brand [delete]
func (b *brand) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Brand.Delete(ctx))
}

// FindPageList
// @Tags 品牌管理
// @Summary 品牌分页列表
// @Description 主要用于查询品牌分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[shopping_admin.BrandParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /brand/findPageList [post]
func (b *brand) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Brand.FindPageList(ctx))
}

// FindById
// @Tags 品牌管理
// @Summary 根据id查询品牌内容详情
// @Description 根据id查询品牌内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "品牌id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /brand/findById/{id} [get]
func (b *brand) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Brand.FindById(ctx))
}

// FindListByProductCategoryId
// @Tags 品牌管理
// @Summary 根据商品分类id查询品牌列表
// @Description 根据productCategoryId查询品牌列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "分类id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /brand/findListByProductCategoryId/{id} [get]
func (b *brand) FindListByProductCategoryId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Brand.FindListByProductCategoryId(ctx))
}
