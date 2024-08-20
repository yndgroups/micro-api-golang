package routes

import (
	"common/common-consumer/apis"
	_ "common/common-consumer/docs"
	"core/middleware"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRoutes 初始化路由
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
	basePath := "/api/common"

	// 接口文档
	routes.GET(basePath+"/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL(basePath+"/v3/api-doc.json")),
	)

	routes.GET(basePath+"/v3/api-doc.json", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL(basePath+"/swagger/doc.json")),
	)

	// 权限管理
	auth := routes.Group(fmt.Sprintf("%s/auth", basePath))
	auth.GET("/sendMsg", middleware.Authorization("anonymous"), apis.Auth.SendMsg)                                            // 测试发送消息
	auth.GET("/createAnonymityToken/:appId", middleware.Authorization("anonymous"), apis.Auth.CreateAnonymityToken)           // 创建匿名token
	auth.POST("/login", middleware.Authorization("anonymous"), apis.Auth.Login)                                               // 网页后台登录
	auth.POST("/wechatLogin", middleware.Authorization("anonymous"), apis.Auth.WechatLogin)                                   // 小程序登录
	auth.POST("/wechatWebLogin", middleware.Authorization("anonymous"), apis.Auth.WechatWebLogin)                             // 公众号登录
	auth.GET("/logout", middleware.Authorization("anonymous"), apis.Auth.Logout)                                              // 网页后台登录
	auth.GET("/refreshPermissions", middleware.Authorization("system:auth:refreshPermissions"), apis.Auth.RefreshPermissions) // 网页后台登录
	auth.POST("/getWeChatJSSDK", middleware.Authorization("system:auth:getWeChatJSSDK"), apis.Auth.GetWeChatJSSDK)            // 获取微信jsSdk
	auth.GET("/getToken/:appId", middleware.Authorization("system:auth:getToken"), apis.Auth.GetToken)                        // 根据应用id获取授权token
	auth.POST("/createRoleMenu", middleware.Authorization("system:role:authMenu"), apis.Auth.CreateRoleMenu)                  // 创建角色关联菜单关系
	auth.POST("/createUserRole", middleware.Authorization("system:user:authUserRole"), apis.Auth.CreateUserRole)              // 创建用户关联角色关系
	auth.POST("/createRoleFunc", middleware.Authorization("system:user:authUserRole"), apis.Auth.CreateRoleFunc)              // 创建角色关联功能关系

	// 用户管理
	user := routes.Group(fmt.Sprintf("%s/user", basePath))
	user.GET("/logout", apis.User.Logout)                                                                              // 退出登陆
	user.POST("", middleware.Authorization("system:user:add"), apis.User.Create)                                       // 创建用户
	user.PUT("", middleware.Authorization("system:user:edit"), apis.User.Update)                                       // 更新用户
	user.DELETE("", middleware.Authorization("system:user:del"), apis.User.Delete)                                     // 删除用户
	user.GET("/findById/:id", middleware.Authorization("system:user:list"), apis.User.FindById)                        // 查询用户详情
	user.GET("/findWeChatWebUserInfo", middleware.Authorization("system:user:list"), apis.User.FindWeChatWebUserInfo)  // 查询用户的微信信息
	user.POST("/findPageList", middleware.Authorization("system:user:list"), apis.User.FindPageList)                   // 查询用户分页列表
	user.POST("/findUserIdByUserParam", middleware.Authorization("system:user:list"), apis.User.FindUserIdByUserParam) // 更具参数查询用户信息
	user.GET("/findDetail", middleware.Authorization("system:user:info"), apis.User.FindDetail)                        // 查询用户信息
	user.POST("/updateDetail", middleware.Authorization("system:user:edit"), apis.User.UpdateDetail)                   // 更新用户基本信息
	user.GET("/findBalance", middleware.Authorization("system:user:findBalance"), apis.User.FindBalance)               // 查询用户余额

	// 功能管理
	sysFunc := routes.Group(fmt.Sprintf("%s/func", basePath))
	sysFunc.POST("", middleware.Authorization("system:func:add"), apis.SysFunc.Create)
	sysFunc.PUT("", middleware.Authorization("system:func:edit"), apis.SysFunc.Update)
	sysFunc.DELETE("", middleware.Authorization("system:func:del"), apis.SysFunc.Delete)
	sysFunc.POST("/findFuncList", middleware.Authorization("system:func:list"), apis.SysFunc.FindList)
	sysFunc.GET("/findAreaById/:id", middleware.Authorization("system:func:list"), apis.SysFunc.FindById)
	sysFunc.GET("/findFuncInfoByRoleIds/:roleId", middleware.Authorization("system:func:list"), apis.SysFunc.FindInfoByRoleIds)

	// 收货地址
	userAddress := routes.Group(fmt.Sprintf("%s/userAddress", basePath))
	userAddress.POST("", middleware.Authorization("system:c:userAddress:manage"), apis.UserAddress.Create)
	userAddress.PUT("", middleware.Authorization("system:c:userAddress:manage"), apis.UserAddress.Update)
	userAddress.DELETE("", middleware.Authorization("system:c:userAddress:manage"), apis.UserAddress.Delete)
	userAddress.GET("/findUserAddressList", middleware.Authorization("system:c:userAddress:manage"), apis.UserAddress.FindList)
	userAddress.GET("/findById/:id", middleware.Authorization("system:c:userAddress:manage"), apis.UserAddress.FindById)

	// 地区管理
	area := routes.Group(fmt.Sprintf("%s/area", basePath))
	area.POST("", middleware.Authorization("system:area:add"), apis.Area.Create)
	area.PUT("", middleware.Authorization("system:area:edit"), apis.Area.Update)
	area.DELETE("", middleware.Authorization("system:area:del"), apis.Area.Delete)
	area.POST("/findAreaPageList", middleware.Authorization("system:area:list"), apis.Area.FindPageList)
	area.GET("/findAreaById/:id", middleware.Authorization("system:area:list"), apis.Area.FindById)
	area.GET("/findAreaByAreaCode", middleware.Authorization("system:area:list"), apis.Area.FindByAreaCode)
	area.GET("/findAreaListByParentCode", middleware.Authorization("system:area:list"), apis.Area.FindListByParentCode)
	area.GET("/findAreaTree", middleware.Authorization("system:area:list"), apis.Area.FindTree)

	// 菜单管理
	menu := routes.Group(fmt.Sprintf("%s/menu", basePath))
	menu.POST("", middleware.Authorization("system:menu:add"), apis.Menu.Create)
	menu.PUT("", middleware.Authorization("system:menu:edit"), apis.Menu.Update)
	menu.DELETE("", middleware.Authorization("system:menu:del"), apis.Menu.Delete)
	menu.GET("/findTreeList", middleware.Authorization("system:menu:list"), apis.Menu.FindTreeList)
	menu.GET("/findAll", middleware.Authorization("system:menu:list"), apis.Menu.FindAll)
	menu.GET("/findByRoleIds/:roleId", middleware.Authorization("system:menu:list"), apis.Menu.FindInfoByRoleIds)

	// 字典管理
	dict := routes.Group(fmt.Sprintf("%s/dict", basePath))
	dict.POST("", middleware.Authorization("system:dict:add"), apis.Dict.Create)
	dict.PUT("", middleware.Authorization("system:dict:edit"), apis.Dict.Update)
	dict.DELETE("", middleware.Authorization("system:dict:del"), apis.Dict.Delete)
	dict.GET("/findDictListByParentId/:parentId", middleware.Authorization("system:dict:list"), apis.Dict.FindListByParentId)
	dict.GET("/findById/:id", middleware.Authorization("system:dict:list"), apis.Dict.FindById)

	// 应用管理
	app := routes.Group(fmt.Sprintf("%s/app", basePath))
	app.POST("", middleware.Authorization("system:app:add"), apis.App.Create)
	app.PUT("", middleware.Authorization("system:app:edit"), apis.App.Update)
	app.DELETE("", middleware.Authorization("system:app:del"), apis.App.Delete)
	app.GET("/findAppList", middleware.Authorization("system:app:list"), apis.App.FindList)          // 查询当前账号所拥有的应用
	app.POST("/findAppPageList", middleware.Authorization("system:app:list"), apis.App.FindPageList) // 查询分页列表下的应用

	// 部门管理
	dept := routes.Group(fmt.Sprintf("%s/dept", basePath))
	dept.POST("", middleware.Authorization("system:dept:add"), apis.Department.Create)
	dept.PUT("", middleware.Authorization("system:dept:edit"), apis.Department.Update)
	dept.DELETE("", middleware.Authorization("system:dept:del"), apis.Department.Delete)
	// 查询当前账号所拥有的部门
	dept.GET("/findTree", middleware.Authorization("system:dept:manage"), apis.Department.FindTree)

	// 岗位管理
	post := routes.Group(fmt.Sprintf("%s/post", basePath))
	post.POST("", middleware.Authorization("system:post:add"), apis.Post.Create)
	post.PUT("", middleware.Authorization("system:post:edit"), apis.Post.Update)
	post.DELETE("", middleware.Authorization("system:post:del"), apis.Post.Delete)
	post.DELETE("/findPageList", middleware.Authorization("system:post:manage"), apis.Post.FindPageList)

	// 应用模块管理
	appModule := routes.Group(fmt.Sprintf("%s/appModule", basePath))
	appModule.POST("", middleware.Authorization("system:appModule:add"), apis.AppModule.Create)
	appModule.PUT("", middleware.Authorization("system:appModule:edit"), apis.AppModule.Update)
	appModule.DELETE("", middleware.Authorization("system:appModule:del"), apis.AppModule.Delete)
	appModule.GET("/findAreaById", middleware.Authorization("system:appModule:list"), apis.AppModule.FindById)
	appModule.GET("/findAppModuleListByAppId", middleware.Authorization("system:appModule:list"), apis.AppModule.FindList)

	// 微应用管理
	mircoApp := routes.Group(fmt.Sprintf("%s/mircoApp", basePath))
	mircoApp.PUT("", middleware.Authorization("system:mircoApp:edit"), apis.Dict.Update)

	// 配置管理
	config := routes.Group(fmt.Sprintf("%s/config", basePath))
	config.POST("", middleware.Authorization("system:config:add"), apis.Config.Create)
	config.PUT("", middleware.Authorization("system:config:edit"), apis.Config.Update)
	config.DELETE("", middleware.Authorization("system:config:del"), apis.Config.Delete)
	config.GET("/findById/:id", middleware.Authorization("system:config:list"), apis.Config.FindById)
	config.POST("/findByConfigExpandParam", middleware.Authorization("system:config:list"), apis.Config.FindByConfigExpandParam)

	// 权限管理
	role := routes.Group(fmt.Sprintf("%s/role", basePath))
	role.POST("", middleware.Authorization("system:role:add"), apis.Role.Create)
	role.PUT("", middleware.Authorization("system:role:add"), apis.Role.Update)
	role.DELETE("", middleware.Authorization("system:role:add"), apis.Role.Delete)
	role.POST("/findPageList", middleware.Authorization("system:role:list"), apis.Role.FindPageList)
	role.GET("/findAuthRoleListByUserId/:userId", middleware.Authorization("system:role:list"), apis.Role.FindAuthListByUserId)

	// 协议声明
	agreement := routes.Group(fmt.Sprintf("%s/agreement", basePath))
	agreement.POST("", middleware.Authorization("system:agreement:add"), apis.Agreement.Create)
	agreement.PUT("", middleware.Authorization("system:agreement:edit"), apis.Agreement.Update)
	agreement.DELETE("", middleware.Authorization("system:agreement:del"), apis.Agreement.Delete)
	agreement.POST("/findPageList", middleware.Authorization("system:agreement:list"), apis.Agreement.FindPageList)
	agreement.GET("/findById/:id", middleware.Authorization("anonymous"), apis.Agreement.FindById)

	// 组织声明
	org := routes.Group(fmt.Sprintf("%s/org", basePath))
	org.POST("", middleware.Authorization("system:org:add"), apis.Org.Create)
	org.PUT("", middleware.Authorization("system:org:edit"), apis.Org.Update)
	org.DELETE("", middleware.Authorization("system:org:delete"), apis.Org.Delete)
	org.POST("/findPageList", apis.Org.FindPageList)
	org.GET("/findById/:id", apis.Org.FindById)

	// 短信发送
	sms := routes.Group(fmt.Sprintf("%s/sms", basePath))
	sms.GET("/:phone/:type", apis.Sms.Send)

	return routes, basePath
}
