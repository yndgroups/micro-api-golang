package routes

import (
	"core/middleware"
	"fmt"
	"order/order-consumer/apis"

	"github.com/gin-gonic/gin"
)

var Pay = &pay{}

type pay struct{}

// InitRoutes 支付
func (t *pay) InitRoutes(routes *gin.Engine, basePath string) {

	pay := routes.Group(fmt.Sprintf("%s/pay", basePath))
	// ====== 修改支付密码 ======
	pay.POST("/updatePayPwd", middleware.Authorization("order:pay:updatePayPwd"), apis.Pay.UpdateUserPayPwd)

}
