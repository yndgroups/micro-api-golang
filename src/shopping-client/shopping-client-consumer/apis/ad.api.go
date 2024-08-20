package apis

import (
	"net/http"
	"shopping-client/shopping-client-consumer/service"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

type ad struct{}

var Ad = &ad{}

// FindById
// @Tags 广告管理
// @Summary 根据id查询广告内容详情
// @Description 根据id查询广告内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "广告id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /ad/findById/{id} [get]
func (a *ad) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Ad.FindById(ctx))
}

// FindList
// @Tags 广告管理
// @Summary 获取广告信息
// @Description 主要用于获取广告信息
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param type path string  true "广告类型[1:轮播图广告,2:单张图片]"
// @Param position path string  true "广告位置[1:首页,2:列表,3:详情页]"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /ad/findList/{type}/{position} [get]
func (a *ad) FindList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Ad.FindList(ctx))
}
