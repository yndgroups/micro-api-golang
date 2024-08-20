package routes

import (
	"core/middleware"
	"fmt"
	"shopping-admin/shopping-admin-consumer/apis"
	_ "shopping-admin/shopping-admin-consumer/docs"
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
	}

	// api基础路径
	basePath := "/api/shopping/admin"

	// 接口文档
	routes.GET(basePath+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(basePath+"/swagger/doc.json")))

	// 商家管理
	business := routes.Group(fmt.Sprintf("%s/business", basePath))
	business.POST("", apis.Business.Create)
	business.PUT("", apis.Business.Update)
	business.DELETE("", apis.Business.Delete)
	business.POST("/findPageList", apis.Business.FindPageList)
	business.GET("/findById/:id", apis.Business.FindById)

	// 优惠券管理
	coupon := routes.Group(fmt.Sprintf("%s/coupon", basePath))
	coupon.POST("", apis.Coupon.Create)
	coupon.PUT("", apis.Coupon.Update)
	coupon.DELETE("", apis.Coupon.Delete)
	coupon.POST("/findPageList", apis.Coupon.FindPageList)
	coupon.GET("/findById/:id", apis.Coupon.FindById)

	// 商家加盟管理
	franchisee := routes.Group(fmt.Sprintf("%s/franchisee", basePath))
	franchisee.POST("", apis.BusinessFranchisee.Create)
	franchisee.PUT("", apis.BusinessFranchisee.Update)
	franchisee.DELETE("", apis.BusinessFranchisee.Delete)
	franchisee.POST("/findPageList", apis.BusinessFranchisee.FindPageList)
	franchisee.GET("/findById/:id", apis.BusinessFranchisee.FindById)

	// 广告管理
	ad := routes.Group(fmt.Sprintf("%s/ad", basePath))
	ad.POST("", apis.Ad.Create)
	ad.PUT("", apis.Ad.Update)
	ad.DELETE("", apis.Ad.Delete)
	ad.POST("/findPageList", apis.Ad.FindPageList)
	ad.GET("/findById/:id", apis.Ad.FindById)

	// 商圈管理
	businessCircle := routes.Group(fmt.Sprintf("%s/businessCircle", basePath))
	businessCircle.DELETE("", middleware.Authorization("shop:businessCircle:del"), apis.BusinessCircle.Delete)
	businessCircle.POST("/findPageList", middleware.Authorization("shop:businessCircle:list"), apis.BusinessCircle.FindPageList)
	businessCircle.GET("/findById/:id", middleware.Authorization("shop:businessCircle:details"), apis.BusinessCircle.FindById)

	// 门店管理
	businessStore := routes.Group(fmt.Sprintf("%s/businessStore", basePath))
	businessStore.DELETE("", middleware.Authorization("shop:businessStore:add"), apis.BusinessStore.Delete)
	businessStore.POST("/findPageList", middleware.Authorization("shop:businessStore:list"), apis.BusinessStore.FindPageList)
	businessStore.GET("/findById/:id", middleware.Authorization("shop:businessStore:details"), apis.BusinessStore.FindById)

	// 商品
	product := routes.Group(fmt.Sprintf("%s/product", basePath))
	product.POST("", middleware.Authorization("shop:product:add"), apis.Product.Create)
	product.PUT("", middleware.Authorization("shop:product:edit"), apis.Product.Update)
	product.DELETE("", middleware.Authorization("shop:product:del"), apis.Product.Delete)
	product.POST("/findPageList", middleware.Authorization("shop:product:list"), apis.Product.FindPageList)
	product.GET("/findById/:productId", middleware.Authorization("shop:product:list"), apis.Product.FindById)
	// 提供给订单查询使用
	product.POST("/findListByIds", middleware.Authorization("shop:product:list"), apis.Product.FindListByProductIds)

	// 商品分类
	category := routes.Group(fmt.Sprintf("%s/category", basePath))
	category.POST("", middleware.Authorization("shop:category:add"), apis.ProductCategory.Create)
	category.PUT("", middleware.Authorization("shop:category:edit"), apis.ProductCategory.Update)
	category.DELETE("", middleware.Authorization("shop:category:del"), apis.ProductCategory.Delete)
	category.GET("/findListByParentId/:parentId", middleware.Authorization("shop:category:list"), apis.ProductCategory.FindListByParentId)
	category.GET("/findTreeByParentId/:parentId", middleware.Authorization("shop:category:list"), apis.ProductCategory.FindTreeByParentId)

	// 品牌管理
	brand := routes.Group(fmt.Sprintf("%s/brand", basePath))
	brand.POST("", middleware.Authorization("shop:brand:add"), apis.Brand.Create)
	brand.PUT("", middleware.Authorization("shop:brand:edit"), apis.Brand.Update)
	brand.DELETE("", middleware.Authorization("shop:brand:del"), apis.Brand.Delete)
	brand.POST("/findPageList", middleware.Authorization("shop:brand:list"), apis.Brand.FindPageList)
	brand.GET("/findListByProductCategoryId/:id", middleware.Authorization("shop:brand:list"), apis.Brand.FindListByProductCategoryId)
	brand.GET("/findById/:id", middleware.Authorization("shop:brand:list"), apis.Brand.FindById)

	// 商品规格管理
	productSpec := routes.Group(fmt.Sprintf("%s/spec", basePath))
	productSpec.POST("", middleware.Authorization("shop:productSpec:add"), apis.ProductSpec.Create)
	productSpec.PUT("", middleware.Authorization("shop:productSpec:edit"), apis.ProductSpec.Update)
	productSpec.DELETE("", middleware.Authorization("shop:productSpec:del"), apis.ProductSpec.Delete)
	productSpec.POST("/findPageList", middleware.Authorization("shop:productSpec:list"), apis.ProductSpec.FindPageList)
	productSpec.GET("/findById/:id", middleware.Authorization("shop:productSpec:list"), apis.ProductSpec.FindById)
	productSpec.GET("/findList/:id", middleware.Authorization("shop:productSpec:list"), apis.ProductSpec.FindListByProductCategoryId)

	// 快递管理
	express := routes.Group(fmt.Sprintf("%s/express", basePath))
	express.POST("", middleware.Authorization("shop:express:add"), apis.Express.Create)
	express.PUT("", middleware.Authorization("shop:express:edit"), apis.Express.Update)
	express.DELETE("", middleware.Authorization("shop:express:del"), apis.Express.Delete)
	express.POST("/findPageList", middleware.Authorization("shop:express:list"), apis.Express.FindPageList)
	express.GET("/findById/:id", middleware.Authorization("shop:express:list"), apis.Express.FindById)

	// 店铺管理
	store := routes.Group(fmt.Sprintf("%s/store", basePath))
	store.POST("", middleware.Authorization("shop:store:add"), apis.Store.Create)
	store.PUT("", middleware.Authorization("shop:store:edit"), apis.Store.Update)
	store.DELETE("", middleware.Authorization("shop:store:del"), apis.Store.Delete)
	store.POST("/findPageList", middleware.Authorization("shop:store:list"), apis.Store.FindPageList)
	store.GET("/findById/:id", middleware.Authorization("shop:store:details"), apis.Store.FindById)

	// 店铺员工管理
	storeStaff := routes.Group(fmt.Sprintf("%s/storeStaff", basePath))
	storeStaff.POST("", middleware.Authorization("shop:storeStaff:add"), apis.StoreStaff.Create)
	storeStaff.PUT("", middleware.Authorization("shop:storeStaff:edit"), apis.StoreStaff.Update)
	storeStaff.DELETE("", middleware.Authorization("shop:storeStaff:del"), apis.StoreStaff.Delete)
	storeStaff.POST("/findPageList", middleware.Authorization("shop:storeStaff:list"), apis.StoreStaff.FindPageList)
	storeStaff.GET("/findById/:id", middleware.Authorization("shop:storeStaff:details"), apis.StoreStaff.FindById)

	return routes, basePath
}
