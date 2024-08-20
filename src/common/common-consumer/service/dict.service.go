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

var Dict = &dict{}

type dict struct{}

// Create 字典新增
func (d *dict) Create(ctx *gin.Context) any {
	dict := common.Dict{}
	if err := ctx.ShouldBindJSON(&dict); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&dict); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	dict.DictId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")
	dict.CreateBy = ctx.Param("requestUserId")

	if result, getResultError := common.NewDictServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &dict); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Insert(result.Count)
	}
}

// Update 字典修改
func (d *dict) Update(ctx *gin.Context) any {
	dict := common.Dict{}
	if err := ctx.ShouldBindJSON(&dict); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if validateMsg, err := validate.Validated(&dict); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	dict.UpdateBy = ctx.Param("requestUserId")

	if result, getResultError := common.NewDictServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &dict); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}

// Delete 字典删除
func (d *dict) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}

	if result, getResultError := common.NewDictServiceClient(app.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.DictIds{
		DictIds: ids,
		UserId:  ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Delete(result.Count)
	}
}

// FindById 字典查找
func (d *dict) FindById(ctx *gin.Context) any {

	if result, getResultError := common.NewDictServiceClient(app.Nacos.CommonGrpcProvider()).FindById(context.Background(), &common.DictIds{
		DictIds: []string{ctx.Param("id")},
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

// FindListByParentId 字典查找
func (d *dict) FindListByParentId(ctx *gin.Context) any {

	if result, getResultError := common.NewDictServiceClient(app.Nacos.CommonGrpcProvider()).FindListByParentId(context.Background(), &common.DictIds{
		ParentId: ctx.Param("parentId"),
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
