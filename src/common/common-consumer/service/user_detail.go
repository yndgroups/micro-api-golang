package service

import (
	"context"
	"core/app"
	"core/config"
	"core/logger"
	"core/resp"
	"core/validate"
	"github.com/gin-gonic/gin"
	"protobuf/build/common"
)

var UserDetail = &userDetail{}

type userDetail struct {
	common.UserDetailServiceServer
}

// FindDetail 获取用户本人详细信息
func (ud *userDetail) FindDetail(ctx *gin.Context) any {
	if result, getResultError := common.NewUserDetailServiceClient(app.Nacos.CommonGrpcProvider()).FindById(context.Background(), &common.UserDetailIds{
		UserId: []string{ctx.Param("requestUserId")},
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		if result.Count > 0 {
			return resp.Back.Result(result.Count, common.BaseUserDetail{
				Phone:    result.Detail.Phone,
				RealName: result.Detail.RealName,
				Email:    result.Detail.Email,
			})
		} else {
			return resp.Fail.WithMsg("未查询到您的信息，请完善")
		}
	}
}

// UpdateDetail 更新用户本人信息
func (u *userDetail) UpdateDetail(ctx *gin.Context) any {
	baseUserInfo := common.BaseUserDetail{}
	if err := ctx.ShouldBindJSON(&baseUserInfo); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&baseUserInfo); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	if result, getResultError := common.NewUserDetailServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &common.UserDetail{
		UserId:   ctx.Param("requestUserId"),
		RealName: baseUserInfo.RealName,
		Phone:    baseUserInfo.Phone,
		Email:    baseUserInfo.Email,
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Back.Update(result.Count)
	}
}
