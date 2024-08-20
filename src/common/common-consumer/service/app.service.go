package service

import (
	"context"
	appClient "core/app"
	"core/config"
	"core/logger"
	"core/model"
	"core/pager"
	"core/redis"
	"core/resp"
	"core/validate"
	"protobuf/build/common"
	"protobuf/build/global"

	"github.com/gin-gonic/gin"
)

var App = &application{}

type application struct{}

// Create 新增应用
func (a *application) Create(ctx *gin.Context) any {
	app := common.App{}
	if err := ctx.ShouldBindJSON(&app); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&app); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	app.AppId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")
	app.CreateBy = ctx.Param("requestUserId")

	if result, getResultError := common.NewAppServiceClient(appClient.Nacos.CommonGrpcProvider()).Create(context.Background(), &app); getResultError != nil {
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
func (a *application) Update(ctx *gin.Context) any {
	app := common.App{}
	if err := ctx.ShouldBindJSON(&app); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&app); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	app.UpdateBy = ctx.Param("requestUserId")

	if result, getResultError := common.NewAppServiceClient(appClient.Nacos.CommonGrpcProvider()).Update(context.Background(), &app); getResultError != nil {
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
func (a *application) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}

	if result, getResultError := common.NewAppServiceClient(appClient.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.AppIds{
		AppIds: ids,
		UserId: ctx.Param("requestUserId"),
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

// FindList 查找应用列表
func (a *application) FindList(ctx *gin.Context) any {
	auth := global.Auth{
		UserId: ctx.Param("requestUserId"),
	}
	if result, getResultError := common.NewAppServiceClient(appClient.Nacos.CommonGrpcProvider()).FindList(context.Background(), &auth); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, result.List)
	}
}

func (a *application) FindPageList(ctx *gin.Context) any {
	request := common.AppPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}

	if result, getResultError := common.NewAppServiceClient(appClient.Nacos.CommonGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}
