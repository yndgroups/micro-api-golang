package apis

import (
	"net/http"
	"order/order-consumer/service"

	"github.com/gin-gonic/gin"

	_ "core/req"
	_ "core/resp"
)

var FreightTemplatesUndelivered = &freightTemplatesUndelivered{}

type freightTemplatesUndelivered struct{}

// Create
// @Tags 模板管理-不送达
// @Summary 模板不送达添加
// @Description 主要用于用户进行模板不送达添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body order.FreightTemplatesUndelivered true "模板不送达新增参数"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /templatesUndelivered [post]
func (ftu *freightTemplatesUndelivered) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesUndelivered.Create(ctx))
}

// Update
// @Tags 模板管理-不送达
// @Summary 模板不送达修改
// @Description 主要用于用户进行模板不送达修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body order.FreightTemplatesUndelivered true "模板不送达参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /templatesUndelivered [put]
func (ftu *freightTemplatesUndelivered) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesUndelivered.Update(ctx))
}

// Delete
// @Tags 模板管理-不送达
// @Summary 模板不送达删除
// @Description 主要用于用户进行模板不送达删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "模板不送达id集合"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /templatesUndelivered [delete]
func (ftu *freightTemplatesUndelivered) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesUndelivered.Delete(ctx))
}

// FindPageList
// @Tags 模板管理-不送达
// @Summary 模板不送达分页列表查询
// @Description 主要用于模板不送达分页列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[any]  true "查询参数"
// @Success 200 {object} resp.Response{data=resp.Pager{list=[]order.FreightTemplatesUndelivered}}  "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /templatesUndelivered/findPageList [post]
func (ftu *freightTemplatesUndelivered) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesUndelivered.FindPageList(ctx))
}

// FindById
// @Tags 模板管理-不送达
// @Summary 根据id模板不送达详情查询
// @Description 主要用于模板不送达详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string  true "模板不送达id"
// @Success 200 {object} resp.Response{data=order.FreightTemplatesUndelivered} "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /templatesUndelivered/findById/{id} [get]
func (ftu *freightTemplatesUndelivered) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesUndelivered.FindById(ctx))
}
