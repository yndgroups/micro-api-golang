package service

import (
	"context"
	"core/app"
	"core/config"
	"core/logger"
	"core/resp"
	"protobuf/build/shopping_client"

	"github.com/gin-gonic/gin"
)

type ad struct{}

var Ad = &ad{}

// FindById 查询广告
func (a *ad) FindById(ctx *gin.Context) any {
	if result, getResultError := shopping_client.NewAdServiceClient(app.Nacos.ShoppingClientProvider()).FindById(context.Background(), &shopping_client.AdIds{
		AdIds:  []string{ctx.Param("id")},
		UserId: ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, result.Detail)
	}
}

// 根据广告类型和广告为位置查询广告列表
func (a *ad) FindList(ctx *gin.Context) any {
	if result, getResultError := shopping_client.NewAdServiceClient(app.Nacos.ShoppingClientProvider()).FindAdList(context.Background(), &shopping_client.AddListParma{
		Type:     ctx.Param("type"),
		Position: ctx.Param("position"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, result.List)
	}
}
