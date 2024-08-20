package routes

import (
	"core/middleware"
	"fmt"
	"order/order-consumer/apis"

	"github.com/gin-gonic/gin"
)

var TemplatesRegion = &templatesRegion{}

type templatesRegion struct{}

// InitRoutes 区域运费-模板管理
func (tr *templatesRegion) InitRoutes(routes *gin.Engine, basePath string) {
	templatesRegion := routes.Group(fmt.Sprintf("%s/templatesRegion", basePath))
	templatesRegion.POST("", middleware.Authorization("order:templatesRegion:add"), apis.FreightTemplatesRegion.Create)
	templatesRegion.PUT("", middleware.Authorization("order:templatesRegion:edit"), apis.FreightTemplatesRegion.Update)
	templatesRegion.DELETE("", middleware.Authorization("order:templatesRegion:del"), apis.FreightTemplatesRegion.Delete)
	templatesRegion.POST("/findPageList", middleware.Authorization("order:templatesRegion:list"), apis.FreightTemplatesRegion.FindPageList)
	templatesRegion.GET("/findById/:id", middleware.Authorization("order:templatesRegion:details"), apis.FreightTemplatesRegion.FindById)
}
