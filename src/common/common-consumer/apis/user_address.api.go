package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	"github.com/gin-gonic/gin"
)

var UserAddress = &userAddress{}

type userAddress struct{}

// Create
// @Tags 收货地址管理
// @Summary 收货地址添加
// @Description 主要用于用户进行收货地址添加
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.UserAddress true "收货地址新增参数"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /userAddress [post]
func (ua *userAddress) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.UserAddress.Create(ctx))
}

// Update
// @Tags 收货地址管理
// @Summary 收货地址修改
// @Description 主要用于用户进行收货地址修改
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body common.UserAddress true "收货地址参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /userAddress [put]
func (ua *userAddress) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.UserAddress.Update(ctx))
}

// Delete
// @Tags 收货地址管理
// @Summary 收货地址删除
// @Description 主要用于用户进行收货地址删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param ids body model.Ids  true "收货地址id集合"
// @Success 200 {object} resp.Response 成功后返回数据结构
// @Failure 400,500 {object} resp.Response 失败后返回数据结构
// @Router /userAddress [delete]
func (ua *userAddress) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.UserAddress.Delete(ctx))
}

// FindList
// @Tags 收货地址管理
// @Summary 收货地址列表查询
// @Description 主要用于收货地址分页列表查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response{data=[]common.UserAddress} "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /userAddress/findUserAddressList [get]
func (ua *userAddress) FindList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.UserAddress.FindList(ctx))
}

// FindById
// @Tags 收货地址管理
// @Summary 根据id收货地址详情查询
// @Description 主要用于收货地址详情查询
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param id path string  true "地区id"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object}  resp.Response "请求异常"
// @Router /userAddress/findById/{id} [get]
func (ua *userAddress) FindById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.UserAddress.FindById(ctx))
}
