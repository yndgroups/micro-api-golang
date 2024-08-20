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
	"protobuf/build/shopping_admin"

	"github.com/gin-gonic/gin"
)

var Ad = &ad{}

type ad struct{}

// Create 广告新增
func (a *ad) Create(ctx *gin.Context) any {
	ad := shopping_admin.Ad{}
	if err := ctx.ShouldBindJSON(&ad); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&ad); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	ad.CreateBy = ctx.Param("requestUserId")
	ad.AdId = redis.GET.GetPrimaryKey("SYS_SHOPP_ADMIN_REQUEST_COUNT")
	if result, getResultError := shopping_admin.NewAdServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Create(context.Background(), &ad); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.Insert(result.Count, result.Msg)
	}
}

// Update 广告修改
func (a *ad) Update(ctx *gin.Context) any {
	ad := shopping_admin.Ad{}
	if err := ctx.ShouldBindJSON(&ad); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&ad); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	ad.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := shopping_admin.NewAdServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Update(context.Background(), &ad); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.Insert(result.Count, result.Msg)
	}
}

// Delete 广告删除
func (a *ad) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := shopping_admin.NewAdServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Delete(context.Background(), &shopping_admin.AdIds{
		Id:     ids,
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

// FindById 查询广告
func (a *ad) FindById(ctx *gin.Context) any {
	if result, getResultError := shopping_admin.NewAdServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindById(context.Background(), &shopping_admin.AdIds{
		Id:     []string{ctx.Param("id")},
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

// FindPageList 查询广告分页列表
func (a *ad) FindPageList(ctx *gin.Context) any {
	request := shopping_admin.AdPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	if result, getResultError := shopping_admin.NewAdServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}
