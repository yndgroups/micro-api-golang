package service

import (
	"core/config"
	"core/model"
	"core/resp"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nfnt/resize"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var File = &file{}

type file struct{}

// SingleUpload 多尺寸单个文件上传
func (f *file) SingleUpload(ctx *gin.Context) any {
	appId := ctx.Param("appId")
	if appId == "" {
		return resp.Fail.WithMsg("appId参数获取失败")
	}
	if ctx.Param("sizeType") == "" {
		return resp.Fail.WithMsg("请正确传入尺寸类型参数不正确")
	}
	file, fileGetErr := ctx.FormFile("file")
	if fileGetErr != nil {
		return resp.Fail.WithMsg("文件表单获取失败")
	}
	// 获取文件后缀名
	suffix := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, "."):])
	if !strings.Contains(".jpg,.jpeg,.png", suffix) {
		return resp.Fail.WithMsg("图片文件格式不正确")
	}
	// 创建存储路径
	timePath := time.Now().Format("20060102")
	uuidName := uuid.New()
	storePath := fmt.Sprintf("%s/%s/%s", config.Global.ServerConfig().UploadPath, appId, timePath)
	exists := f.FileOrDirIsExist(storePath)
	if !exists {
		os.MkdirAll(storePath, os.ModePerm)
	}
	fileName := fmt.Sprintf("%s%s", uuidName, suffix)
	filePath := fmt.Sprintf("%s/source-%s", storePath, fileName)
	if ctx.Param("sizeType") == "1" {
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			return resp.Fail.WithMsg("文件上传失败")
		} else {
			reqPath := fmt.Sprintf("/resources/%s/%s/source-%s", appId, timePath, fileName)
			return resp.Success.WithData(reqPath)
		}
	} else if ctx.Param("sizeType") == "2" {
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			return resp.Fail.WithMsg("文件上传失败")
		}
		reqPath := fmt.Sprintf("/resources/%s/%s/source-%s", appId, timePath, fileName)
		if result, err := f.uploadFileMoreSize(filePath, suffix, reqPath); err != nil {
			return resp.Fail.WithMsg(err.Error())
		} else {
			return resp.Success.WithData(result)
		}
	} else {
		return resp.Fail.WithMsg("尺寸类型参数不正确,只能传入1或者2")
	}
}

// MultipleUpload 文件批量上传
func (f *file) MultipleUpload(ctx *gin.Context) any {
	appId := ctx.Param("appId")
	if appId == "" {
		return resp.Fail.WithMsg("appId参数获取失败")
	}
	sizeType := ctx.Param("sizeType")
	if sizeType == "" {
		return resp.Fail.WithMsg("请正确传入id参数")
	} else {
		if !strings.Contains("1,2", sizeType) {
			return resp.Fail.WithMsg("类型参数不正确!只能传入1或者2")
		}
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		return resp.Fail.WithMsg("文件获取失败:" + err.Error())
	}
	// 获取所有图片
	files := form.File["file"]
	// 遍历所有图片
	for _, file := range files {
		// 逐个存
		if err := ctx.SaveUploadedFile(file, file.Filename); err != nil {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
		}
	}
	return resp.Success.WithData(files)
}

// GetList 获取文件列表
func (f *file) GetList(ctx *gin.Context) any {
	return "开发中..."
}

// Delete 文件删除
func (f *file) Delete(ctx *gin.Context) any {
	file := model.File{}
	if err := ctx.ShouldBindJSON(&file); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if file.FilePath == "" {
		return resp.Fail.WithMsg("必须传入要删除的图片路径")
	}
	if file.SizeType == 0 {
		return resp.Fail.WithMsg("必须传入要删除的图片尺寸类型")
	} else {
		if file.SizeType > 2 {
			return resp.Fail.WithMsg("参数只能传1或2")
		}
	}
	if file.SizeType == 1 {
		if _, e := f.fileOrDirRemove(config.Global.ServerConfig().UploadPath + strings.Replace(file.FilePath, "/resources", "", 1)); e != nil {
			return resp.Fail.WithMsg(e.Error())
		} else {
			return resp.Success.WithMsg("删除成功")
		}
	} else {
		var filePaths []string
		path := strings.Replace(file.FilePath, "/resources", "", 1)
		filePaths = append(filePaths, path)
		filePaths = append(filePaths, strings.Replace(path, "source-", "medium-", 1))
		filePaths = append(filePaths, strings.Replace(path, "source-", "thumbnail-", 1))
		filePaths = append(filePaths, strings.Replace(path, "source-", "large-", 1))
		println(filePaths)
		for i := 0; i < len(filePaths); i++ {
			if _, e := f.fileOrDirRemove(config.Global.ServerConfig().UploadPath + filePaths[i]); e != nil {
				if e.Error() == "文件不存在" {
					continue
				} else {
					return resp.Fail.WithDataAndMsg(config.Global.ServerConfig().UploadPath+filePaths[i], e.Error())
				}
			}
		}
		return resp.Success.WithData("删除成功")
	}
}

// fileOrDirRemove 删除文件夹或者文件
func (f *file) fileOrDirRemove(path string) (bool, error) {
	b := f.FileOrDirIsExist(path)
	if b {
		if err := os.Remove(path); err != nil {
			return false, errors.New("文件删除失败")
		}
		return true, nil
	}
	return false, errors.New("文件不存在")
}

// FileOrDirIsExist 判断文件或文件夹是否存在
func (f *file) FileOrDirIsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// uploadFileMoreSize 多文件上传
func (f *file) uploadFileMoreSize(filePath string, suffix string, reqPath string) (any, error) {
	// 读取源文件进行尺寸修改重新存储
	openFile, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("图片件路径读取错误！")
	}
	result := map[string]string{}
	result["source"] = reqPath
	if img, _, err := image.Decode(openFile); err != nil {
		return nil, errors.New("文件读取失败！")
	} else {
		dx := img.Bounds().Dx() // 宽度
		dy := img.Bounds().Dy() // 高度
		options := jpeg.Options{Quality: 50}
		prefix := "source-"
		if dx > 1000 {
			largeImg := resize.Resize(1000, 0, img, resize.NearestNeighbor)
			largeImgPath := strings.ReplaceAll(filePath, prefix, "large-")
			if err != nil {
				return nil, errors.New("图片转码错误！")
			}
			out, err := os.Create(largeImgPath)
			if err != nil {
				return nil, errors.New("图片创建失败！")
			}
			defer out.Close()
			if strings.Contains(".jpg,.jpeg", suffix) {
				jpeg.Encode(out, largeImg, &options)
			} else {
				png.Encode(out, largeImg)
			}
			result["large"] = strings.ReplaceAll(reqPath, prefix, "large-")
		}
		if dx > 400 {
			mediumImg := resize.Resize(400, uint(400*dy/dx), img, resize.NearestNeighbor)
			mediumPath := strings.ReplaceAll(filePath, prefix, "medium-")
			if err != nil {
				return nil, errors.New("图片转码错误！")
			}
			out, err := os.Create(mediumPath)
			if err != nil {
				return nil, errors.New("图片创建失败！")
			}
			defer out.Close()
			if strings.Contains(".jpg,.jpeg", suffix) {
				jpeg.Encode(out, mediumImg, &options)
			} else {
				png.Encode(out, mediumImg)
			}
			result["medium"] = strings.ReplaceAll(reqPath, prefix, "medium-")
		}
		if dx > 100 {
			thumbnailImg := resize.Resize(100, uint(100*dy/dx), img, resize.NearestNeighbor)
			thumbnailImgPath := strings.ReplaceAll(filePath, prefix, "thumbnail-")
			if err != nil {
				return nil, errors.New("文件转码错误！")
			}
			out, err := os.Create(thumbnailImgPath)
			if err != nil {
				return nil, errors.New("文件创建失败！")
			}
			defer out.Close()
			if strings.Contains(".jpg,.jpeg", suffix) {
				jpeg.Encode(out, thumbnailImg, &options)
			} else {
				png.Encode(out, thumbnailImg)
			}
			result["thumbnail"] = strings.ReplaceAll(reqPath, prefix, "thumbnail-")
		}
	}
	return result, nil
}
