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

var FreightTemplatesFree = &freightTemplatesFree{}

type freightTemplatesFree struct{}

func (ftf *freightTemplatesFree) Create(ctx *gin.Context) any {
	freightTemplatesFree := order.FreightTemplatesFree{}
	if err := ctx.ShouldBindJSON(&freightTemplatesFree); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&freightTemplatesFree); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	freightTemplatesFree.FreeId = redis.GET.GetPrimaryKey("SYS_ORDER_REQUEST_COUNT")
	freightTemplatesFree.CreateBy = ctx.Param("requestUserId")
	if result, getResultError := order.NewFreightTemplatesFreeServiceClient(app.Nacos.OrderGrpcProvider()).Create(context.Background(), &freightTemplatesFree); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Insert(result.Count)
	}
}

func (ftf *freightTemplatesFree) Update(ctx *gin.Context) any {
	freightTemplatesFree := order.FreightTemplatesFree{}
	if err := ctx.ShouldBindJSON(&freightTemplatesFree); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&freightTemplatesFree); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	freightTemplatesFree.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := order.NewFreightTemplatesFreeServiceClient(app.Nacos.OrderGrpcProvider()).Update(context.Background(), &freightTemplatesFree); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Update(result.Count)
	}
}

func (ftf *freightTemplatesFree) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := order.NewFreightTemplatesFreeServiceClient(app.Nacos.OrderGrpcProvider()).Delete(context.Background(), &order.FreightTemplatesFreeIds{
		FreeId: ids,
		UserId: ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Delete(result.Count)
	}
}

func (ftf *freightTemplatesFree) FindPageList(ctx *gin.Context) any {
	request := order.FreightTemplatesFreePageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	if result, getResultError := order.NewFreightTemplatesFreeServiceClient(app.Nacos.OrderGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}

func (ftf *freightTemplatesFree) FindById(ctx *gin.Context) any {
	if result, getResultError := order.NewFreightTemplatesFreeServiceClient(app.Nacos.OrderGrpcProvider()).FindById(context.Background(), &order.FreightTemplatesFreeIds{
		FreeId: []string{ctx.Param("id")},
		UserId: ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, result.Detail)
	}
}
