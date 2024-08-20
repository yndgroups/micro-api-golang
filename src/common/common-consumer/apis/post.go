package apis

import (
	"common/common-consumer/service"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "core/req"
	_ "core/resp"
)

var Post = &post{}

type post struct{}

// Create
// @Tags 岗位管理
// @Summary 创建岗位添加
// @Description 主要用于岗位添加新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param org body common.Post true "新增岗位添加参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /org [post]
func (p *post) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Post.Create(ctx))
}

// Create
// @Tags 岗位管理
// @Summary 岗位修改
// @Description 主要用于岗位修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param org body common.Post true "新增岗位添加参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /org [put]
func (p *post) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Post.Update(ctx))
}

// Delete
// @Tags 岗位管理
// @Summary 岗位删除
// @Description 岗位删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "地区id集合"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /area [delete]
func (p *post) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Post.Delete(ctx))
}

// FindPageList
// @Tags 地区管理
// @Summary 岗位分页列表查询
// @Description 主要用于岗位分页列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[common.AreaParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /area/findAreaPageList [post]
func (p *post) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Post.FindPageList(ctx))
}
