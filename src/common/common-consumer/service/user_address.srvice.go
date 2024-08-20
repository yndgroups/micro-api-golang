package service

import (
	"context"
	"core/app"
	"core/config"
	"core/logger"
	"core/model"
	"core/redis"
	"core/resp"
	"core/validate"
	"protobuf/build/common"

	"github.com/gin-gonic/gin"
)

var UserAddress = &userAddress{}

type userAddress struct{}

// CreateUser 收货地址新增信息
func (ua *userAddress) Create(ctx *gin.Context) any {
	userAddress := common.UserAddress{}
	if err := ctx.ShouldBindJSON(&userAddress); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&userAddress); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	userAddress.CreateBy = ctx.Param("requestUserId")
	userAddress.AddressId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")
	if result, getResultError := common.NewUserAddressServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &userAddress); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Insert(result.Count)
	}
}

// Update 更新收货地址信息
func (ua *userAddress) Update(ctx *gin.Context) any {
	userAddress := common.UserAddress{}
	if err := ctx.ShouldBindJSON(&userAddress); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&userAddress); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	userAddress.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := common.NewUserAddressServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &userAddress); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}

// Delete 删除收货地址信息
func (ua *userAddress) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := common.NewUserAddressServiceClient(app.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.UserAddressIds{
		AddressIds: ids,
		UserId:     ctx.Param("requestUserId"),
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

// FindList 查询收货地址列表
func (ua *userAddress) FindList(ctx *gin.Context) any {
	if result, getResultError := common.NewUserAddressServiceClient(app.Nacos.CommonGrpcProvider()).FindList(context.Background(), &common.UserAddressIds{
		AddressIds: nil,
		UserId:     ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, result.List)
	}
}

// FindById 查询收货地址详情
func (ua *userAddress) FindById(ctx *gin.Context) any {
	if result, getResultError := common.NewUserAddressServiceClient(app.Nacos.CommonGrpcProvider()).FindById(context.Background(), &common.UserAddressIds{
		AddressIds: []string{ctx.Param("id")},
		UserId:     ctx.Param("requestUserId"),
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
