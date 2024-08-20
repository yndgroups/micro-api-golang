package service

import (
	"context"
	"core/app"
	"core/config"
	"core/logger"
	"core/model"
	"core/pager"
	"core/resp"
	"core/validate"
	"protobuf/build/shopping_admin"

	"github.com/gin-gonic/gin"
)

var BusinessStore = &businessStore{}

type businessStore struct{}

// Create 商家加盟新增
func (bf *businessStore) Create(ctx *gin.Context) any {
	businessStore := shopping_admin.BusinessStore{}
	if err := ctx.ShouldBindJSON(&businessStore); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&businessStore); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	businessStore.CreateBy = ctx.Param("requestUserId")
	//businessFranchisee.BrandId = redis.GET.GetPrimaryKey("SYS_SHOPP_ADMIN_REQUEST_COUNT")
	if result, getResultError := shopping_admin.NewBusinessStoreServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Create(context.Background(), &businessStore); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.Insert(result.Count, result.Msg)
	}
}

// Update 商家加盟修改
func (bf *businessStore) Update(ctx *gin.Context) any {
	businessStore := shopping_admin.BusinessStore{}
	if err := ctx.ShouldBindJSON(&businessStore); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&businessStore); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	businessStore.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := shopping_admin.NewBusinessStoreServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Create(context.Background(), &businessStore); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.Insert(result.Count, result.Msg)
	}
}

// Delete 商家门店删除
func (bs *businessStore) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := shopping_admin.NewBusinessStoreServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Delete(context.Background(), &shopping_admin.BusinessStoreIds{
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

// FindPageList 查询商家门店分页列表
func (bs *businessStore) FindPageList(ctx *gin.Context) any {
	request := shopping_admin.BusinessStorePageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	if result, getResultError := shopping_admin.NewBusinessStoreServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}

// FindById 查询商家门店
func (bs *businessStore) FindById(ctx *gin.Context) any {
	if result, getResultError := shopping_admin.NewBusinessStoreServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindById(context.Background(), &shopping_admin.BusinessStoreIds{
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
