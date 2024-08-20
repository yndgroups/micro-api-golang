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

var AppModule = &appModule{}

type appModule struct{}

func (am *appModule) Create(ctx *gin.Context) any {
	appModule := common.AppModule{}
	if err := ctx.ShouldBindJSON(&appModule); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&appModule); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	appModule.CreateBy = ctx.Param("requestUserId")
	appModule.MdId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")

	if result, getResultError := common.NewAppModuleServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &appModule); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Insert(result.Count)
	}
}

func (am *appModule) Update(ctx *gin.Context) any {
	appModule := common.AppModule{}
	if err := ctx.ShouldBindJSON(&appModule); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&appModule); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	appModule.UpdateBy = ctx.Param("requestUserId")

	if result, getResultError := common.NewAppModuleServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &appModule); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}

func (am *appModule) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}

	if result, getResultError := common.NewAppModuleServiceClient(app.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.AppModuleIds{
		MdIds:  ids,
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

func (am *appModule) FindById(ctx *gin.Context) any {

	if result, getResultError := common.NewAppModuleServiceClient(app.Nacos.CommonGrpcProvider()).FindById(context.Background(), &common.AppModuleIds{
		MdIds:  []string{ctx.Param("id")},
		UserId: ctx.Param("requestUserId"),
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

func (am *appModule) FindList(ctx *gin.Context) any {
	request := common.AppModulePageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}

	if result, getResultError := common.NewAppModuleServiceClient(app.Nacos.CommonGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}
