package redis

type Config struct {
	// 配置id
	ConfId string `gorm:"primary_key" json:"confId"`
	// 应用名称
	AppName string `json:"appName"`
	// 应用id
	AppId string `json:"appId"`
	// 来源id
	FromAppId string `json:"fromAppId"`
	// 密钥
	AppSecret string `json:"appSecret"`
	// 支付商户号
	MchId string `json:"mchId"`
	//
	PaterNerKey string `json:"paterNerKey"`
	// 电话号码
	Phone string `json:"phone"`
	// 邮箱
	Email string `json:"email"`
	// 异步回调地址
	NotifyUrl string `json:"notifyUrl"`
	// 同步回调地址
	ReturnUrl string `json:"returnUrl"`
	// 公众号用于调用微信JS接口的临时票据
	JsApiTicket string `json:"jsApiTicket"`
	// Token可以由开发者任意填写，用作生成签名（该Token会配置在服务器中和接口URL中包含的Token进行比对，从而验证安全性）。
	Token string `json:"token"`
	// 默认支付金额
	Amount float64 `json:"amount"`
	// 消息格式化（比如：JSON）
	MsgDataFormat string `json:"msgDataFormat"`
	// EncodingAESKey由开发者手动填写或随机生成，将用作消息体加解密密钥。（43位）
	AesKey string `json:"aesKey"`
	// 时间戳
	Timestamp int64 `json:"timestamp"`
} // @name Config

type LoginUser struct {
	// 用户id
	UserId string `json:"userId"`
	// 用户名
	UserName string `json:"userName"`
	// 用户密码
	UserPassword string `json:"userPassword"`
	// 加密盐
	Salt string `json:"salt"`
	// 用户电话
	Phone string `json:"phone"`
	// 用户邮箱登录
	Email string `json:"email"`
	// 微信OpenId
	OpenId string `json:"OpenId"`
	// 禁用原因
	ReasonsProhibition string `json:"reasonsProhibition"`
	// 账号类型
	AccountType int8 `json:"accountType"`
	// 账号状态（1：启用（审核通过），2：禁用，3：待审核）
	Status int8 `json:"status"`
	// 帐号来源分类（1：后台端：2前端：3: 微信）
	RegType int8 `json:"regType"`
	// 帐号来源某个应用的appId
	SourceId string `json:"sourceId"`
	// 通过code登录时获取的AccessToken
	WxAccessToken string `json:"wxAccessToken"`
	// 权限值
	Permission string `json:"permission"`
	// 配置id
	ConfId string `json:"confId"`
}
