package model

// BaseModel 基础
type BaseModel struct {
	DelStatus  bool      `gorm:"-" json:"delStatus,bool" swaggerignore:"true" minimum:"0" maximum:"1" default:"0"`
	CreateBy   string    `gorm:"->;<-:create" json:"createBy" swaggerignore:"true" minLength:"6" maxLength:"19"`
	UpdateBy   string    `gorm:"->;<-:update" json:"updateBy,omitempty" swaggerignore:"true" minLength:"6" maxLength:"19"`
	CreateTime LocalDate `gorm:"->" json:"createTime,string,omitempty" swaggerignore:"true"`
	UpdateTime LocalDate `gorm:"->" json:"updateTime,string,omitempty" swaggerignore:"true"`
} // @Name BaseModel 基础字段模型

// Ids id集合
type Ids []string

// WechatRequestBack 微信请求实体
type WechatRequestBack struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	SessionKey   string `json:"session_key"`
	ExpiresIn    int64  `json:"expires_in"`
	Token        string `json:"token"`
	Scope        string `json:"scope"`
	OpenId       string `json:"openid"`
	Errcode      int    `json:"errcode"`
	Msg          string `json:"Msg"`
}

// BatchDeleteIds 批量删除id
type BatchDeleteIds struct {
	Ids string `validate:"required" json:"ids"`
}

// ProductParam 商品下单列表参数
type ProductParam struct {
	// 商品id
	ProductId string `validate:"required" json:"productId"`
	// 购买数量
	Quantity uint64 `validate:"required" json:"quantity"`
	// 单规格还是多规格
	SpecType uint `validate:"required" json:"specType"`
	// 邮费模板id
	TempId string `validate:"required" json:"tempId"`
	// 优惠券id
	CouponId string `json:"couponId"`
} // @name ProductParma

// ProductExpand 扩展参数
type ProductExpand struct {
	// 商品id
	ProductId string `json:"productId"`
	// SkuId
	SkuId string `json:"skuId"`
	// 分类id
	CategoryId string `json:"categoryId"`
	// 商品名称
	ProductName string `json:"productName"`
	// 关键字
	Keyword string `json:"keyword"`
	// 商品简介
	Introduction string `json:"introduction"`
	// 商品价格(单规格时必填)
	Price float64 `json:"price"`
	// 成本价(单规格时必填)
	Cost float64 `json:"cost"`
	// 会员价格(单规格时必填)
	VipPrice float64 `json:"vipPrice"`
	// 市场价(单规格时必填)
	MarketPrice float64 `json:"marketPrice"`
	// 商品图片
	Image string `json:"image"`
	// 轮播图
	SliderImage string `json:"sliderImage"`
	// 推荐图
	RecommendImage string `json:"recommendImage"`
	// 库存
	Stock uint64 `json:"stock"`
	// 库存预警
	StockWarning uint64 `json:"stockWarning"`
	// 单位
	UnitName string `json:"unitName"`
	// 运费模板ID
	TempId string `json:"tempId"`
	// 规格(0:单 1:多)
	SpecType int `json:"specType"`
	// 规格名称
	SpecTypeName string `json:"specTypeName"`
	// 上架状态（0：未上架，1：上架）
	IsDisplay int `json:"isDisplay"`
	// 上架状态名称
	IsDisplayName string `json:"isDisplayName"`
	// 商品是否是虚拟商品(0:否 1：是)
	IsVirtual int `json:"isVirtual"`
	// 商品是否是虚拟商品名称
	IsVirtualName string `json:"isVirtualName"`
	// 虚拟商品类型
	VirtualType int `json:"virtualType"`
	// 虚拟销量
	VirtualSales int64 `json:"virtualSales"  swaggerignore:"true"`
	// 是否开启限购(0:否 1:是)
	IsLimit int `json:"isLimit"`
	// 是否开启限购名称
	IsLimitName string `json:"isLimitName"`
	// 限购类型(1:单次限购 2:永久限购)
	LimitType int `json:"limitType"`
	// 限购类型名称
	LimitTypeName string `json:"limitTypeName"`
	// 限购数量
	LimitNum int `json:"limitNum"`
	// 商品详情
	Content string `json:"content"`
}

// MicroServiceHttpParam 微服务http请求参数
type MicroServiceHttpParam[T any] struct {
	ServiceName string
	HttpMethod  string
	ApiPath     string
	Url         string
	Method      string
	AccessToken string
	ReqParam    T
}

// IPConnectionPool ip链接池
type IPConnectionPool map[string][]Instance

// Instance 地址对
type Instance struct {
	Ip      string
	Port    uint64
	Healthy bool
	Weight  float64
}

// InitParam 初始化参数
type InitParam struct {
	NamespaceId string   // 命名空间
	DataId      string   // 配置id
	GroupName   string   // 分组名称
	ServiceName string   // 服务名称
	IpAddr      string   // 服务启动地址
	Port        uint64   // nacos注册端口
	LogDir      string   // 日志路径
	CacheDir    string   // 缓存路径
	LogLevel    string   //日志级别
	RegisterIp  string   // 当前启动服务的Ip
	Clusters    []string // 集群名称
	ServerPort  uint64   // 启动端口
	IsClient    bool     // 是否客户端
	UserName    string
	Password    string
}

// WeUserInfo 微信用户信息
type WeUserInfo struct {
	City           string  `json:"city"`
	Country        string  `json:"country"`
	GroupId        int     `json:"groupid"`
	HeadImgurl     string  `json:"headimgurl"` // "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJVsGvqgby9gHKnggqX6oHIejYu5C8lhAqOXAS0uDopog4To9jwHw4iauVbSg1PZ97RjBbhjzvczUw/132"
	Language       string  `json:"language"`
	NickName       string  `json:"nickname"`
	Openid         string  `json:"openid"` // : "o2Tu45zDgz7c_rtcGz9xBSvUp7JA"
	Province       string  `json:"province"`
	QrScene        int     `json:"qr_scene"`
	QrSceneStr     int     `json:"qr_scene_str"`
	Remark         string  `json:"remark"`
	Sex            int     `json:"sex"`
	Subscribe      int     `json:"subscribe"`
	SubscribeScene string  `json:"subscribe_scene"`
	SubscribeTime  int     `json:"subscribe_time"`
	TagidList      *string `json:"tagid_list"`
}

// SysConfig 系统配置
type SysConfig struct {
	// 配置id
	ConfId string `json:"confId"`
	// 应用名称
	AppName string `json:"appName"`
	// 应用id
	AppId string `json:"appId"`
	// 来源id
	FromAppId int64 `json:"fromAppId"`
	// 密钥
	AppSecret string `json:"appSecret"`
	// 支付商户号
	MchId string `json:"mchId"`
	// 私人密钥
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
	// 商户证书的证书序列号
	SerialNo string `json:"serialNo"`
	// apiV2Key，商户平台获取
	ApiV2Key string `json:"apiV2Key"`
	// apiV3Key，商户平台获取
	ApiV3Key string `json:"apiV3Key"`
	// 私钥
	PrivateKey string `json:"privateKey"`
	// 删除状态
	DelStatus bool `json:"delStatus"`
	// 创建人
	CreateBy string `json:"createBy"`
	// 修改人
	UpdateBy string `json:"updateBy"`
	// 创建时间
	CreateTime string `json:"createTime"`
	// 修改时间
	UpdateTime string `json:"cpdateTime"`
}

// LoginDTO 登录参数
type LoginDTO struct {
	// 用户名 [必填]
	UserName string `json:"userName"`
	// 用户密码 [必填]
	UserPassword string `json:"userPassword"`
	// 电话号码
	Phone string `json:"phone"`
	// 邮箱
	Email string `json:"email"`
} //@Name LoginDTO 登录参数

// JsSdkDTO 获取微信js-sdk的请求参数
type JsSdkDTO struct {
	// 随机地址不能为空
	RandomUrl string `json:"randomUrl" validate:"required"`
	// 配置id
	ConfId string `json:"confId" validate:"required"`
} // @Name JsSdkDTO 获取微信js-sdk的请求参数

// WechatLoginParams 微信登陆参数
type WechatLoginParams struct {
	// 微信wxCode
	WxCode string `json:"wxCode" validate:"required"`
	// 微信昵称
	NickName string `json:"nickName"`
	// 姓名
	TrueName string `json:"trueName"`
	// 头像[必填]
	AvatarUrl string `json:"avatarUrl"`
	// 性别（1：男， 2：女）")
	Gender uint32 `json:"gender"`
	// 用户类型（1：教师，2:学生，3：家长，4：其他社会人员）
	UserType string `json:"userType"`
	// 所在省
	Province string `json:"province"`
	// 所在市
	City string `json:"city"`
}

// WechatLoginDTO 微信实体
type WechatLoginDTO struct {
	// 微信wxCode
	WxCode string `json:"wxCode" validate:"required"`
	// 微信昵称
	NickName string `json:"nickName"`
	// 姓名
	TrueName string `json:"trueName"`
	// 头像[必填]
	AvatarUrl string `json:"avatarUrl"`
	// 性别（1：男， 2：女）")
	Gender int64 `json:"gender"`
	// 用户类型（1：教师，2:学生，3：家长，4：其他社会人员）
	UserType string `json:"userType"`
	// 所在省
	Province string `json:"province"`
	// 所在市
	City string `json:"city"`
	// 注册来源(1:后台端 2:前端 3: 微信公众号 4:微信小程序）
	RegType int64 `json:"regType"`
	// 配置的id
	AppId string `json:"appId"`
}

// WeChatWebLoginDTO 微信公众号登录参数
type WeChatWebLoginDTO struct {
	// 配置的id
	ConfId string `json:"confId" validate:"required"`
	// 微信wxCode
	WxCode string `json:"wxCode" validate:"required"`
	// 登录类型
	Type int `json:"type" validate:"required"`
} // @Name WeChatWebLoginDTO 微信公众号登录参数

// WeChatLoginDTO 微信登录参数
type WeChatLoginDTO struct {
	// 配置的id [必填]
	ConfId string `json:"confId" validate:"required"`
	// 微信code [必填]
	WxCode string `json:"wxCode" validate:"required"`
	// 用户id 主要用于同步其他应用 不让前端传参
	UserId string `json:"userId" swaggerignore:"true"`
	// 微信昵称
	NickName string `json:"nickName"`
	// 姓名
	TrueName string `json:"trueName"`
	// 头像[必填]
	AvatarUrl string `json:"avatarUrl" validate:"required"`
	// 性别（1：男， 2：女）")
	Gender string `json:"gender"`
	// 用户类型（1：教师，2:学生，3：家长，4：其他社会人员）
	UserType string `json:"userType"`
	// 所在省
	Province string `json:"province"`
	// 所在市
	City string `json:"city"`
}

// AccessToken 微信AccessToken信息
// "{\"access_token\":\"53_qjyVwyO-aAbNo0ZjT1PUynSnWdGHbaKK8kjJVGg17aiCUMMC5X7XvRT5vRuJUuT6asgMfgQwaBWbOz_zj6gs5vavWvzY14ZICqV-YhAr2oF6EGiC9mwsybGBLfj4ssfqrMbFwl99A-n7unykNZZjAJAQCC\",\"expires_in\":7200}"
// {"errcode":40164,"Msg":"invalid ip 125.33.201.21 ipv6 ::ffff:125.33.201.21, not in whitelist rid: 6207df1a-08ef8516-2a6c18c6"}
type AccessToken struct {
	// AccessToken值
	AccessToken string `json:"access_token"`
	// 过期时间
	ExpiresIn int `json:"expires_in"`
	// 错误编码
	ErrCode int64 `json:"errcode"`
	// 错误信息
	Msg string `json:"Msg"`
}

type WeChatUser struct {
	Subscribe      int           `json:"subscribe"`
	Openid         string        `json:"openid"`
	Nickname       string        `json:"nickname"`
	Sex            uint32        `json:"sex"`
	Language       string        `json:"language"`
	City           string        `json:"city"`
	Province       string        `json:"province"`
	Country        string        `json:"country"`
	Headimgurl     string        `json:"headimgurl"`
	SubscribeTime  int           `json:"subscribe_time"`
	Remark         string        `json:"remark"`
	Groupid        int           `json:"groupid"`
	TagidList      []interface{} `json:"tagid_list"`
	SubscribeScene string        `json:"subscribe_scene"`
	QrScene        int           `json:"qr_scene"`
	QrSceneStr     string        `json:"qr_scene_str"`
}

// LoginUser 用户登录参数
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
}

// File 文件相关
type File struct {
	// 文件路径
	FilePath string `json:"filePath" validate:"required"`
	// 尺寸类型
	SizeType int `json:"sizeType" validate:"required"`
}
