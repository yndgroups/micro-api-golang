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

var ProductSpec = productSpec{}

type productSpec struct{}

// Create 商品规格参数新增
func (ps *productSpec) Create(ctx *gin.Context) any {
	productSpec := shopping_admin.ProductSpec{}
	if err := ctx.ShouldBindJSON(&productSpec); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&productSpec); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	productSpec.CreateBy = ctx.Param("requestUserId")
	if result, getResultError := shopping_admin.NewProductSpecServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Create(context.Background(), &productSpec); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Insert(result.Count)
	}
}

// Update 商品规格参数修改
func (ps *productSpec) Update(ctx *gin.Context) any {
	productSpec := shopping_admin.ProductSpec{}
	if err := ctx.ShouldBindJSON(&productSpec); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&productSpec); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	productSpec.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := shopping_admin.NewProductSpecServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Update(context.Background(), &productSpec); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}

// Delete 商品规格参数删除
func (ps *productSpec) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}

	if result, getResultError := shopping_admin.NewProductSpecServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Delete(context.Background(), &shopping_admin.ProductSpecIds{
		Ids:    ids,
		UserId: ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Delete(result.Count)
	}
}

// FindById 查询商品规格参数
func (ps *productSpec) FindById(ctx *gin.Context) any {
	if result, getResultError := shopping_admin.NewProductSpecServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindById(context.Background(), &shopping_admin.ProductSpecIds{
		Ids: []string{ctx.Param("id")},
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

// FindListByProductCategoryId 根据productCategoryId查询规格列表
func (ps *productSpec) FindListByProductCategoryId(ctx *gin.Context) any {
	if result, getResultError := shopping_admin.NewProductSpecServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindListByProductCategoryId(context.Background(), &shopping_admin.ProductSpecIds{
		ProductCategoryId: ctx.Param("id"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.WithData(result.List)
	}
}

// FindPageList 查询商品规格参数分页列表
func (ps *productSpec) FindPageList(ctx *gin.Context) any {
	request := shopping_admin.ProductSpecPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	if result, getResultError := shopping_admin.NewProductSpecServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}
