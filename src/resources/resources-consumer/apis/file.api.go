package apis

import (
	_ "core/req"
	_ "core/resp"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"resources/resources-consumer/service"

	"github.com/gin-gonic/gin"
)

var File = &file{}

type file struct{}

// SingleUpload
// @Tags 文件上传管理
// @Summary 多尺寸单个文件上传
// @Description 主要用于多尺寸单个文件上传
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param sizeType path string true "上传类型: 1：单尺寸上传 2：多尺寸上传"
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,404,500 {object} resp.Response "请求失败"
// @Router /file/singleUpload/{sizeType} [post]
func (f *file) SingleUpload(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.File.SingleUpload(ctx))
}

// MultipleUpload
// @Tags 文件上传管理
// @Summary 多尺寸多个文件上传
// @Description 主要用于多尺寸单个文件上传
// @Accept multipart/form-data
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param sizeType path string true "上传类型: 1：单尺寸上传 2：多尺寸上传"
// @Param file formData file  true "文件"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 300,400,500 {object} resp.Response "请求失败"
// @Router /file/multipleUpload/{sizeType} [post]
func (f *file) MultipleUpload(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.File.MultipleUpload(ctx))
}

// GetList
// @Tags 文件上传管理
// @Summary 获取文件列表
// @Description 主要用于获取文件列表
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,404,500 {object} resp.Response "请求失败"
// @Router /file/getList [post]
func (f *file) GetList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.File.GetList(ctx))
}

// Delete
// @Tags 文件上传管理
// @Summary 文件删除
// @Description 主要用于文件删除
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param object body model.File true "应用参数"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,404,500 {object} resp.Response "请求失败"
// @Router /file/delete [post]
func (f *file) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.File.Delete(ctx))
}
