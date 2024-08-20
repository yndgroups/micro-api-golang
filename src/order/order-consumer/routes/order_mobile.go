package routes

import (
	"core/middleware"
	"fmt"
	"order/order-consumer/apis"

	"github.com/gin-gonic/gin"
)

var OrderMobile = &orderMobile{}

type orderMobile struct{}

// InitRoutes 订单移动端版
func (t *orderMobile) InitRoutes(routes *gin.Engine, basePath string) {

	orderMobile := routes.Group(fmt.Sprintf("%s/order/m", basePath))

	orderMobile.DELETE("", middleware.Authorization("order:c:manage"), apis.OrderMobile.Delete)

	orderMobile.POST("/findPageList", middleware.Authorization("order:c:manage"), apis.OrderMobile.FindPageList)
}
