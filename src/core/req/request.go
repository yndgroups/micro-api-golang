package req

import (
	"encoding/json"
	"fmt"
	"protobuf/build/global"
	"reflect"
	"unsafe"

	"github.com/gin-gonic/gin"
)

/*
	@description Request 请求参数统一封装
	@author yangDaqiong
 	@date  2022-01-02
*/

type Request[T any] struct {
	// 分页索引
	PageIndex int64 `json:"pageIndex" validate:"required,min=1" example:"1"`
	// 分页长度
	PageSize int64 `json:"pageSize" validate:"required,min=10,max=100" example:"10"`
	// 请求查询参数
	Params T `json:"params"`
} // @name Request 请求参数封装

func CreateAuthDTO[T any, S any](req Request[T], authParams S) (S, error) {
	println("ValueOf ===================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	authParam := reflect.ValueOf(&authParams)
	println("for  ===================>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", authParam.Kind(), reflect.Ptr)
	if authParam.Kind() == reflect.Ptr {
		println(authParam.Kind() == reflect.Ptr)
		elem := authParam.Elem()
		pageIndex := elem.FieldByName("PageIndex")
		pageSize := elem.FieldByName("PageSize")
		if pageIndex.Kind() == reflect.Int64 && pageSize.Kind() == reflect.Int64 {
			var index int64 = (req.PageIndex - 1) * req.PageSize
			if index > 100 {
				index = 100
			}
			*(*int64)(unsafe.Pointer(pageIndex.Addr().Pointer())) = index
			*(*int64)(unsafe.Pointer(pageSize.Addr().Pointer())) = req.PageSize
		}

		params := elem.FieldByName("Params")
		println(params.Kind() == reflect.Struct, ">>>>>>>>> Struct  >>>>>>>>>>>>>>>>>>>>>>>>>")

		// 无法被设置通常代表其非指针
		if params.Kind() == reflect.Ptr {
			println(fmt.Sprintf("%T", params.Type()))
			*(*any)(unsafe.Pointer(params.Addr().Pointer())) = reflect.ValueOf(req.Params)
			t := authParam.Elem()
			for i := 0; i < t.NumField(); i++ {
				// t.Field(i).Set()
				fieldInfo := t.Type().Field(i)
				tag := fieldInfo.Tag
				name := tag.Get("json")
				println(fmt.Sprintf("i >> %v", name))
			}
		}
	}
	println(fmt.Sprintf("authParams = %v", authParams))
	return authParams, nil
}

// 结构体转JSON
func CreateAuthParams[T any, S any](req Request[T], authParams S, ctx *gin.Context) S {
	if req.PageSize > 100 {
		req.PageSize = 100
	}
	if req.PageIndex > 0 {
		req.PageIndex = (req.PageIndex - 1) * req.PageSize
	} else {
		req.PageIndex = 0
	}
	reqJson, _ := json.Marshal(struct {
		PageIndex int64
		PageSize  int64
		Params    T
		Auth      global.Auth `json:"auth"`
	}{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		Params:    req.Params,
		Auth: global.Auth{
			UserId: ctx.Param("requestUserId"),
			AppId:  ctx.Param("appId"),
		},
	})
	json.Unmarshal([]byte(reqJson), &authParams)
	// 组装权限参数
	return authParams
}
