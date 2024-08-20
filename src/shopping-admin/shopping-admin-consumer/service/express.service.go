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

var Express = &express{}

type express struct{}

// Create 快递新增
func (e *express) Create(ctx *gin.Context) any {
	express := shopping_admin.Express{}
	if err := ctx.ShouldBindJSON(&express); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&express); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	express.CreateBy = ctx.Param("requestUserId")
	express.ExpressId = redis.GET.GetPrimaryKey("SYS_SHOPP_ADMIN_REQUEST_COUNT")
	if result, getResultError := shopping_admin.NewExpressServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Create(context.Background(), &express); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.Insert(result.Count, result.Msg)
	}
}

// UpdateExpress 快递修改
func (e *express) Update(ctx *gin.Context) any {
	express := shopping_admin.Express{}
	if err := ctx.ShouldBindJSON(&express); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&express); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	express.CreateBy = ctx.Param("requestUserId")
	express.ExpressId = redis.GET.GetPrimaryKey("SYS_SHOPP_ADMIN_REQUEST_COUNT")
	if result, getResultError := shopping_admin.NewExpressServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Update(context.Background(), &express); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.Insert(result.Count, result.Msg)
	}
}

// Delete 快递删除
func (e *express) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := shopping_admin.NewExpressServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Delete(context.Background(), &shopping_admin.ExpressIds{
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

// FindById 查询快递
func (e *express) FindById(ctx *gin.Context) any {
	if result, getResultError := shopping_admin.NewExpressServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindById(context.Background(), &shopping_admin.ExpressIds{
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

// FindPageList 查询快递分页列表
func (e *express) FindPageList(ctx *gin.Context) any {
	request := shopping_admin.ExpressPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	if result, getResultError := shopping_admin.NewExpressServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}
