package apis

import (
	"net/http"
	"shopping-client/shopping-client-consumer/service"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var BusinessCircle = businessCircle{}

type businessCircle struct{}

// @Tags 商圈管理
// @Summary 删除商圈
// @Description 主要用于商圈删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "商圈是删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /businessCircle [delete]
func (bc *businessCircle) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessCircle.Delete(ctx))
}

// @Tags 商圈管理
// @Summary 商圈分页列表
// @Description 主要用于查询商圈分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[shopping_client.BusinessCircleParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /businessCircle/findPageList [post]
func (bc *businessCircle) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessCircle.FindPageList(ctx))
}

// @Tags 商圈管理
// @Summary 根据id查询商圈内容详情
// @Description 根据id查询商圈内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "商圈id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /businessCircle/findById/{id} [get]
func (bc *businessCircle) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessCircle.FindById(ctx))
}
