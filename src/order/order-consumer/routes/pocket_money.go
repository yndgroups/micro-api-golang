package routes

import (
	"core/middleware"
	"fmt"
	"order/order-consumer/apis"

	"github.com/gin-gonic/gin"
)

var PocketMoney = &pocketMoney{}

type pocketMoney struct{}

func (w *pocketMoney) InitRoutes(routes *gin.Engine, basePath string) {
	pocketMoney := routes.Group(fmt.Sprintf("%s/pocketMoney", basePath))
	// 查询余额
	pocketMoney.GET("/getBalance", middleware.Authorization("order:pocketMoney:manage"), apis.PocketMoney.GetBalance)
	// 充值记录
	pocketMoney.POST("/findPageList", middleware.Authorization("order:pocketMoney:manage"), apis.PocketMoney.FindPageList)
	// 创建充值订单
	pocketMoney.GET("/create/:amount", middleware.Authorization("order:pocketMoney:manage"), apis.PocketMoney.Create)
	// 详情
	pocketMoney.GET("/findById/:id", middleware.Authorization("order:pocketMoney:manage"), apis.PocketMoney.FindById)

}
