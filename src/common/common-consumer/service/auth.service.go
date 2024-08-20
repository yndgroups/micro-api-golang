package service

import (
	"context"
	"core/app"
	"core/config"
	"core/jwt"
	"core/logger"
	cm "core/model"
	"core/redis"
	"core/resp"
	"core/utils"
	"core/validate"
	"encoding/json"
	"errors"
	"fmt"
	"protobuf/build/common"
	"protobuf/build/global"
	"strings"
	"time"
	"unsafe"

	"github.com/blinkbean/dingtalk"

	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/HttpRequest"
)

var Auth = &auth{}

type auth struct{}

// SendMsg 测试消息发送
func (a *auth) SendMsg(ctx *gin.Context) any {
	// 文本发送
	// cli := dingtalk.InitDingTalkWithSecret("cd6af33e4ffd4756ec1aaf9e0e71e0e349db978cedb0e20e425a18641d47eb08", "SEC8c3059b47cfb77d769ccf7f7dace8ee6a441df22bb57f744b6e9bbf67e12a869")
	//方法一、 cli.SendTextMessage("您有一个新订单，测试", dingtalk.WithAtAll())

	// 方法二、按行写入数组，增强markdown的可读性
	// maps := make(map[string]string)
	// maps["productName"] = "五香辣椒面"
	// maps["price"] = "28.8"
	// maps["orderNo"] = "SP_2023090509351688"
	// maps["customerName"] = "杨大琼"
	// maps["phpne"] = "15085999726"
	// maps["address"] = "北京市房山区良乡大南关村"
	// a.sendMsgToDingTalk(maps)

	// 方法三
	//list := []string{"测试多个A测试多个A测试多个A测试多个A测试多个A测试多个A测试多个A测试多个A", "测试多个A测试多个A测试多个A测试多个A测试多个A测试多个A测试多个A测试多个A"}
	//cli := dingtalk.InitDingTalkWithSecret("cd6af33e4ffd4756ec1aaf9e0e71e0e349db978cedb0e20e425a18641d47eb08", "SEC8c3059b47cfb77d769ccf7f7dace8ee6a441df22bb57f744b6e9bbf67e12a869")
	//dm := dingtalk.DingMap()
	//dm.Set("新订单提醒", dingtalk.H2)
	//for i, v := range list {
	//	dm.Set(fmt.Sprintf("商品%d: $$ %v $$", i+1, v), dingtalk.RED)
	//	dm.Set(fmt.Sprintf("数  量: $$ %v $$", i+1), dingtalk.RED)
	//	dm.Set(fmt.Sprintf("价  格: $$ %v $$", i+1), dingtalk.RED)
	//	dm.Set("---", dingtalk.N)
	//}
	//dm.Set(fmt.Sprintf("客户姓名: %v", "大内密探00*"), dingtalk.N)
	//dm.Set(fmt.Sprintf("联系电话: %v", 15085999726), dingtalk.N)
	//dm.Set(fmt.Sprintf("收货地址: $$ %v $$", "北京市房山区良乡大南关村129号"), dingtalk.RED)
	//cli.SendMarkDownMessageBySlice("新订单提醒", dm.Slice())
	return resp.Success.WithData("发送成功")
}

// sendOrderMsgToDingTalk 发送消息到钉钉
func (w *auth) sendMsgToDingTalk(msgs map[string]string) {
	cli := dingtalk.InitDingTalkWithSecret("cd6af33e4ffd4756ec1aaf9e0e71e0e349db978cedb0e20e425a18641d47eb08", "SEC8c3059b47cfb77d769ccf7f7dace8ee6a441df22bb57f744b6e9bbf67e12a869")
	msg := []string{
		"### 新订单提醒",
		"---",
		"- 商品名称：<font color=#ff0000 size=6>" + msgs["productName"] + "</font>",
		"- 商品单价：<font color=#ff0000 size=6>" + msgs["price"] + "元</font>",
		"- 商品单号：<font color=#ff0000 size=6>" + msgs["orderNo"] + "</font>",
		"- 客户姓名：<font color=#ff0000 size=6>" + msgs["customerName"] + "</font>",
		"- 联系电话：<font color=#ff0000 size=6>" + msgs["phpne"] + "</font>",
		"- 收货地址：<font color=#ff0000 size=6>" + msgs["address"] + "</font>",
	}
	cli.SendMarkDownMessageBySlice("新订单提醒", msg, dingtalk.WithAtAll())
}

// CreateAnonymityToken 创建匿名token
func (a *auth) CreateAnonymityToken(ctx *gin.Context) any {
	if !strings.Contains(config.Global.ServerConfig().AppRoleUniqueIdentification, ctx.Param("appId")) {
		return resp.AuthorizedErr
	}
	ip := ctx.ClientIP()
	ipArr := strings.Split(ip, ".")
	ipSecret := "95"
	for i, v := range ipArr {
		switch i {
		case 0:
			ipSecret += v
		case 1:
			ipSecret += "27" + v
		case 2:
			ipSecret += "1317" + v
		case 3:
			ipSecret += "159" + v
		}
	}
	token, _ := jwt.CreateToken(jwt.TokenInfo{
		UserId:  "anonymous_" + ipSecret,
		AppId:   ctx.Param("appId"),
		SubId:   "",
		RegType: 0,
	})
	return resp.Success.WithData(token)
}

// Login 后台登录
func (a *auth) Login(ctx *gin.Context) any {
	// defer exception.Txception()()
	//注意该结构接收的内容
	loginUser := cm.LoginDTO{}
	ctx.ShouldBindJSON(&loginUser)
	// 查询数据库是否有用户信息
	// 查询用户信息
	userResult, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).FindByUserParam(context.Background(), &common.UserParam{
		UserName: loginUser.UserName,
		AppId:    ctx.Param("appId"),
		Email:    loginUser.Email,
		Phone:    loginUser.Phone,
	})

	if getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	}
	// 错误问题
	if userResult.Count == 0 {
		return resp.Fail.WithMsg(userResult.Msg)
	}
	//println(utils.HmacSha512(loginUser.UserPassword, user.Salt)) // 盐加密
	// 判断密码是否正确
	var user = userResult.Detail
	if user.UserPassword != utils.HmacSha512(loginUser.UserPassword, user.Salt) {
		return resp.Fail.WithMsg("用户密码不正确！")
	}
	// 验证是否有登陆后台权限
	if permission := a.isLoginPermission(user.UniqueIdentification, ctx.Param("appId")); !permission {
		return resp.Fail.WithMsg("该账号权限不支持后台登陆！")
	}
	// 创建token
	token, createTokenErr := jwt.CreateToken(
		jwt.TokenInfo{
			UserId:  user.UserId,
			AppId:   ctx.Param("appId"),
			RegType: user.RegType,
			SubId:   "",
		})
	if createTokenErr != nil {
		return resp.Fail.WithMsg(fmt.Sprintf("token创建失败,失败原因：%v", createTokenErr.Error()))
	}

	// 存储后台用户登陆redisKey
	var redisKey = fmt.Sprintf("%v%s_%s", config.Global.ServerConfig().LoginPrefix, ctx.Param("appId"), user.UserId)
	// 后台登陆保持不间断两小时在线
	redisErr := redis.SET.Set(redisKey, token, config.Global.RedisConfig().Expire)
	if redisErr != nil {
		return resp.Fail.WithMsg(fmt.Sprintf("数据存储失败,失败原因：%v", redisErr.Error()))
	}
	// 存储权限值
	var redisPowerKey = fmt.Sprintf("%v%s_%s", config.Global.ServerConfig().PowerPrefix, ctx.Param("appId"), user.UserId)
	// 后台登陆保持不间断两小时在线，因此在此也设置两小时保持与用户信息一致
	redisPowerErr := redis.SET.Set(redisPowerKey, userResult.PowerSign, config.Global.RedisConfig().Expire)
	if redisPowerErr != nil {
		return resp.Fail.WithMsg(fmt.Sprintf("数据存储失败,失败原因：%v", redisPowerErr.Error()))
	}
	// 返回登录信息
	var inMap = make(map[string]interface{})
	inMap["userName"] = user.UserName
	inMap["userId"] = user.UserId
	inMap["token"] = token
	return resp.Success.WithData(inMap)
}

// Logout 用户退出登陆
func (a *auth) Logout(ctx *gin.Context) any {
	redis.DELETE.DelByKey(ctx.Param("loginRedisKey"))
	redis.DELETE.DelByKey(ctx.Param("powerKey"))
	loginUser, _ := redis.GET.GetJson(ctx.Param("loginRedisKey"))
	if len(loginUser) == 0 {
		return resp.Success.WithMsg("退出成功!")
	} else {
		return resp.Fail.WithMsg("退出失败!")
	}
}

// RefreshPermissions 用户刷新权限
func (a *auth) RefreshPermissions(ctx *gin.Context) any {
	result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).FindPermissions(context.Background(), &common.UserParam{
		UserId: ctx.Param("requestUserId"),
		AppId:  ctx.Param("appId"),
	})
	if getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	}
	if result.PowerSign == "" {
		return resp.Fail.WithMsg("权限信息查询失败")
	}
	redisPowerErr := redis.SET.Set(ctx.Param("powerKey"), result.PowerSign, config.Global.RedisConfig().WeChatExpire)
	if redisPowerErr != nil {
		logger.SugarLogger.Warnf("权限数据更新失败: %v", redisPowerErr.Error())
		return resp.Fail.WithMsg("权限数据更新失败")
	}

	println(ctx.Param("appId") + "_" + ctx.Param("requestUserId"))
	redis.DELETE.HashDel("power_change", ctx.Param("appId")+"_"+ctx.Param("requestUserId"))
	return resp.Success.WithData(result.PowerSign)
}

// WechatLogin 微信小程序登录
func (a *auth) WechatLogin(ctx *gin.Context) any {
	wechatLoginDto := cm.WechatLoginDTO{}
	ctx.ShouldBindJSON(&wechatLoginDto)
	// 校验请求参数
	if validateMsg, err := validate.Validated(wechatLoginDto); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	wechatLoginDto.RegType = 4
	wechatLoginDto.AppId = ctx.Param("appId")
	return a.wechatLogins(wechatLoginDto)
}

// wechatLogins 微信小程序或公众号公共登录方法提取
func (a *auth) wechatLogins(wechatLoginDto cm.WechatLoginDTO) any {
	// 查询服务提供商配置
	confId := ""
	if wechatLoginDto.RegType == 3 {
		confId = config.Global.ServerConfig().WechatConfId
	} else if wechatLoginDto.RegType == 4 {
		confId = config.Global.ServerConfig().AppletConfId
	} else {
		return resp.Fail.WithMsg("登陆账号类型不正确")
	}
	conf, configErr := a.getConfig(&common.ConfigIds{
		ConfIds: []string{confId},
		AppId:   wechatLoginDto.AppId,
	})
	if configErr != nil {
		return resp.Fail.WithMsg(configErr.Error())
	}
	if conf == nil {
		return resp.Fail.WithMsg("配置信息获取失败")
	}
	var url string
	if wechatLoginDto.RegType == 3 { //公众号登录
		url = fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", conf.OutsideAppId, conf.OutsideAppSecret, wechatLoginDto.WxCode)
		println("公众号登录 = url = " + url)
	}
	if wechatLoginDto.RegType == 4 { // 小程序登录
		url = fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", conf.OutsideAppId, conf.OutsideAppSecret, wechatLoginDto.WxCode)
		println("小程序登录 = url = " + url)
	}

	// 请求微信接口
	wxRequest, httpRequestErr := HttpRequest.Get(url)
	if httpRequestErr != nil {
		return resp.Fail.WithMsg(fmt.Sprintf("微信服务器请求失败, 失败原因：%v", httpRequestErr.Error()))
	}

	// 获取微信数据
	wxResp, wxRequestErr := wxRequest.Body()
	if wxRequestErr != nil {
		return resp.Fail.WithMsg("微信数据获取失败！")
	}

	// 解析微信返回信息
	//str := "eyJzZXNzaW9uX2tleSI6ImVpZmQrcWIxcm1aT3ZGRm9vU2FKbWc9PSIsIm9wZW5pZCI6Im85X1JTNDY5ZXZUM1Y4OEN0UmNGdjg4N05CSlEifQ=="
	wxBack := (*string)(unsafe.Pointer(&wxResp))

	// 解析微信信息
	var wechatRequestBack cm.WechatRequestBack
	jsonParaseErr := json.Unmarshal([]byte(*wxBack), &wechatRequestBack)
	if jsonParaseErr != nil {
		return resp.Fail.WithData("微信返回信息解析失败，失败原因：" + jsonParaseErr.Error())
	}

	// 解析成功之后进行判断
	if wechatRequestBack.OpenId != "" {
		// 根据openId查找是否已经注册
		result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).FindByUserParam(context.Background(), &common.UserParam{
			OpenId:  wechatRequestBack.OpenId,
			RegType: wechatLoginDto.RegType,
			AppId:   conf.AppId,
		})
		if getResultError != nil {
			logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().CommonProvider, getResultError.Error())
			return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
		}
		// return resp.Back.WithData(result)
		if result.Count > 0 { // 用户存在直接登录并返回 token
			var user = result.Detail
			if user.Status == 2 {
				return resp.Fail.WithData("该账号被禁用！禁用原因：" + user.ReasonsProhibition)
			}
			// 由于access_token拥有较短的有效期，当access_token超时后，可以使用refresh_token进行刷新，
			// refresh_token有效期为30天，当refresh_token失效之后，需要用户重新授权。
			token, createTokenErr := jwt.CreateToken(jwt.TokenInfo{
				UserId:  user.UserId,
				AppId:   wechatLoginDto.AppId,
				RegType: wechatLoginDto.RegType,
				SubId:   result.Detail.OpenId,
			})
			if createTokenErr != nil {
				return resp.Fail.WithMsg(fmt.Sprintf("token创建失败，失败原因： %v", createTokenErr.Error()))
			}
			// 保存用户信息
			return a.saveLoginUserInfo(&common.UserDetail{
				UserId:    user.UserId,
				NickName:  wechatLoginDto.NickName,
				Province:  wechatLoginDto.Province,
				City:      wechatLoginDto.City,
				Avatar:    wechatLoginDto.AvatarUrl,
				Gender:    wechatLoginDto.Gender,
				LoginType: "applet",
			}, token, wechatLoginDto.AppId, result.PowerSign)

		} else { // 用户不存在立即注册并返回token
			userId := redis.GET.GetPrimaryKey("USER_REGISTER_COUNT")
			// 保存用户信息并初始化角色信息
			result, getResultError := common.NewUserServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), &common.CreateUserDTO{
				User: &common.User{
					OpenId:   wechatRequestBack.OpenId,
					UserId:   userId,
					Salt:     fmt.Sprintf("%d", time.Now().Unix()),
					UserName: "wx_" + userId,
					RegType:  wechatLoginDto.RegType,
					Status:   1,
					CreateBy: userId,
				},
				PowerParma: &common.PowerParma{
					AppId: wechatLoginDto.AppId,
				},
			})
			if getResultError != nil {
				logger.SugarLogger.Errorf("登录失败，原因：%v", getResultError.Error())
				return resp.Fail.WithMsg("登录失败，请重新尝试登录！")
			}
			if result.Count > 0 {
				token, createTokenErr := jwt.CreateToken(jwt.TokenInfo{
					UserId:  userId,
					AppId:   wechatLoginDto.AppId,
					RegType: wechatLoginDto.RegType,
					SubId:   wechatRequestBack.OpenId,
				})
				if createTokenErr != nil {
					logger.SugarLogger.Errorf("token创建失败,原因：%v", createTokenErr.Error())
					return resp.Fail.WithMsg("token创建失败")
				}
				// 保存更新用户
				return a.saveLoginUserInfo(&common.UserDetail{
					UserId:    userId,
					NickName:  wechatLoginDto.NickName,
					Province:  wechatLoginDto.Province,
					City:      wechatLoginDto.City,
					Avatar:    wechatLoginDto.AvatarUrl,
					Gender:    wechatLoginDto.Gender,
					LoginType: "applet",
				}, token, wechatLoginDto.AppId, result.PowerSign)
			} else {
				return resp.Fail.WithMsg("登录失败，请重新尝试登录！")
			}
		}
	} else {
		return resp.Fail.WithData(fmt.Sprintf("微信调用失败，失败原因:%v", utils.WeChatAccessTokenErrorMessage(wechatRequestBack.Errcode)))
	}
}

// WechatWebLogin 微信公众号登录
func (a *auth) WechatWebLogin(ctx *gin.Context) any {
	wechatLoginDto := cm.WechatLoginDTO{}
	ctx.ShouldBindJSON(&wechatLoginDto)
	// 校验请求参数
	if validateMsg, err := validate.Validated(&wechatLoginDto); err != nil {
		return resp.ErrParam.WithData(validateMsg)
	}
	wechatLoginDto.RegType = 3
	return a.wechatLogins(wechatLoginDto)
}

// 保存用户登陆信息
func (a *auth) saveLoginUserInfo(userInfo *common.UserDetail, token string, appId string, powerSign string) any {
	result, getResultError := common.NewUserDetailServiceClient(app.Nacos.CommonGrpcProvider()).Create(context.Background(), userInfo)
	if getResultError != nil {
		logger.SugarLogger.Errorf("接口地址：【小程序登陆接口】,服务生产者:【%v】发生错误,错误信息：%v", config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("登陆失败")
	}
	if result.Count >= 0 {
		// 微信登陆存储redisKey
		var redisKey = config.Global.ServerConfig().LoginPrefix + appId + "_" + userInfo.UserId
		// 微信登陆存储7天之内不在进行登陆毫秒存储
		if redisErr := redis.SET.Set(redisKey, token, config.Global.RedisConfig().WeChatExpire); redisErr != nil {
			logger.SugarLogger.Errorf("登录失败，数据存储失败,失败原因：%v", redisErr)
			return resp.Fail.WithMsg(fmt.Sprintf("数据存储失败,失败原因：%v", redisErr))
		}
		// 权限值存储redisPowerKey
		var redisPowerKey = fmt.Sprintf("%v%s_%s", config.Global.ServerConfig().PowerPrefix, appId, userInfo.UserId)
		// 存储权限值，与微信登陆用户的时间保持一致
		redisPowerErr := redis.SET.Set(redisPowerKey, powerSign, config.Global.RedisConfig().WeChatExpire)
		if redisPowerErr != nil {
			return resp.Fail.WithMsg(fmt.Sprintf("数据存储失败,失败原因：%v", redisPowerErr.Error()))
		}
		return resp.Success.WithData(map[string]string{"token": token})
	}
	return resp.Fail.WithMsg("登陆失")
}

// getConfig 获取配置
func (a *auth) getConfig(configExpandParam *common.ConfigIds) (*common.Config, error) {
	// 优先从redis取出配置信息，如果没有再去数据库查询
	redisConfigKey := config.Global.ServerConfig().ConfigPrefix + configExpandParam.ConfIds[0]
	configJson, getConfigErr := redis.GET.GetJson(redisConfigKey) // redis获取配置
	var conf *common.Config
	if getConfigErr != nil {
		// 根据openId查找是否已经注册
		result, getResultError := common.NewConfigServiceClient(app.Nacos.CommonGrpcProvider()).FindByConfigExpandParam(context.Background(), configExpandParam)
		if getResultError != nil {
			logger.SugarLogger.Errorf("服务生产者【%v】发生错误,错误信息:%v", config.Global.ServerConfig().CommonProvider, getResultError.Error())
			return nil, errors.New("服务生产者调用异常,请联系管理员进行检查！")
		}
		if result.Count > 0 {
			// 和微信登陆存储时间一直保持最大有效性
			redis.SET.SetJson(redisConfigKey, result.Detail, config.Global.RedisConfig().WeChatExpire)
			conf = result.Detail
		} else {
			return nil, errors.New("未查询到您的配置信息，请去系统进行添加或修改！")
		}
	} else {
		// 读取redis数据
		if jsonErr := json.Unmarshal(configJson, &conf); jsonErr != nil {
			return nil, errors.New("配置文件json解析失败！")
		}
	}
	if conf.ConfId == "" {
		return nil, errors.New("没有查询到相关配置！")
	}
	return conf, nil
}

// 角色关联菜单
func (a *auth) CreateRoleMenu(ctx *gin.Context) any {
	roleMenus := make([]*common.RoleMenu, 0)
	if err := ctx.ShouldBindJSON(&roleMenus); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if len(roleMenus) == 0 {
		return resp.Fail.WithMsg("列表参数不能为空")
	}
	createBy := ctx.Param("requestUserId")
	createRoleMenuArr := make([]*common.RoleMenu, 0)
	for i, v := range roleMenus {
		if v.MenuId == "" || v.RoleId == "列表参数不能为空" {
			return resp.Fail.WithMsg(fmt.Sprintf("列表下标为%v的菜单或角色id不能为空", i+1))
		} else {
			createRoleMenuArr = append(createRoleMenuArr, &common.RoleMenu{MenuId: v.MenuId, RoleId: v.RoleId, CreateBy: createBy})
		}
	}

	if result, getResultError := common.NewAuthClient(app.Nacos.CommonGrpcProvider()).CreateRoleMenu(context.Background(), &common.RoleMenuList{
		List: createRoleMenuArr,
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		a.addPowerChangeUser(result.UserIds, ctx.Param("appId"))
		return resp.Back.WithMsg("菜单授权成功")
	}
}

// CreateRoleFunc 角色关联功能
func (a *auth) CreateRoleFunc(ctx *gin.Context) any {
	roleFuncs := make([]*common.RoleFunc, 0)
	if err := ctx.ShouldBindJSON(&roleFuncs); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	// 校验请求参数
	if len(roleFuncs) == 0 {
		return resp.Fail.WithMsg("列表参数不能为空")
	}
	createBy := ctx.Param("requestUserId")
	createRoleFuncArr := make([]*common.RoleFunc, 0)
	for i, v := range roleFuncs {
		if v.FuncId == "" || v.RoleId == "列表参数不能为空" {
			return resp.Fail.WithMsg(fmt.Sprintf("列表下标为%v的菜单或角色id不能为空", i+1))
		} else {
			createRoleFuncArr = append(createRoleFuncArr, &common.RoleFunc{FuncId: v.FuncId, RoleId: v.RoleId, CreateBy: createBy})
		}
	}

	if result, getResultError := common.NewAuthClient(app.Nacos.CommonGrpcProvider()).CreateRoleFunc(context.Background(), &common.RoleFuncList{
		List: createRoleFuncArr,
	}); getResultError != nil {
		logger.SugarLogger.Errorf("接口地址: %v, 服务生产者: %v发生错误,错误信息: %v", ctx.FullPath(), config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		a.addPowerChangeUser(result.UserIds, ctx.Param("appId"))
		return resp.Back.WithMsg("菜单授权成功")
	}
}

// 将权限变更影响的用户添加到列表中
func (a *auth) addPowerChangeUser(userIds []string, appId string) error {
	if len(userIds) > 0 {
		appUserIds, _ := redis.GET.Keys(config.Global.ServerConfig().LoginPrefix + appId + "_*")
		// 存在有登陆信息的就将其加入需要刷新的缓存里
		if len(appUserIds) > 0 {
			appUserIdsStr := strings.Join(appUserIds, ",")
			for _, userId := range userIds {
				addKey := appId + "_" + userId
				if strings.Contains(appUserIdsStr, addKey) {
					// 通知用户进行权限变更
					// redis.HSet("isResetAuth", setKey, "1")
					// 删除用户已经存储的权限
					redis.SET.HashSet("power_change", addKey, "1")
					// res := redis.SET.HashSet("power_change", addKey, "1")
					// if res.Error() != "" {
					// 	return errors.New("redis存储权限变更信息失败")
					// }
				}
			}
		}
	}
	return nil
}

// 用户关联角色
func (a *auth) CreateUserRole(ctx *gin.Context) any {
	userRoleList := make([]*common.UserRole, 0)
	if err := ctx.ShouldBindJSON(&userRoleList); err != nil {
		return resp.Fail.WithMsg(err.Error())
	}
	createBy := ctx.Param("requestUserId")
	if len(userRoleList) == 0 {
		return resp.Fail.WithMsg("列表参数不能为空")
	}
	for i, v := range userRoleList {
		if v.UserId == "" || v.RoleId == "" {
			return resp.Fail.WithMsg(fmt.Sprintf("列表下标为%v的某项参数值为空", i+1))
		} else {
			v.CreateBy = createBy
			v.AppId = ctx.Param("appId")
		}
	}

	if result, getResultError := common.NewAuthClient(app.Nacos.CommonGrpcProvider()).CreateUserRole(context.Background(), &common.UserRoleList{
		List: userRoleList,
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发NewAuthClient生错误：错误信息：%v", config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Count > 0 {
			a.addPowerChangeUser([]string{userRoleList[0].UserId}, ctx.Param("appId"))
		}
		return resp.Back.WithMsg("角色授权成功")
	}
}

// 根据当前应用id获取新的token，用于新应用数据获取或这切换应用之间进行切换
func (a *auth) GetToken(ctx *gin.Context) any {
	if result, getResultError := common.NewAuthClient(app.Nacos.CommonGrpcProvider()).GetToken(context.Background(), &global.Auth{
		AppId: ctx.Param("appId"),
	}); getResultError != nil {
		logger.SugarLogger.Errorf("服务生产者【%v】发NewAuthClient生错误：错误信息：%v", config.Global.ServerConfig().CommonProvider, getResultError.Error())
		return resp.Fail.WithMsg("服务生产者调用异常,请联系管理员进行检查")
	} else {
		if result.Count > 0 {
			// 通知用户进行权限变更
			return resp.Success.WithData(result)
		}
		return resp.Fail.WithMsg(result.Msg)
	}
}

// 判断是否有登陆后台的权限
// uniqueIdentification 数据查出
func (a *auth) isLoginPermission(uniqueIdentification string, appId string) bool {
	b := false
	strArr := strings.Split(uniqueIdentification, ",")
	// 如果权限标识不存在，无权限登陆
	if len(strArr) == 0 {
		return b
	}
	if config.Global.ServerConfig().AppRoleUniqueIdentification != "" {
		appStrArr := strings.Split(config.Global.ServerConfig().AppRoleUniqueIdentification, "|")
		if len(appStrArr) > 0 {
			for i := 0; i < len(appStrArr); i++ {
				// 匹配上appId 不包含在对于配置里的唯一标识不能进行后台登陆
				if strings.Contains(appStrArr[0], appId) {
					appIdInfo := strings.Split(appStrArr[i], "-")
					// 配额里的关系不存在也直接无权限登陆
					if len(appIdInfo) == 0 {
						continue
					}
					for j := 0; j < len(strArr); j++ {
						// 如果后台配置的权限标识存在，可登陆
						if strings.Contains(appIdInfo[1], strArr[j]) {
							b = true
							break
						}
					}
				}
				// 内层循环取出结果之后为 true 直接终止整个循环
				if b {
					break
				}
			}
			return b
		}
		return b
	}
	return b
}
