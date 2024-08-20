package apis

import (
	"net/http"
	"shopping-client/shopping-client-consumer/service"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Business = &business{}

type business struct{}

// @Tags 商家管理
// @Summary 创建商家
// @Description 主要用于商家新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param business body shopping_client.Business true "新增商家参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /business [post]
func (b *business) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Business.Create(ctx))
}

// @Tags 商家管理
// @Summary 修改商家
// @Description 主要用于商家修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param business body shopping_client.Business true "修改商家参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /business [put]
func (b *business) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Business.Update(ctx))
}

// @Tags 商家管理
// @Summary 删除商家
// @Description 主要用于商家删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /business [delete]
func (b *business) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Business.Delete(ctx))
}

// @Tags 商家管理
// @Summary 商家分页列表
// @Description 主要用于查询商家分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[shopping_client.BusinessParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /business/findPageList [post]
func (b *business) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Business.FindPageList(ctx))
}

// @Tags 商家管理
// @Summary 根据id查询商家内容详情
// @Description 根据id查询商家内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "商家id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /business/findById/{id} [get]
func (b *business) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Business.FindById(ctx))
}
