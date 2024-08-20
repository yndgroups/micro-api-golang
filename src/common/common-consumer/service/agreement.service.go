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

var Agreement = &agreement{}

type agreement struct{}

// Create 协议新增
func (a *agreement) Create(ctx *gin.Context) any {
	agreement := common.Agreement{}
	if err := ctx.ShouldBindJSON(&agreement); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&agreement); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	agreement.AgreeId = redis.GET.GetPrimaryKey("SYS_SHOPP_ADMIN_REQUEST_COUNT")
	agreement.CreateBy = ctx.Param("requestUserId")
	if result, getResultError := common.NewAgreementServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &agreement); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Insert(result.Count, result.Msg)
	}
}

// Update 协议修改
func (a *agreement) Update(ctx *gin.Context) any {
	agreement := common.Agreement{}
	if err := ctx.ShouldBindJSON(&agreement); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&agreement); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	agreement.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := common.NewAgreementServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &agreement); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}

// Delete 协议删除
func (a *agreement) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := common.NewAgreementServiceClient(app.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.AgreementIds{
		AgreeIds: ids,
		UserId:   ctx.Param("requestUserId"),
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

// FindById 查询协议
func (a *agreement) FindById(ctx *gin.Context) any {
	if result, getResultError := common.NewAgreementServiceClient(app.Nacos.CommonGrpcProvider()).FindById(context.Background(), &common.AgreementIds{
		AgreeIds: []string{ctx.Param("id")},
		UserId:   ctx.Param("requestUserId"),
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

// FindPageList 查询协议分页列表
func (a *agreement) FindPageList(ctx *gin.Context) any {
	request := common.AgreementPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	if result, getResultError := common.NewAgreementServiceClient(app.Nacos.CommonGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}
