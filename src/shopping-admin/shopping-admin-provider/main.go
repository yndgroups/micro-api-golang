package main

import (
	"core/app"
	"protobuf/build/shopping_admin"
	"shopping-admin/shopping-admin-provider/service"
)

/**
 * 商品服务生产者
 */
func main() {
	app := app.AppLaunch{}
	app.StartGrpcServer()
	shopping_admin.RegisterProductServiceServer(app.GrpcServer, service.ProductService)
	shopping_admin.RegisterProductCategoryServiceServer(app.GrpcServer, service.ProductCategoryService)
	shopping_admin.RegisterStoreStaffServiceServer(app.GrpcServer, service.StoreStaffService)
	shopping_admin.RegisterProductSpecServiceServer(app.GrpcServer, service.ProductSpecService)
	app.StartTcp()
}
