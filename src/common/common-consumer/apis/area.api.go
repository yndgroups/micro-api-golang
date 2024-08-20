package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var Area = &area{}

type area struct{}

// Create
// @Tags 地区管理
// @Summary 地区添加
// @Description 主要用于用户进行地区添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Area true "地区新增参数"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /area [post]
func (area *area) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Area.Create(ctx))
}

// Update
// @Tags 地区管理
// @Summary 地区修改
// @Description 主要用于用户进行地区修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.Area true "地区参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /area [put]
func (area *area) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Area.Update(ctx))
}

// Delete
// @Tags 地区管理
// @Summary 地区删除
// @Description 主要用于用户进行地区删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "地区id集合"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /area [delete]
func (area *area) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Area.Delete(ctx))
}

// FindPageList
// @Tags 地区管理
// @Summary 地区分页列表查询
// @Description 主要用于地区分页列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[common.AreaParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /area/findAreaPageList [post]
func (area *area) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Area.FindPageList(ctx))
}

// FindListByParentCode
// @Tags 地区管理
// @Summary 根据父编码查询下级列表
// @Description 主要用于父级编码查询下级列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param parentCode query string  true "地区父亲编码"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /area/findAreaListByParentCode [get]
func (area *area) FindListByParentCode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Area.FindListByParentCode(ctx))
}

// FindById
// @Tags 地区管理
// @Summary 根据id地区详情查询
// @Description 主要用于地区详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string  true "地区id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /area/findAreaById/{id} [get]
func (area *area) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Area.FindById(ctx))
}

// FindByAreaCode
// @Tags 地区管理
// @Summary 根据areaCode地区详情查询
// @Description 主要用于地区详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param areaCode query string  true "地区编码"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /area/findAreaByAreaCode [get]
func (area *area) FindByAreaCode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Area.FindByAreaCode(ctx))
}

// FindTree
// @Tags 地区管理
// @Summary 查询省市区三级数据树状结构
// @Description 主要用于地区详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /area/findAreaTree [get]
func (area *area) FindTree(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Area.FindTree(ctx))
}
