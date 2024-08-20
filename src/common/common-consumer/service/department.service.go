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

var Department = &department{}

type department struct{}

// Create 新增应用
func (dp *department) Create(ctx *gin.Context) any {
	dept := common.Department{}
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&dept); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	dept.DeptId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")
	dept.CreateBy = ctx.Param("requestUserId")

	if result, getResultError := common.NewDepartmentServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &dept); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Insert(result.Count)
	}
}

// Update 修改应用
func (dp *department) Update(ctx *gin.Context) any {
	dept := common.Department{}
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&dept); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	if dept.DeptId == "" {
		return resp.Fail.WithMsg("部门id不能空")
	}
	dept.UpdateBy = ctx.Param("requestUserId")

	if result, getResultError := common.NewDepartmentServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &dept); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}

// Delete 删除应用
func (dp *department) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := common.NewDepartmentServiceClient(app.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.DepartmentIds{
		DeptIds: ids,
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Delete(result.Count)
	}
}

// FindTree 查找应用列表
func (dp *department) FindTree(ctx *gin.Context) any {
	auth := global.Auth{
		UserId: ctx.Param("requestUserId"),
	}
	if result, getResultError := common.NewDepartmentServiceClient(app.Nacos.CommonGrpcProvider()).FindTree(context.Background(), &auth); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, dp.treeRecursion(result.List, "1"))
	}
}

// 递归处理菜单数据
func (dp *department) treeRecursion(list []*common.Department, parent string) []*common.Department {
	// 递归出传入的父级编码数据及非父级编码数据
	parentList := make([]*common.Department, 0)
	childList := make([]*common.Department, 0)
	for _, v := range list {
		if v.ParentId == parent {
			parentList = append(parentList, v)
		} else {
			childList = append(childList, v)
		}
	}
	// 将非父级别数据进行分组
	childMapList := make(map[string][]*common.Department)
	for _, v := range childList {
		t := childMapList[v.ParentId]
		t = append(t, v)
		childMapList[v.ParentId] = t
	}
	// 最总结果
	for i, v := range parentList {
		child := childMapList[v.DeptId]
		if len(child) > 0 {
			t := dp.treeRecursion(childList, v.DeptId)
			parentList[i].Children = t
		} else {
			parentList[i].Children = child
		}
	}
	return parentList
}
