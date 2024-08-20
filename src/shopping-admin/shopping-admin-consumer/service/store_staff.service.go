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
	"fmt"
	"protobuf/build/common"
	shoppingClient "protobuf/build/shopping_admin"

	"github.com/gin-gonic/gin"
)

var StoreStaff = &storeStaff{}

type storeStaff struct{}

// Create 商家新增
func (ss *storeStaff) Create(ctx *gin.Context) any {
	storeStaff := shoppingClient.StoreStaff{}
	if err := ctx.ShouldBindJSON(&storeStaff); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&storeStaff); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	// 查询用户的店铺信息拿到店铺id 临时写一个
	storeStaff.StoreId = "2022091323000000001" // 农品万联店铺
	// 调取公共服务获取用户基本信息
	result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).FindById(context.Background(), &common.UserIds{
		UserId: []string{storeStaff.UserId},
	})
	if getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	}
	if result.Count == 0 {
		return resp.Fail.WithMsg(fmt.Sprintf("用户ID为%v的员工信息不存在", storeStaff.UserId))
	} else {
		// storeStaff.RealName = result.Detail.RealName
		storeStaff.OpenId = result.Detail.OpenId
		storeStaff.Phone = result.Detail.Phone
		storeStaff.StaffId = redis.GET.GetPrimaryKey("SYS_SHOPP_ADMIN_REQUEST_COUNT")
		storeStaff.Email = result.Detail.Email
		storeStaff.CreateBy = ctx.Param("requestUserId")
	}

	if result, getResultError := shoppingClient.NewStoreStaffServiceClient(app.Nacos.ShoppingClientProvider()).Create(context.Background(), &storeStaff); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Insert(result.Count)
	}
}

// Update 商家修改
func (ss *storeStaff) Update(ctx *gin.Context) any {
	userId, _ := ctx.Params.Get("requestUserId")
	storeStaff := shoppingClient.StoreStaff{}
	if err := ctx.ShouldBindJSON(&storeStaff); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&storeStaff); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	// 查询用户的店铺信息拿到店铺id 临时写一个
	storeStaff.StoreId = "2022091323000000001" // 农品万联店铺
	result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).FindById(context.Background(), &common.UserIds{
		UserId: []string{storeStaff.UserId},
	})
	if getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	}
	if result.Count == 0 {
		return resp.Fail.WithMsg(fmt.Sprintf("用户ID为%v的员工信息不存在", storeStaff.UserId))
	} else {
		// storeStaff.RealName = result.Detail.RealName
		storeStaff.OpenId = result.Detail.OpenId
		storeStaff.Phone = result.Detail.Phone
		storeStaff.UpdateBy = userId
	}
	if result, getResultError := shoppingClient.NewStoreStaffServiceClient(app.Nacos.ShoppingClientProvider()).Update(context.Background(), &storeStaff); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Update(result.Count)
	}
}

// Delete 商家删除
func (ss *storeStaff) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := shoppingClient.NewStoreStaffServiceClient(app.Nacos.ShoppingClientProvider()).Delete(context.Background(), &shoppingClient.StoreStaffIds{
		StaffId: ids,
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Delete(result.Count)
	}
}

// FindById 查询商家
func (ss *storeStaff) FindById(ctx *gin.Context) any {
	if result, getResultError := shoppingClient.NewStoreStaffServiceClient(app.Nacos.ShoppingClientProvider()).FindById(context.Background(), &shoppingClient.StoreStaffIds{
		StaffId: []string{ctx.Param("id")},
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, result.Detail)
	}
}

// FindPageList 查询商家分页列表
func (ss *storeStaff) FindPageList(ctx *gin.Context) any {
	request := shoppingClient.StoreStaffPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	if result, getResultError := shoppingClient.NewStoreStaffServiceClient(app.Nacos.ShoppingClientProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().ShoppingAdminProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		return resp.Back.Result(result.Count, pager.NewPager{List: result.List, TotalCount: result.Count, PageIndex: request.PageIndex, PageSize: request.PageSize})
	}
}
