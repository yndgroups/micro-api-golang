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
	"protobuf/build/shopping_client"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var Cart = &cart{}

type cart struct{}

// Create 添加购物车
func (c *cart) Create(ctx *gin.Context) any {
	cart := shopping_client.Cart{}
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	cart.CartId = redis.GET.GetPrimaryKey("SYS_CART_COUNT")
	cart.CreateBy = ctx.Param("requestUserId")
	if result, getResultError := shopping_client.NewCartServiceClient(app.Nacos.ShoppingClientProvider()).Create(context.Background(), &cart); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.WithMsg("购物车添加成功")
	}
}

// Delete 根据商品id删除购物车
func (c *cart) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := shopping_client.NewCartServiceClient(app.Nacos.ShoppingClientProvider()).Delete(context.Background(), &shopping_client.CartIds{
		CartIds: ids,
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Delete(result.Count)
	}
	// redis
	/* ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if len(ids) == 0 {
		return resp.Fail.WithMsg("参数不能为空")
	}
	tempIds := model.Ids{}
	for _, v := range ids {
		if err := redis.DELETE.HashDel(fmt.Sprintf("cart_%s", ctx.Param("requestUserId")), v); err != nil {
			tempIds = append(tempIds, v)
		}
	}
	if len(tempIds) > 0 {
		return resp.DeleteFail.WithData(tempIds)
	}
	return resp.DeleteSuccess */
}

// FindCount 查询用户的购物车数量
func (c *cart) FindCount(ctx *gin.Context) any {
	if result, getResultError := shopping_client.NewCartServiceClient(app.Nacos.ShoppingClientProvider()).FindCount(context.Background(), &shopping_client.CartIds{
		UserId: ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			resp.Success.WithMsg(result.Msg)
		}
		return resp.Success.WithData(result.Count)
	}
	// redis
	/* cartInfo, redisGetErr := redis.GET.HashLen(fmt.Sprintf("cart_%v", ctx.Param("requestUserId")))
	if redisGetErr != nil {
		return resp.Fail.WithMsg(redisGetErr.Error())
	} else {
		return resp.Success.WithData(cartInfo)
	} */
}

// ChangeQuantity 修改购物车商品数量
func (c *cart) ChangeQuantity(ctx *gin.Context) any {
	if result, getResultError := shopping_client.NewCartServiceClient(app.Nacos.ShoppingClientProvider()).ChangeQuantity(context.Background(), &shopping_client.CartIds{
		UserId:   ctx.Param("requestUserId"),
		CartIds:  []string{ctx.Param("cartId")},
		Quantity: cast.ToInt64(ctx.Param("quantity")),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithDataAndMsg(result.Count, result.Msg)
		}
		if result.Count > 0 {
			return resp.Success.WithMsg("修改成功")
		} else {
			return resp.Fail.WithMsg("购物车数量修改失败")
		}
	}
}

// FindPageList 获取购物车列表
func (c *cart) FindPageList(ctx *gin.Context) any {
	request := shopping_client.CartPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	// 查询用户购物车
	request.UserId = ctx.Param("requestUserId")
	if result, getResultError := shopping_client.NewCartServiceClient(app.Nacos.ShoppingClientProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().ShoppingClientProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}

	// redis
	/* cartInfo, redisGetErr := redis.GET.HashValues(fmt.Sprintf("cart_%s", ctx.Param("requestUserId")))
	if redisGetErr != nil {
		return resp.Fail.WithMsg(redisGetErr.Error())
	} else {
		return resp.Success.WithData(cartInfo)
	} */
}
