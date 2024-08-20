package routes

import (
	"core/middleware"
	"fmt"
	"order/order-consumer/apis"

	"github.com/gin-gonic/gin"
)

var TemplatesUndelivered = &templatesUndelivered{}

type templatesUndelivered struct{}

// InitRoutes 不送达-模板管理
func (t *templatesUndelivered) InitRoutes(routes *gin.Engine, basePath string) {
	templatesUndelivered := routes.Group(fmt.Sprintf("%s/templatesUndelivered", basePath))
	templatesUndelivered.POST("", middleware.Authorization("order:templatesUndelivered:add"), apis.FreightTemplatesUndelivered.Create)
	templatesUndelivered.PUT("", middleware.Authorization("order:templatesUndelivered:edit"), apis.FreightTemplatesUndelivered.Update)
	templatesUndelivered.DELETE("", middleware.Authorization("order:templatesUndelivered:del"), apis.FreightTemplatesUndelivered.Delete)
	templatesUndelivered.POST("/findPageList", middleware.Authorization("order:templatesUndelivered:list"), apis.FreightTemplatesUndelivered.FindPageList)
	templatesUndelivered.GET("/findById/:id", middleware.Authorization("order:templatesUndelivered:details"), apis.FreightTemplatesUndelivered.FindById)

}
