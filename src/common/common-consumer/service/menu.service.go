package service

import (
	"context"
	"core/app"
	"core/config"
	"core/logger"
	"core/model"
	"core/redis"
	"core/resp"
	"core/validate"
	"protobuf/build/common"

	"github.com/gin-gonic/gin"
)

var Menu = &menu{}

type menu struct{}

// Create 创建菜单
func (m *menu) Create(ctx *gin.Context) any {
	menu := common.Menu{}
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&menu); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	menu.MenuId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")
	menu.CreateBy = ctx.Param("requestUserId") // 取登录的Redis用户的userId
	menu.AppId = ctx.Param("appId")            // 不能让用户传避免bug

	if result, getResultError := common.NewMenuServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &menu); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Insert(result.Count)
	}
}

// Update 更新菜单
func (m *menu) Update(ctx *gin.Context) any {
	menu := common.Menu{}
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if menu.MenuId == "" {
		return resp.Fail.WithMsg("menuId不能为空")
	}
	if validateMsg, err := validate.Validated(&menu); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	menu.UpdateBy = ctx.Param("requestUserId") // 取登录的Redis用户的userId
	menu.AppId = ctx.Param("appId")            // 不能让用户传避免bug
	if result, getResultError := common.NewMenuServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &menu); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}

// Delete 删除菜单
func (m *menu) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if len(ids) == 0 {
		return resp.Fail.WithMsg("menuId参数不能为空")
	}

	if result, getResultError := common.NewMenuServiceClient(app.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.MenuIds{
		MenuIds: ids,
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Delete(result.Count)
	}
}

// FindTreeList 根据用户权限进行菜单查询
func (m *menu) FindTreeList(ctx *gin.Context) any {
	// 查询权限数据
	params := common.RoleMenuParam{
		AppId:  ctx.Param("appId"),
		UserId: ctx.Param("requestUserId"),
	}
	if result, getResultError := common.NewMenuServiceClient(app.Nacos.CommonGrpcProvider()).FindTreeList(context.Background(), &params); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		if result.Count > 0 {
			if len(result.List) > 0 {
				menus := make(map[string]any)
				// 创建树状结构菜单
				menus["menuList"] = m.treeRecursion(result.List, "1")
				menus["powerSign"] = result.PowerSign
				redis.SET.Set("power_"+ctx.Param("appId")+"_"+ctx.Param("requestUserId"), result.PowerSign, config.Global.RedisConfig().Expire)
				redis.DELETE.HashDel("power_change", ctx.Param("appId")+"_"+ctx.Param("requestUserId"))
				return resp.Success.WithData(menus)
			}
		}
		return resp.NotData
	}
}

// FindAll 查询应用下的所有菜单
func (m *menu) FindAll(ctx *gin.Context) any {
	// 查询权限数据
	params := common.RoleMenuParam{
		AppId:  ctx.Param("appId"),
		UserId: ctx.Param("requestUserId"),
	}
	if result, getResultError := common.NewMenuServiceClient(app.Nacos.CommonGrpcProvider()).FindAll(context.Background(), &params); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		if result.Count > 0 {
			if len(result.List) > 0 {
				return resp.Success.WithData(m.treeRecursion(result.List, "1"))
			}
		}
		return resp.NotData
	}
}

// FindInfoByRoleIds 根据角色id查询菜单相关信息，主要用于给角色授权菜单
func (m *menu) FindInfoByRoleIds(ctx *gin.Context) any {
	if result, getResultError := common.NewMenuServiceClient(app.Nacos.CommonGrpcProvider()).FindByRoleIds(context.Background(), &common.RoleMenuParam{
		RoleIds: []string{ctx.Param("roleId")},
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		if len(result.List) > 0 {
			// 遍历树状菜单
			menus := make(map[string]any)
			menus["menuList"] = m.treeRecursion(result.List, "1")
			menus["roleMenuIds"] = result.RoleMenuIds
			return resp.Success.WithData(menus)
		}
		return resp.NotData
	}
}

// treeRecursion 递归处理菜单数据
func (m *menu) treeRecursion(list []*common.Menu, parent string) []*common.Menu {
	// 递归出传入的父级编码数据及非父级编码数据
	parentList := make([]*common.Menu, 0)
	childList := make([]*common.Menu, 0)
	for _, v := range list {
		if v.ParentId == parent {
			parentList = append(parentList, v)
		} else {
			childList = append(childList, v)
		}
	}
	// 将非父级别数据进行分组
	childMapList := make(map[string][]*common.Menu)
	for _, v := range childList {
		t := childMapList[v.ParentId]
		t = append(t, v)
		childMapList[v.ParentId] = t
	}
	// 最总结果
	for i, v := range parentList {
		child := childMapList[v.MenuId]
		if len(child) > 0 {
			t := m.treeRecursion(childList, v.MenuId)
			parentList[i].Children = t
		} else {
			parentList[i].Children = child
		}
	}
	return parentList
}
