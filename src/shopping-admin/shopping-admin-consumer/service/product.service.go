package service

import (
	"context"
	"core/app"
	"core/config"
	"core/db"
	"core/logger"
	"core/model"
	"core/redis"
	"core/req"
	"core/resp"
	"core/validate"
	"protobuf/build/shopping_admin"
	"shopping-admin/shopping-admin-consumer/enum"

	"github.com/gin-gonic/gin"
)

var Product = product{}

type product struct{}

// Create 商品新增
func (p *product) Create(ctx *gin.Context) any {
	product := shopping_admin.Product{}
	if err := ctx.ShouldBindJSON(&product); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&product); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	product.CreateBy = ctx.Param("requestUserId")
	product.ProductId = redis.GET.GetPrimaryKey("SYS_SHOPP_ADMIN_REQUEST_COUNT")
	// 组装sku数据
	if result, getResultError := shopping_admin.NewProductServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Create(context.Background(), &product); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.Insert(result.Count, result.Msg)
	}
}

// Update 商品修改
func (p *product) Update(ctx *gin.Context) any {
	product := shopping_admin.Product{}
	if err := ctx.ShouldBindJSON(&product); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&product); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	product.UpdateBy = ctx.Param("requestUserId")
	// 组装sku数据
	if result, getResultError := shopping_admin.NewProductServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Create(context.Background(), &product); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.Update(result.Count, result.Msg)
	}
}

// Delete 商品删除
func (p *product) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	return db.BatchDeleteByIds(ids, enum.Product.TableName, enum.Product.PrimaryKey, ctx.Param("requestUserId"))
}

// FindById 查询商品详情
func (p *product) FindById(ctx *gin.Context) any {
	if result, getResultError := shopping_admin.NewProductServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindById(context.Background(), &shopping_admin.ProductIds{
		ProductId: []string{ctx.Param("productId")},
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.WithData(result.Detail)
	}
}

// FindPageList 查询商品列表
func (p *product) FindPageList(ctx *gin.Context) any {
	request := req.Request[shopping_admin.ProductRequestParams]{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validate, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validate)
	}
	auth := shopping_admin.ProductPageAuthQuery{}
	authParams := req.CreateAuthParams(request, auth, ctx)
	if result, getResultError := shopping_admin.NewProductServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindPageList(context.Background(), &authParams); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.WithData(resp.CreatePager(request, result.Count, result.List))
	}
}

// FindListByProductIds 根据商品id查询商品
func (p *product) FindListByProductIds(ctx *gin.Context) any {
	productParams := shopping_admin.ProductRequestParamMicroList{}
	if err := ctx.ShouldBindJSON(&productParams); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if len(productParams.ParamList) == 0 {
		return resp.Fail.WithMsg("id不能为空")
	}
	if result, getResultError := shopping_admin.NewProductServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindListByProductIds(context.Background(), &productParams); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Success.WithData(result.List)
	}
}
