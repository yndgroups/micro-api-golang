package apis

import (
	"net/http"
	"order/order-consumer/service"

	"github.com/gin-gonic/gin"

	_ "core/req"
	_ "core/resp"
)

var FreightTemplatesFree = &freightTemplatesFree{}

type freightTemplatesFree struct{}

// Create
// @Tags 模板管理-免邮
// @Summary 模板免邮添加
// @Description 主要用于用户进行模板免邮添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body order.FreightTemplatesFree true "模板免邮新增参数"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /templatesFree [post]
func (ftf *freightTemplatesFree) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesFree.Create(ctx))
}

// Update
// @Tags 模板管理-免邮
// @Summary 模板免邮修改
// @Description 主要用于用户进行模板免邮修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body order.FreightTemplatesFree true "模板免邮参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /templatesFree [put]
func (ftf *freightTemplatesFree) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesFree.Update(ctx))
}

// Delete
// @Tags 模板管理-免邮
// @Summary 模板免邮删除
// @Description 主要用于用户进行模板免邮删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "模板免邮id集合"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /templatesFree [delete]
func (ftf *freightTemplatesFree) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesFree.Delete(ctx))
}

// FindPageList
// @Tags 模板管理-免邮
// @Summary 模板免邮分页列表查询
// @Description 主要用于模板免邮分页列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[any] true "查询参数"
// @Success 200 {object} resp.Response{data=resp.Pager{list=[]order.FreightTemplatesFree}} "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /templatesFree/findPageList [post]
func (ftf *freightTemplatesFree) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesFree.FindPageList(ctx))
}

// FindById
// @Tags 模板管理-免邮
// @Summary 根据id模板免邮详情查询
// @Description 主要用于模板免邮详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string  true "模板免邮id"
// @Success 200 {object} resp.Response{data=order.FreightTemplatesFree} "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /templatesFree/findById/{id} [get]
func (ftf *freightTemplatesFree) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.FreightTemplatesFree.FindById(ctx))
}
