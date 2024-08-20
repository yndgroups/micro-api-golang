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
	"protobuf/build/global"

	"github.com/gin-gonic/gin"
)

var SysFunc = &sysFunc{}

type sysFunc struct{}

// Create 功能新增信息
func (sf *sysFunc) Create(ctx *gin.Context) any {
	SysFunc := common.SysFunc{}
	if err := ctx.ShouldBindJSON(&SysFunc); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&SysFunc); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	SysFunc.FuncId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")
	SysFunc.AppId = ctx.Param("appId")
	SysFunc.CreateBy = ctx.Param("requestUserId")

	if result, getResultError := common.NewSysFuncServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &SysFunc); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Insert(result.Count)
	}
}

// Update 更新功能信息
func (sf *sysFunc) Update(ctx *gin.Context) any {
	SysFunc := common.SysFunc{}
	if err := ctx.ShouldBindJSON(&SysFunc); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&SysFunc); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	SysFunc.UpdateBy = ctx.Param("requestUserId")
	SysFunc.AppId = ctx.Param("appId")
	if result, getResultError := common.NewSysFuncServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &SysFunc); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}

// Delete 删除功能信息
func (sf *sysFunc) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := common.NewSysFuncServiceClient(app.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.SysFuncIds{
		FuncIds: ids,
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

// FindById 查询功能详情
func (sf *sysFunc) FindById(ctx *gin.Context) any {
	if result, getResultError := common.NewSysFuncServiceClient(app.Nacos.CommonGrpcProvider()).FindById(context.Background(), &common.SysFuncIds{
		FuncIds: []string{ctx.Param("id")},
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, result.Detail)
	}
}

// FindList 查询功能列表
func (sf *sysFunc) FindList(ctx *gin.Context) any {
	request := common.SysFuncAuthQuery{}
	auth := global.Auth{AppId: ctx.Param("appId"), UserId: ctx.Param("requestUserId")}
	request.Auth = &auth
	if result, getResultError := common.NewSysFuncServiceClient(app.Nacos.CommonGrpcProvider()).FindList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		if result.Count > 0 {
			if len(result.List) > 0 {
				// 权限值提取
				powerSign := make([]string, 0)
				for _, v := range result.List {
					if v.PowerSign != "" {
						powerSign = append(powerSign, v.PowerSign)
					}
				}
				// 遍历树状菜单
				funcs := make(map[string]any)
				funcs["funcsList"] = sf.treeRecursion(result.List, "1")
				funcs["powerSign"] = powerSign
				// 删除权限变更信息
				// redis.HDel("isResetAuth", ctx.Param("appId")+"_"+ctx.Param("requestUserId"))
				return resp.Success.WithData(funcs)
			}
		}
		return resp.NotData
	}
}

// FindInfoByRoleIds 根据角色id查询功能相关信息，主要用于给角色授功能
func (sf *sysFunc) FindInfoByRoleIds(ctx *gin.Context) any {
	if result, getResultError := common.NewSysFuncServiceClient(app.Nacos.CommonGrpcProvider()).FindFuncInfoByRoleIds(context.Background(), &common.RoleFuncParam{
		RoleId: ctx.Param("roleId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		if result.Count > 0 {
			if len(result.List) > 0 {
				// 遍历树状菜单
				menus := make(map[string]any)
				menus["funcsList"] = sf.treeRecursion(result.List, "1")
				menus["roleFuncIds"] = result.RoleFuncIds
				return resp.Success.WithData(menus)
			}
		}
		return resp.NotData
	}
}

// treeRecursion 递归处理功能数据
func (sf *sysFunc) treeRecursion(list []*common.SysFuncListVo, parentId string) []*common.SysFuncListVo {
	// 递归出传入的父级编码数据及非父级编码数据
	parentList := make([]*common.SysFuncListVo, 0)
	childList := make([]*common.SysFuncListVo, 0)
	for _, v := range list {
		if v.ParentId == parentId {
			parentList = append(parentList, v)
		} else {
			childList = append(childList, v)
		}
	}
	// 将非父级别数据进行分组
	childMapList := make(map[string][]*common.SysFuncListVo)
	for _, v := range childList {
		t := childMapList[v.ParentId]
		t = append(t, v)
		childMapList[v.ParentId] = t
	}
	// 最总结果
	for i, v := range parentList {
		child := childMapList[v.FuncId]
		if len(child) > 0 {
			t := sf.treeRecursion(childList, v.FuncId)
			parentList[i].Children = t
		} else {
			parentList[i].Children = child
		}
	}
	return parentList
}
