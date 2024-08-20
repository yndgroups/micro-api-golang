package apis

import (
	"net/http"
	"shopping-client/shopping-client-consumer/service"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var BusinessFranchisee = &businessFranchisee{}

type businessFranchisee struct{}

// Create
// @Tags 商家入住管理
// @Summary 创建商家入住
// @Description 主要用于商家入住新增
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param franchisee body shopping_client.BusinessFranchisee true "新增商家入住参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /franchisee [post]
func (bf *businessFranchisee) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessFranchisee.Create(ctx))
}

// Update
// @Tags 商家入住管理
// @Summary 修改商家入住
// @Description 主要用于商家入住修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param franchisee body shopping_client.BusinessFranchisee true "修改商家入住参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /franchisee [put]
func (bf *businessFranchisee) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessFranchisee.Update(ctx))
}

// Delete
// @Tags 商家入住管理
// @Summary 删除商家入住
// @Description 主要用于商家入住删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids true "删除参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /franchisee [delete]
func (bf *businessFranchisee) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessFranchisee.Delete(ctx))
}

// FindPageList
// @Tags 商家入住管理
// @Summary 商家入住分页列表
// @Description 主要用于查询商家入住分页列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param pageInfo body req.Request[shopping_client.BusinessFranchiseeParam] true "查询参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /franchisee/findPageList [post]
func (bf *businessFranchisee) FindPageList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessFranchisee.FindPageList(ctx))
}

// FindById
// @Tags 商家入住管理
// @Summary 根据id查询商家入住内容详情
// @Description 根据id查询商家入住内容详情
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string true "商家入住id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /franchisee/findById/{id} [get]
func (bf *businessFranchisee) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.BusinessFranchisee.FindById(ctx))
}
