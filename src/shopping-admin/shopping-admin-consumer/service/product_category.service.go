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
	"protobuf/build/shopping_admin"

	"github.com/gin-gonic/gin"
)

var ProductCategory = &productCategory{}

type productCategory struct{}

// Create 商品分类新增
func (p *productCategory) Create(ctx *gin.Context) any {
	productCategory := shopping_admin.ProductCategory{}
	if err := ctx.ShouldBindJSON(&productCategory); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&productCategory); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	productCategory.CreateBy = ctx.Param("requestUserId")
	productCategory.ProductCategoryId = redis.GET.GetPrimaryKey("SYS_SHOPP_ADMIN_REQUEST_COUNT")
	if result, getResultError := shopping_admin.NewProductCategoryServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Create(context.Background(), &productCategory); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Insert(result.Count)
	}
}

// Update 商品分类修改
func (p *productCategory) Update(ctx *gin.Context) any {
	productCategory := shopping_admin.ProductCategory{}
	if err := ctx.ShouldBindJSON(&productCategory); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&productCategory); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	productCategory.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := shopping_admin.NewProductCategoryServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Update(context.Background(), &productCategory); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Update(result.Count)
	}
}

// Delete 商品分类删除
func (p *productCategory) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// return db.BatchDeleteByIds(ids, enum.ProductCategory.TableName, enum.ProductCategory.PrimaryKey, ctx.Param("requestUserId"))
	if result, getResultError := shopping_admin.NewProductCategoryServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).Delete(context.Background(), &shopping_admin.ProductCategoryIds{
		Id: ids,
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Delete(result.Count)
	}
}

// FindListByParentId 根据parentId查询商品分类
func (p *productCategory) FindListByParentId(ctx *gin.Context) any {
	parentId := ctx.Param("parentId")
	if result, getResultError := shopping_admin.NewProductCategoryServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindListByParentId(context.Background(), &shopping_admin.ProductCategoryIds{
		ParentId: parentId,
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Success.WithData(result.List)
	}
}

// FindTreeByParentId 查询商品树转给结构分类
func (p *productCategory) FindTreeByParentId(ctx *gin.Context) any {
	parentId := ctx.Param("parentId")
	if result, getResultError := shopping_admin.NewProductCategoryServiceClient(app.Nacos.ShoppingGrpcAdminProvider()).FindTreeByParentId(context.Background(), &shopping_admin.ProductCategoryIds{
		ParentId: parentId,
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Success.WithData(p.treeRecursion(result.List, parentId))
	}
}

// treeRecursion 递归处理分类
func (p *productCategory) treeRecursion(list []*shopping_admin.ProductCategoryListVO, parent string) []*shopping_admin.ProductCategoryListVO {
	// 递归出传入的父级编码数据及非父级编码数据
	parentList := make([]*shopping_admin.ProductCategoryListVO, 0)
	childList := make([]*shopping_admin.ProductCategoryListVO, 0)
	for _, v := range list {
		if v.ParentId == parent {
			parentList = append(parentList, v)
		} else {
			childList = append(childList, v)
		}
	}
	// 将非父级别数据进行分组
	childMapList := make(map[string][]*shopping_admin.ProductCategoryListVO)
	for _, v := range childList {
		t := childMapList[v.ParentId]
		t = append(t, v)
		childMapList[v.ParentId] = t
	}
	// 最总结果
	for i, v := range parentList {
		child := childMapList[v.ProductCategoryId]
		if len(child) > 0 {
			t := p.treeRecursion(childList, v.ProductCategoryId)
			parentList[i].Children = t
		} else {
			parentList[i].Children = child
		}
	}
	return parentList
}
