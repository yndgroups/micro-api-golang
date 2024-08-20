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

var FreightTemplatesRegion = &freightTemplatesRegion{}

type freightTemplatesRegion struct{}

func (ftr *freightTemplatesRegion) Create(ctx *gin.Context) any {
	freightTemplatesRegion := order.FreightTemplatesRegion{}
	if err := ctx.ShouldBindJSON(&freightTemplatesRegion); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&freightTemplatesRegion); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	freightTemplatesRegion.RegionId = redis.GET.GetPrimaryKey("SYS_ORDER_REQUEST_COUNT")
	freightTemplatesRegion.CreateBy = ctx.Param("requestUserId")
	if result, getResultError := order.NewFreightTemplatesRegionServiceClient(app.Nacos.OrderGrpcProvider()).Create(context.Background(), &freightTemplatesRegion); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Insert(result.Count)
	}
}

func (ftr *freightTemplatesRegion) Update(ctx *gin.Context) any {
	freightTemplatesRegion := order.FreightTemplatesRegion{}
	if err := ctx.ShouldBindJSON(&freightTemplatesRegion); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&freightTemplatesRegion); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	freightTemplatesRegion.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := order.NewFreightTemplatesRegionServiceClient(app.Nacos.OrderGrpcProvider()).Update(context.Background(), &freightTemplatesRegion); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Update(result.Count)
	}
}

func (ftr *freightTemplatesRegion) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := order.NewFreightTemplatesRegionServiceClient(app.Nacos.OrderGrpcProvider()).Delete(context.Background(), &order.FreightTemplatesRegionIds{
		RegionId: ids,
		UserId:   ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Delete(result.Count)
	}
}

func (ftr *freightTemplatesRegion) FindPageList(ctx *gin.Context) any {
	request := order.FreightTemplatesRegionPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	if result, getResultError := order.NewFreightTemplatesRegionServiceClient(app.Nacos.OrderGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}

func (ftr *freightTemplatesRegion) FindById(ctx *gin.Context) any {
	if result, getResultError := order.NewFreightTemplatesRegionServiceClient(app.Nacos.OrderGrpcProvider()).FindById(context.Background(), &order.FreightTemplatesRegionIds{
		RegionId: []string{ctx.Param("id")},
		UserId:   ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, result.Detail)
	}
}
