package configureation

import (
	"core/logger"
	"errors"
	"protobuf/build/common"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
)

// InitWechatConfig 初始化微信支付配置
func InitWechatConfig(config *common.Config) (*wechat.ClientV3, error) {
	// NewClientV3 初始化微信客户端 v3
	//	mchid：商户ID 或者服务商模式的 sp_mchid
	// 	serialNo：商户证书的证书序列号
	//	apiV3Key：apiV3Key，商户平台获取
	//	privateKey：私钥 apiclient_key.pem 读取后的内容
	if client, err := wechat.NewClientV3(config.MchId, config.SerialNo, config.ApiV3Key, config.PrivateKey); err != nil {
		logger.SugarLogger.Errorf("微信初始化失败，失败原因：%v", err.Error())
		return nil, errors.New("微信初始化失败！")
	} else {
		// 启用自动同步返回验签，并定时更新微信平台API证书
		/*err = client.AutoVerifySign()
		if err != nil {
			logger.SugarLogger.Errorf("启用自动同步返回验签，并定时更新微信平台API证书失败，失败原因：%v", err.Error())
		}*/
		// 打开Debug开关，输出日志，默认是关闭的
		client.DebugSwitch = gopay.DebugOn
		return client, nil
	}
}
