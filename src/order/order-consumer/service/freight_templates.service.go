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
	"protobuf/build/order"

	"github.com/gin-gonic/gin"
)

var FreightTemplates = &freightTemplates{}

type freightTemplates struct{}

func (ft *freightTemplates) Create(ctx *gin.Context) any {
	freightTemplates := order.FreightTemplates{}
	if err := ctx.ShouldBindJSON(&freightTemplates); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&freightTemplates); err != nil {
		return resp.ErrValidate.WithData(validateMsg)
	}
	freightTemplates.TempId = redis.GET.GetPrimaryKey("SYS_ORDER_REQUEST_COUNT")
	freightTemplates.CreateBy = ctx.Param("requestUserId")
	if result, getResultError := order.NewFreightTemplatesServiceClient(app.Nacos.OrderGrpcProvider()).Create(context.Background(), &freightTemplates); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Insert(result.Count)
	}
}

func (ft *freightTemplates) Update(ctx *gin.Context) any {
	freightTemplates := order.FreightTemplates{}
	if err := ctx.ShouldBindJSON(&freightTemplates); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&freightTemplates); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	freightTemplates.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := order.NewFreightTemplatesServiceClient(app.Nacos.OrderGrpcProvider()).Update(context.Background(), &freightTemplates); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Update(result.Count)
	}
}

func (ft *freightTemplates) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := order.NewFreightTemplatesServiceClient(app.Nacos.OrderGrpcProvider()).Delete(context.Background(), &order.FreightTemplateIds{
		TempId: ids,
		UserId: ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Delete(result.Count)
	}
}

func (ft *freightTemplates) FindPageList(ctx *gin.Context) any {
	request := order.FreightTemplatesPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	if result, getResultError := order.NewFreightTemplatesServiceClient(app.Nacos.OrderGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}

func (ft *freightTemplates) FindById(ctx *gin.Context) any {
	if result, getResultError := order.NewFreightTemplatesServiceClient(app.Nacos.OrderGrpcProvider()).FindById(context.Background(), &order.FreightTemplateIds{
		TempId: []string{ctx.Param("id")},
		UserId: ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, result.Detail)
	}
}
