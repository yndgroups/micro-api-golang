package apis

import (
	"net/http"
	"order/order-consumer/service"

	"github.com/gin-gonic/gin"

	_ "core/req"
	_ "core/resp"
)

var FreightTemplatesRegion = &freightTemplatesRegion{}

type freightTemplatesRegion struct{}

// Create
// @Tags 模板管理-区域运费
// @Summary 模板区域运费添加
// @Description 主要用于用户进行模板区域运费添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body order.FreightTemplatesRegion true "模板区域运费新增参数"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /templatesRegion [post]
func (ftr *freightTemplatesRegion) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesRegion.Create(ctx))
}

// Update
// @Tags 模板管理-区域运费
// @Summary 模板区域运费修改
// @Description 主要用于用户进行模板区域运费修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body order.FreightTemplatesRegion true "模板区域运费参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /templatesRegion [put]
func (ftr *freightTemplatesRegion) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesRegion.Update(ctx))
}

// Delete
// @Tags 模板管理-区域运费
// @Summary 模板区域运费删除
// @Description 主要用于用户进行模板区域运费删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "模板区域运费id集合"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /templatesRegion [delete]
func (ftr *freightTemplatesRegion) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesRegion.Delete(ctx))
}

// FindPageList
// @Tags 模板管理-区域运费
// @Summary 模板区域运费分页列表查询
// @Description 主要用于模板区域运费分页列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[any]  true "查询参数"
// @Success 200 {object} resp.Response{data=resp.Pager{list=[]order.FreightTemplatesRegion}} "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /templatesRegion/findPageList [post]
func (ftr *freightTemplatesRegion) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesRegion.FindPageList(ctx))
}

// FindById
// @Tags 模板管理-区域运费
// @Summary 根据id模板区域运费详情查询
// @Description 主要用于模板区域运费详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string  true "模板区域运费id"
// @Success 200 {object} resp.Response{data=order.FreightTemplatesRegion} "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /templatesRegion/findById/{id} [get]
func (ftr *freightTemplatesRegion) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesRegion.FindById(ctx))
}
