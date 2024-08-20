package apis

import (
	_ "core/req"
	_ "core/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-admin/shopping-admin-consumer/service"
)

var ProductSpec = &productSpec{}

type productSpec struct{}

// Create
// @Tags 商品规格管理
// @Summary 创建商品规格
// @Description 主要用于商品规格新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param spec body shopping_admin.ProductSpec true "新增商品规格参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /spec [post]
func (ps *productSpec) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductSpec.Create(ctx))
}

// Update
// @Tags 商品规格管理
// @Summary 修改商品规格
// @Description 主要用于商品规格修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param spec body shopping_admin.ProductSpec true "修改商品规格参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /spec [put]
func (ps *productSpec) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductSpec.Update(ctx))
}

// Delete
// @Tags 商品规格管理
// @Summary 删除商品规格
// @Description 主要用于商品规格删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /spec [delete]
func (ps *productSpec) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductSpec.Delete(ctx))
}

// FindPageList
// @Tags 商品规格管理
// @Summary 商品规格分页列表
// @Description 主要用于查询商品规格分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[shopping_admin.ProductSpecParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /spec/findPageList [post]
func (ps *productSpec) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductSpec.FindPageList(ctx))
}

// FindById
// @Tags 商品规格管理
// @Summary 根据id查询商品规格内容详情
// @Description 根据id查询商品规格内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "商品规格id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /spec/findById/{id} [get]
func (ps *productSpec) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductSpec.FindById(ctx))
}

// FindListByProductCategoryId
// @Tags 商品规格管理
// @Summary 根据productCategoryId查询规格列表
// @Description 根据productCategoryId查询规格列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "商品分类id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /spec/findList/{id} [get]
func (ps *productSpec) FindListByProductCategoryId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.ProductSpec.FindListByProductCategoryId(ctx))
}
