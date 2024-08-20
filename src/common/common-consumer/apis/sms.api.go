package apis

import (
	"common/common-consumer/service"
	"net/http"

	_ "core/req"
	_ "core/resp"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/gin-gonic/gin"
)

var Sms = &sms{}

type sms struct{}

// Send
// @Tags 短信服务
// @Summary 发送验证码
// @Description 主要用于发送验证码
// @Param accessToken header string true "授权令牌" default(Bearer token)
// @Param phone path string true "手机号码"
// @Param type path string true "类型"
// @Success 200 {object} resp.Response "请求成功"
// @Failure 400,500 {object} resp.Response "请求异常"
// @Router /sms/{phone}/{type} [get]
func (s *sms) Send(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Sms.SmsSend(ctx))
}

/*
	// CreateClient

// 使用AK&SK初始化账号Client
//	@param accessKeyId
//	@param accessKeySecret
//	@return Client
//	@throws Exception
*/
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	host := "dysmsapi.aliyuncs.com"
	config.Endpoint = &host
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

/* func testMain() (_err error) {
	client, _err := CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		ResourceOwnerAccount: tea.String("your_value"),
		ResourceOwnerId:      tea.Int64(1),
		PhoneNumbers:         tea.String("your_value"),
		SignName:             tea.String("your_value"),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
*/
