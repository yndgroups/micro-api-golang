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
	"encoding/json"
	"errors"
	"protobuf/build/common"
	"protobuf/build/global"

	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/HttpRequest"
)

var User = &user{}

type user struct{}

// Create 创建用户
func (u *user) Create(ctx *gin.Context) any {
	user := common.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&user); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	user.UserId = redis.GET.GetPrimaryKey("SYS_COMMON_REQUEST_COUNT")
	user.CreateBy = ctx.Param("requestUserId")
	if result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &common.CreateUserDTO{
		User: &user,
		PowerParma: &common.PowerParma{
			AppId: ctx.Param("appId"),
		},
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		if result.Count > 0 {
			return resp.Success.WithData(result)
		}
		return resp.NotData
	}
}

// Update 更新用户信息
func (u *user) Update(ctx *gin.Context) any {
	user := common.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&user); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	user.UpdateBy = ctx.Param("requestUserId")
	if result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).Update(context.Background(), &user); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.WithData(result)
	}
}

// Delete 批或单条删除用户信息
func (u *user) Delete(ctx *gin.Context) any {
	ids := model.Ids{}
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	if result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).Delete(context.Background(), &common.UserIds{
		UserId:         ids,
		OperatorUserid: ctx.Param("requestUserId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.WithData(result)
	}
}

// FindPageList 查询用户分页列表
func (u *user) FindPageList(ctx *gin.Context) any {
	request := common.UserPageAuthQuery{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	// 校验请求参数
	if validateMsg, err := validate.Validated(&request); err != nil {
		return resp.ValidateErr.WithData(validateMsg)
	}
	request.Auth = &global.Auth{
		AppId: ctx.Param("appId"),
	}
	if result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).FindPageList(context.Background(), &request); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.WithData(result)
	}
}

// FindBalance 获取用户本人信息
func (u *user) FindBalance(ctx *gin.Context) any {
	if result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).FindBalance(context.Background(), &common.UserIds{
		UserId: []string{ctx.Param("requestUserId")},
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		return resp.Success.WithData(result.Balance)
	}
}

// FindWeChatWebUserInfo 查询微信信息
func (u *user) FindWeChatWebUserInfo(ctx *gin.Context) any {
	if wxUserInfo, getUserInfoErr := u.getWeChatWebUserInfo(ctx.Param("requestUserId")); getUserInfoErr != nil {
		return resp.Fail.WithMsg(getUserInfoErr.Error())
	} else {
		// 同步商城信息
		/* user := order.User{
			UserId:    ctx.Param("requestUserId"),
			NickName:  wxUserInfo.Nickname,
			Province:  wxUserInfo.Province,
			City:      wxUserInfo.City,
			Avatar:    wxUserInfo.Headimgurl,
			Gender:    wxUserInfo.Sex,
			LoginType: "h5",
		}
		if _, getResultError := order.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).UpdateUser(context.Background(), &user); getResultError != nil {
			logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().OrderProvider, getResultError.Error())
			return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
		} */
		return resp.Success.WithData(wxUserInfo)
	}
}

// getWeChatWebUserInfo 获取微信用户信息
func (u *user) getWeChatWebUserInfo(userId string) (model.WeChatUser, error) {
	redisUserKey := config.Global.ServerConfig().LoginPrefix + userId
	var weChatUser model.WeChatUser
	if loginUserJson, err := redis.GET.GetJson(redisUserKey); err != nil {
		return weChatUser, errors.New("登录信息已过期，请重新登录！")
	} else {
		var loginUser model.LoginUser
		if jsonErr := json.Unmarshal(loginUserJson, &loginUser); jsonErr != nil {
			return weChatUser, errors.New("用户信息json转换失败")
		}
		// https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html
		// -微信网页开发- 微信网页开发第四步 第四步：拉取用户信息(需scope为 snsapi_userinfo)
		// -微信网页开发- https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
		// - 用户管理- GET https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
		url := "https://api.weixin.qq.com/sns/userinfo?access_token=" + loginUser.WxAccessToken + "&openid=" + loginUser.OpenId + "&lang=zh_CN"
		if wxRequest, getErr := HttpRequest.Get(url); getErr != nil {
			return weChatUser, errors.New("微服务器请求失败！")
		} else {
			if body, err := wxRequest.Body(); err != nil {
				return weChatUser, errors.New("微用户信息数据获取失败！")
			} else {
				if err := json.Unmarshal(body, &weChatUser); err != nil {
					return weChatUser, errors.New("微用户信息解析失败！")
				} else {
					return weChatUser, nil
				}
			}
		}
	}
}

// FindById 查询用户信息
func (u *user) FindById(ctx *gin.Context) any {
	return nil
	/*var user common.User
	if res := db.DB.First(&user, ctx.Param("id")); res.RowsAffected > 0 {
		return resp.Success.WithData(user)
	} else {
		return resp.Fail.WithMsg("用户名不存在!")
	}*/
}

// Logout 退出登录
func (u *user) Logout(ctx *gin.Context) any {
	userId, _ := ctx.Params.Get("requestUserId")
	redisKey := config.Global.ServerConfig().LoginPrefix + userId
	_ = redis.DELETE.DelByKey(redisKey)
	loginUser, _ := redis.GET.GetJson(redisKey)
	if len(loginUser) == 0 {
		return resp.Success.WithMsg("退出成功!")
	} else {
		return resp.Fail.WithMsg("退出失败!")
	}
}

// FindUserIdByUserParam 根据指定参数获取用户信息
func (u *user) FindUserIdByUserParam(ctx *gin.Context) any {
	params := common.UserParam{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		return resp.Fail.WithData(err.Error())
	}
	if result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).FindByUserParam(context.Background(), &params); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", "common-provider", getResultError.Error())
		return resp.Fail.WithMsg("服务生产者发生错误,请联系管理员进行检查")
	} else {
		if result.Msg != "" {
			return resp.Fail.WithMsg(result.Msg)
		}
		userInfo := map[string]string{"userId": result.Detail.UserId, "email": result.Detail.Email}
		return resp.Back.Result(result.Count, userInfo)
	}
}
