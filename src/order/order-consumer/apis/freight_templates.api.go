package apis

import (
	"net/http"
	"order/order-consumer/service"

	"github.com/gin-gonic/gin"

	_ "core/req"
	_ "core/resp"
)

var FreightTemplates = &freightTemplates{}

type freightTemplates struct{}

// CreateFreightTemplates
// @Tags 模板管理-运费模板
// @Summary 模板添加
// @Description 主要用于用户进行模板添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body order.FreightTemplates true "模板新增参数"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /templates [post]
func (ft *freightTemplates) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplates.Create(ctx))
}

// UpdateFreightTemplates
// @Tags 模板管理-运费模板
// @Summary 模板修改
// @Description 主要用于用户进行模板修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body order.FreightTemplates true "模板参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /templates [put]
func (ft *freightTemplates) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplates.Update(ctx))
}

// DeleteFreightTemplates
// @Tags 模板管理-运费模板
// @Summary 模板删除
// @Description 主要用于用户进行模板删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "模板id集合"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /templates [delete]
func (ft *freightTemplates) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplates.Delete(ctx))
}

// FindPageList
// @Tags 模板管理-运费模板
// @Summary 模板分页列表查询
// @Description 主要用于模板分页列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body order.FreightTemplatesPageAuthQuery true "查询参数"
// @Success 200 {object} resp.Response{data=resp.Pager{list=[]order.FreightTemplates}} "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /templates/findPageList [post]
func (ft *freightTemplates) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplates.FindPageList(ctx))
}

// FindById
// @Tags 模板管理-运费模板
// @Summary 根据id模板详情查询
// @Description 主要用于模板详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string  true "模板id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /templates/findById/{id} [get]
func (ft *freightTemplates) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplates.FindById(ctx))
}
