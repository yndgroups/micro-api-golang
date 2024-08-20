package main

import (
	"core/app"
	"order/order-provider/service"
	"protobuf/build/order"
)

/**
 * 订单服务提供者
 */
func main() {
	app := app.AppLaunch{}
	app.StartGrpcServer()
	order.RegisterOrderServiceServer(app.GrpcServer, service.OrderService)
	order.RegisterFreightTemplatesServiceServer(app.GrpcServer, service.FreightTemplatesService)
	order.RegisterFreightTemplatesFreeServiceServer(app.GrpcServer, service.FreightTemplatesFreeService)
	order.RegisterFreightTemplatesUndeliveredServiceServer(app.GrpcServer, service.FreightTemplatesUndeliveredService)
	order.RegisterPocketMoneyServiceServer(app.GrpcServer, service.PocketMoneyService)
	app.StartTcp()
}
