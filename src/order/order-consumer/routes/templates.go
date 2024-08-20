package routes

import (
	"core/middleware"
	"fmt"
	"order/order-consumer/apis"

	"github.com/gin-gonic/gin"
)

var Templates = &templates{}

type templates struct{}

// InitRoutes 运费-模板管理
func (t *templates) InitRoutes(routes *gin.Engine, basePath string) {
	templates := routes.Group(fmt.Sprintf("%s/templates", basePath))
	templates.POST("", middleware.Authorization("order:templates:add"), apis.FreightTemplates.Create)
	templates.PUT("", middleware.Authorization("order:templates:edit"), apis.FreightTemplates.Update)
	templates.DELETE("", middleware.Authorization("order:templates:del"), apis.FreightTemplates.Delete)
	templates.POST("/findPageList", middleware.Authorization("order:templates:list"), apis.FreightTemplates.FindPageList)
	templates.GET("/findById/:id", middleware.Authorization("order:templates:details"), apis.FreightTemplates.FindById)

}
