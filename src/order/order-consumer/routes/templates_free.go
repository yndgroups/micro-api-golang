package routes

import (
	"core/middleware"
	"fmt"
	"order/order-consumer/apis"

	"github.com/gin-gonic/gin"
)

var TemplatesFree = &templatesFree{}

type templatesFree struct{}

// InitRoutes 免邮-模板管理
func (t *templatesFree) InitRoutes(routes *gin.Engine, basePath string) {
	templatesFree := routes.Group(fmt.Sprintf("%s/templatesFree", basePath))
	templatesFree.POST("", middleware.Authorization("order:templatesFree:add"), apis.FreightTemplatesFree.Create)
	templatesFree.PUT("", middleware.Authorization("order:templatesFree:edit"), apis.FreightTemplatesFree.Update)
	templatesFree.DELETE("", middleware.Authorization("order:templatesFree:del"), apis.FreightTemplatesFree.Delete)
	templatesFree.POST("/findPageList", middleware.Authorization("order:templatesFree:list"), apis.FreightTemplatesFree.FindPageList)
	templatesFree.GET("/findById/:id", middleware.Authorization("order:templatesFree:details"), apis.FreightTemplatesFree.FindById)
}
