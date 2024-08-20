package middleware

import (
	"core/redis"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Permission 权限值验证中间件
func Permission(powerSign string) gin.HandlerFunc {
	return func(c *gin.Context) {
		powerSignAll, _ := redis.GET.GetString(c.Param("powerKey"))
		// 判断权限是否已经变更
		if powerSign == "" || powerSign == "anonymous" || strings.Contains(powerSignAll, powerSign) {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.JSON(http.StatusOK, gin.H{"msg": "无访问权限,请登陆之后再进行访问", "code": 40103, "data": nil})
			c.Abort()
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
	}
}
