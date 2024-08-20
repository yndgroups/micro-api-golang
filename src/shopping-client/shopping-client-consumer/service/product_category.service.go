package service

import (
	"context"
	"core/app"
	"core/config"
	"core/logger"
	"core/resp"
	"protobuf/build/global"
	"protobuf/build/shopping_client"

	"github.com/gin-gonic/gin"
)

var ProductCategory = &productCategory{}

type productCategory struct{}

// FindCategoryByParentId 根据parentId查询商品分类
func (pc *productCategory) FindCategoryByParentId(ctx *gin.Context) any {
	request := shopping_client.ProductCategoryAuthParams{
		ParentId: ctx.Param("parentId"),
		Auth: &global.Auth{
			UserId: ctx.Param("requestUserId"),
		},
	}
	if result, getResultError := shopping_client.NewProductCategoryServiceClient(app.Nacos.ShoppingClientProvider()).FindByParentId(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingClientProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, result.List)
	}
}

// 查询商品树转给结构分类
func (pc *productCategory) FindCategoryTreeByParentId(ctx *gin.Context) any {
	request := shopping_client.ProductCategoryAuthParams{
		ParentId: ctx.Param("parentId"),
		Auth: &global.Auth{
			UserId: ctx.Param("requestUserId"),
		},
	}
	if result, getResultError := shopping_client.NewProductCategoryServiceClient(app.Nacos.ShoppingClientProvider()).FindTreeByParentId(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingClientProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, pc.treeRecursion(result.TreeList, ctx.Param("parentId")))
	}
}

// 递归处理数据
func (pc *productCategory) treeRecursion(list []*shopping_client.ProductCategoryTreeList, parentId string) []*shopping_client.ProductCategoryTreeList {
	// 递归出传入的父级编码数据及非父级编码数据
	parentList := make([]*shopping_client.ProductCategoryTreeList, 0)
	childList := make([]*shopping_client.ProductCategoryTreeList, 0)
	for _, v := range list {
		if v.ParentId == parentId {
			parentList = append(parentList, v)
		} else {
			childList = append(childList, v)
		}
	}
	// 将非父级别数据进行分组
	childMapList := make(map[string][]*shopping_client.ProductCategoryTreeList)
	for _, v := range childList {
		t := childMapList[v.ParentId]
		t = append(t, v)
		childMapList[v.ParentId] = t
	}
	// 最总结果
	for i, v := range parentList {
		child := childMapList[v.CategoryId]
		if len(child) > 0 {
			t := pc.treeRecursion(childList, v.CategoryId)
			parentList[i].Children = t
		} else {
			parentList[i].Children = child
		}
	}
	return parentList
}
