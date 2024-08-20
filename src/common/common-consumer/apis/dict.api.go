package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Dict = &dict{}

type dict struct{}

// Create
// @Tags 字典管理
// @Summary 字典添加
// @Description 主要用于用户进行字典添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Dict true "字典新增"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /dict [post] 字典添加
func (d *dict) Create(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, service.Dict.Create(ctx))
}

// Update
// @Tags 字典管理
// @Summary 字典修改
// @Description 主要用于用户进行字典修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Dict true "登录对象参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /dict [put] 字典添加
func (d *dict) Update(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, service.Dict.Update(ctx))
}

// Delete
// @Tags 字典管理
// @Summary 字典删除
// @Description 主要用于用户进行字典删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "字典id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /dict [delete] 字典删除
func (d *dict) Delete(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, service.Dict.Delete(ctx))
}

// FindById
// @Tags 字典管理
// @Summary 根据字典id查询字典列表
// @Description 主要用于字典列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string  true "字典id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /dict/findById/{id} [get]
func (d *dict) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Dict.FindById(ctx))
}

// FindListByParentId
// @Tags 字典管理
// @Summary 根据字典父id查询字典列表
// @Description 主要用于字典列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param parentId path string  true "字典父亲id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /dict/findDictListByParentId/{parentId} [get]
func (d *dict) FindListByParentId(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Dict.FindListByParentId(ctx))
}
