package main

import (
	"common/common-provider/service"
	"core/app"
	"protobuf/build/common"
)

func main() {
	app := app.AppLaunch{}
	app.StartGrpcServer()
	// 注册服务
	common.RegisterUserServiceServer(app.GrpcServer, service.UserService)
	common.RegisterRoleServiceServer(app.GrpcServer, service.RoleService)
	common.RegisterMenuServiceServer(app.GrpcServer, service.MenuService)
	common.RegisterAppServiceServer(app.GrpcServer, service.AppService)
	common.RegisterAppModuleServiceServer(app.GrpcServer, service.AppModuleService)
	common.RegisterAgreementServiceServer(app.GrpcServer, service.AgreementService)
	common.RegisterUserAddressServiceServer(app.GrpcServer, service.UserAddressService)
	common.RegisterAreaServiceServer(app.GrpcServer, service.AreaService)
	common.RegisterConfigServiceServer(app.GrpcServer, service.ConfigService)
	common.RegisterDictServiceServer(app.GrpcServer, service.DictService)
	common.RegisterOrgServiceServer(app.GrpcServer, service.OrgService)
	common.RegisterAuthServer(app.GrpcServer, service.AuthService)
	common.RegisterDepartmentServiceServer(app.GrpcServer, service.DepartmentService)
	common.RegisterPostServiceServer(app.GrpcServer, service.PostService)
	common.RegisterSysFuncServiceServer(app.GrpcServer, service.SysFuncService)
	common.RegisterUserDetailServiceServer(app.GrpcServer, service.UserDetailService)
	app.StartTcp()
}
