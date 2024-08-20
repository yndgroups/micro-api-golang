package apis

import (
	"net/http"
	"shopping-client/shopping-client-consumer/service"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var BusinessStore = &businessStore{}

type businessStore struct{}

// @Tags 商家门店管理
// @Summary 删除商家门店
// @Description 主要用于商家门店删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "商家门店是删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /businessStore [delete]
func (bs *businessStore) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessStore.Delete(ctx))
}

// @Tags 商家门店管理
// @Summary 商家门店分页列表
// @Description 主要用于查询商家门店分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[shopping_client.BusinessStoreParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /businessStore/findPageList [post]
func (bs *businessStore) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessStore.FindPageList(ctx))
}

// @Tags 商家门店管理
// @Summary 根据id查询商家门店内容详情
// @Description 根据id查询商家门店内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "商家门店id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /businessStore/findById/{id} [get]
func (bs *businessStore) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessStore.FindById(ctx))
}
