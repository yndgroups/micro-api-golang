package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 全局异常处理
func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// debug.PrintStack()
			println(fmt.Sprintf("%v", ctx.Errors))
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  errInfoHandle(err),
				"data": nil,
			})
			ctx.Abort()
		}
	}()
	ctx.Next()
}

// 统一处理信息
func errInfoHandle(r interface{}) string {
	var errInfo string
	switch v := r.(type) {
	case error:
		errInfo = v.Error()
	default:
		errInfo = r.(string)
	}
	return Msg(errInfo)
}

func Msg(errStr string) string {
	switch {
	// 无效的内存地址或无指针取消引用
	case strings.Contains(errStr, "invalid memory address or nil pointer dereference"):
		return "服务调用异常"
	}
	return errStr
}
