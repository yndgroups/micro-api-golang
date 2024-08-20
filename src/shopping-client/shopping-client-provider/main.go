package main

import (
	"core/app"
	"protobuf/build/shopping_client"

	"shopping-client/shopping-client-provider/service"
)

/**
 * 商品服务生产者
 */
func main() {
	app := app.AppLaunch{}
	app.StartGrpcServer()
	shopping_client.RegisterProductServiceServer(app.GrpcServer, service.ProductService)
	shopping_client.RegisterStoreStaffServiceServer(app.GrpcServer, service.StoreStaffService)
	shopping_client.RegisterAdServiceServer(app.GrpcServer, service.AdService)
	shopping_client.RegisterProductCategoryServiceServer(app.GrpcServer, service.ProductCategorySerice)
	shopping_client.RegisterCartServiceServer(app.GrpcServer, service.CartService)
	app.StartTcp()
}
