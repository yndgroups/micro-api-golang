package routes

import (
	"core/middleware"
	"fmt"
	"shopping-client/shopping-client-consumer/apis"
	_ "shopping-client/shopping-client-consumer/docs"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(env string) (route *gin.Engine, baseApiPath string) {
	routes := gin.Default()
	// 环境模式设置
	switch true {
	case strings.Contains(env, "test"):
		gin.SetMode(gin.TestMode)
	case strings.Contains(env, "prod"):
		gin.SetMode(gin.ReleaseMode)
		// 生产环境捕获异常并让成功不要挂掉
		routes.Use(middleware.Recover)
	default:
		gin.SetMode(gin.TestMode)
		routes.Use(middleware.Recover)
	}

	// api基础路径
	basePath := "/api/shopping/client"

	// 接口文档
	routes.GET(basePath+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(basePath+"/swagger/doc.json")))

	// 商品
	product := routes.Group(fmt.Sprintf("%s/product", basePath))
	product.POST("/findPageList", middleware.Authorization("anonymous"), apis.Product.FindPageList)
	product.GET("/findById/:productId", middleware.Authorization("anonymous"), apis.Product.FindById)
	product.GET("/categoryByParentId/:parentId", middleware.Authorization("anonymous"), apis.ProductCategory.FindCategoryByParentId)         // 商品分类
	product.GET("/categoryTreeByParentId/:parentId", middleware.Authorization("anonymous"), apis.ProductCategory.FindCategoryTreeByParentId) // 商品分类

	// 广告管理
	ad := routes.Group(fmt.Sprintf("%s/ad", basePath))
	ad.GET("/findById/:id", middleware.Authorization("anonymous"), apis.Ad.FindById)
	ad.GET("/findList/:type/:position", middleware.Authorization("anonymous"), apis.Ad.FindList)

	// 购物车
	cart := routes.Group(fmt.Sprintf("%s/cart", basePath))
	cart.POST("/add", middleware.Authorization("shop:c:cart:manage"), apis.Cart.Create)                                     // 添加购物车
	cart.GET("/changeQuantity/:cartId/:quantity", middleware.Authorization("shop:c:cart:manage"), apis.Cart.ChangeQuantity) // 修改购物车数量
	cart.POST("/findPageList", middleware.Authorization("shop:c:cart:manage"), apis.Cart.FindPageList)                      // 购物车商品列表
	cart.GET("/findCount", middleware.Authorization("shop:c:cart:manage"), apis.Cart.FindCount)                             // 查询购物车数量
	cart.DELETE("/delete", middleware.Authorization("shop:c:cart:manage"), apis.Cart.Delete)                                // 清除购物车

	// 商家管理
	business := routes.Group(fmt.Sprintf("%s/business", basePath))
	business.POST("", middleware.Authorization("shop:c:business:add"), apis.Business.Create) // 商家申请入住
	business.PUT("", middleware.Authorization("shop:c:business:edit"), apis.Business.Update)
	business.GET("/findById/:id", middleware.Authorization("shop:c:business:details"), apis.Business.FindById) // 商家详情

	// 商家加盟管理
	franchisee := routes.Group(fmt.Sprintf("%s/franchisee", basePath))
	franchisee.POST("", middleware.Authorization("shop:c:franchisee:add"), apis.BusinessFranchisee.Create)
	franchisee.PUT("", middleware.Authorization("shop:c:franchisee:edit"), apis.BusinessFranchisee.Update)
	franchisee.DELETE("", middleware.Authorization("shop:c:franchisee:del"), apis.BusinessFranchisee.Delete)
	franchisee.POST("/findPageList", middleware.Authorization("shop:c:franchisee:list"), apis.BusinessFranchisee.FindPageList)
	franchisee.GET("/findById/:id", middleware.Authorization("shop:c:franchisee:details"), apis.BusinessFranchisee.FindById)

	// 商圈管理
	businessCircle := routes.Group(fmt.Sprintf("%s/businessCircle", basePath))
	businessCircle.DELETE("", middleware.Authorization("shop:c:businessCircle:del"), apis.BusinessCircle.Delete)
	businessCircle.POST("/findPageList", middleware.Authorization("shop:c:businessCircle:list"), apis.BusinessCircle.FindPageList)
	businessCircle.GET("/findById/:id", middleware.Authorization("shop:c:businessCircle:details"), apis.BusinessCircle.FindById)

	// 门店管理
	businessStore := routes.Group(fmt.Sprintf("%s/businessStore", basePath))
	businessStore.DELETE("", middleware.Authorization("shop:c:businessStore:details"), apis.BusinessStore.Delete)
	businessStore.POST("/findPageList", middleware.Authorization("shop:c:businessStore:details"), apis.BusinessStore.FindPageList)
	businessStore.GET("/findById/:id", middleware.Authorization("shop:c:businessStore:details"), apis.BusinessStore.FindById)

	return routes, basePath
}
