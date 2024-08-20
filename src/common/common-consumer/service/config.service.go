package service

import (
	"context"
	"core/app"
	"core/logger"
	"core/resp"
	"core/validate"
	"fmt"
	"protobuf/build/common"

	coreConfig "core/config"

	"github.com/gin-gonic/gin"
)

var Config = &configService{}

type configService struct{}

// FindByConfigExpandParam 根据参数查询配置文件
func (cfg *configService) FindByConfigExpandParam(ctx *gin.Context) any {
	configIds := common.ConfigIds{}
	if err := ctx.ShouldBindJSON(&configIds); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&configIds); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	configIds.AppId = ctx.Param("appId")
	// 根据openId查找是否已经注册
	result, getResultError := common.NewConfigServiceClient(app.Nacos.CommonGrpcProvider()).FindByConfigExpandParam(context.Background(), &configIds)
	if getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", coreConfig.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg(fmt.Sprintf("%v服务生产者调用异常", coreConfig.Global.ServerConfig().CommonProvider))
	}
	if result.Msg != "" {
		return resp.Fail.WithData(result.Msg)
	}
	if len(result.List) > 1 {
		return resp.Success.WithData(result.List)
	}
	return resp.Success.WithData(result.Detail)
}
