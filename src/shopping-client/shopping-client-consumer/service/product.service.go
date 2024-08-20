package service

import (
	"context"
	"core/app"
	"core/config"
	"core/logger"
	"core/pager"
	"core/resp"
	"core/validate"
	"protobuf/build/shopping_client"

	"github.com/gin-gonic/gin"
)

var Product = &product{}

type product struct{}

// FindPageList 查询商品列表
func (p *product) FindPageList(ctx *gin.Context) any {
	request := shopping_client.ProductPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	if result, getResultError := shopping_client.NewProductServiceClient(app.Nacos.ShoppingClientProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingClientProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}

// FindById 查询商品详情
func (p *product) FindById(ctx *gin.Context) any {
	// 校验请求参数
	if result, getResultError := shopping_client.NewProductServiceClient(app.Nacos.ShoppingClientProvider()).FindById(context.Background(), &shopping_client.ProductIds{
		ProductIds: []string{ctx.Param("productId")},
		UserId:     ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, result.Detail)
	}
}

func (p *product) FindListByProductIds(ctx *gin.Context) any {
	return nil
}
