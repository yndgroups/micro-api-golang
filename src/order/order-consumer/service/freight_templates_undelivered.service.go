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

var FreightTemplatesUndelivered = &freightTemplatesUndelivered{}

type freightTemplatesUndelivered struct{}

func (ftu *freightTemplatesUndelivered) Create(ctx *gin.Context) any {
	freightTemplatesUndelivered := order.FreightTemplatesUndelivered{}
	if err := ctx.ShouldBindJSON(&freightTemplatesUndelivered); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&freightTemplatesUndelivered); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	freightTemplatesUndelivered.UndeliveredId = redis.GET.GetPrimaryKey("SYS_ORDER_REQUEST_COUNT")
	freightTemplatesUndelivered.CreateBy = ctx.Param("requestUserId")
	if result, getResultError := order.NewFreightTemplatesUndeliveredServiceClient(app.Nacos.OrderGrpcProvider()).Create(context.Background(), &freightTemplatesUndelivered); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Insert(result.Count)
	}
}

func (ftu *freightTemplatesUndelivered) Update(ctx *gin.Context) any {
	freightTemplatesUndelivered := order.FreightTemplatesUndelivered{}
	if err := ctx.ShouldBindJSON(&freightTemplatesUndelivered); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&freightTemplatesUndelivered); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	freightTemplatesUndelivered.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := order.NewFreightTemplatesUndeliveredServiceClient(app.Nacos.OrderGrpcProvider()).Update(context.Background(), &freightTemplatesUndelivered); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Update(result.Count)
	}
}

// DeleteFreightTemplatesUndelivered 删除不送达模板
func (ftu *freightTemplatesUndelivered) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := order.NewFreightTemplatesUndeliveredServiceClient(app.Nacos.OrderGrpcProvider()).Delete(context.Background(), &order.FreightTemplatesUndeliveredIds{
		UndeliveredId: ids,
		UserId:        ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Delete(result.Count)
	}
}

// FindPageList 不送达详情分页
func (ftu *freightTemplatesUndelivered) FindPageList(ctx *gin.Context) any {
	request := order.FreightTemplatesUndeliveredPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	if result, getResultError := order.NewFreightTemplatesUndeliveredServiceClient(app.Nacos.OrderGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}

// FindById 查询不送达模板详情
func (ftu *freightTemplatesUndelivered) FindById(ctx *gin.Context) any {
	if result, getResultError := order.NewFreightTemplatesUndeliveredServiceClient(app.Nacos.OrderGrpcProvider()).FindById(context.Background(), &order.FreightTemplatesUndeliveredIds{
		UndeliveredId: []string{ctx.Param("id")},
		UserId:        ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, result.Detail)
	}
}
