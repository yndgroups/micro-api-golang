package service

import (
	"context"
	"core/app"
	"core/config"
	"core/logger"
	"core/model"
	"core/pager"
	"core/redis"
	"core/resp"
	"core/validate"
	"protobuf/build/common"

	"github.com/gin-gonic/gin"
)

var Role = &role{}

type role struct{}

// Create 创建角色
func (r *role) Create(ctx *gin.Context) any {
	role := common.Role{}
	if err := ctx.ShouldBindJSON(&role); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&role); err != nil {
		return resp.ErrValidate.WithData(validateMsg)
	}
	role.RoleId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")
	role.CreateBy = ctx.Param("requestUserId")
	if result, getResultError := common.NewRoleServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &role); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Insert(result.Count)
	}
}

// Update 修改角色
func (r *role) Update(ctx *gin.Context) any {
	role := common.Role{}
	if err := ctx.ShouldBindJSON(&role); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&role); err != nil {
		return resp.ErrValidate.WithData(validateMsg)
	}
	role.RoleId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")
	role.CreateBy = ctx.Param("requestUserId")
	if result, getResultError := common.NewRoleServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &role); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}

// Delete 删除角色
func (r *role) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := common.NewRoleServiceClient(app.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.RoleIds{
		RoleId: ids,
		UserId: ctx.Param("requestUserId"),
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

// FindPageList 查询服务列表
func (r *role) FindPageList(ctx *gin.Context) any {
	request := common.RolePageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	// 只能使用当前系统颁发token的appId,就算前传入该参数也给覆盖
	request.Params.AppId = ctx.Param("appId")
	if result, getResultError := common.NewRoleServiceClient(app.Nacos.CommonGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}

// FindAuthListByUserId 根据用户id查询其在应用下的角色及授权信息
func (r *role) FindAuthListByUserId(ctx *gin.Context) any {
	if result, getResultError := common.NewRoleServiceClient(app.Nacos.CommonGrpcProvider()).FindAuthRoleListByUserId(context.Background(), &common.RoleAuthParam{
		UserId: ctx.Param("userId"),
		AppId:  ctx.Param("appId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, result)
	}
}
