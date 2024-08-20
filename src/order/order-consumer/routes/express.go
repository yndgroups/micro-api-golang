package routes

import (
	"core/middleware"
	"fmt"
	"order/order-consumer/apis"

	"github.com/gin-gonic/gin"
)

var Express = &express{}

type express struct{}

func (o *express) InitRoutes(routes *gin.Engine, basePath string) {

	express := routes.Group(fmt.Sprintf("%s/express", basePath))

	express.POST("/kd100/", middleware.Authorization("order:c:manage"), apis.Order.Create)

}
