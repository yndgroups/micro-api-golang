package config

import (
	"encoding/json"
	"fmt"
	"time"
)

// Server 服务配置
type Server struct {
	// 可以登陆的唯一权限标识
	AppRoleUniqueIdentification string `json:"appRoleUniqueIdentification"`
	// token 秘钥
	Secret string `json:"secret"`
	// 运维管理员
	Administrator string `json:"administrator"`
	// 上传地址
	UploadPath string `json:"uploadPath"`
	// 登录存储前缀
	LoginPrefix string `json:"loginPrefix"`
	// 配置文件前缀
	ConfigPrefix string `json:"configPrefix"`
	// token前缀
	TokenPrefix string `json:"tokenPrefix"`
	// 权限前缀
	PowerPrefix string `json:"powerPrefix"`
	// 公共服务
	CommonProvider string `json:"common"`
	// 公共服务
	YsjzProvider string `json:"ysjz"`
	// 订单服务
	OrderProvider string `json:"order"`
	// 商城后台服务
	ShoppingAdminProvider string `json:"shoppingAdmin"`
	// 商城后前端服
	ShoppingClientProvider string `json:"shoppingClient"`
	// 资源服务
	Resources string `json:"resources"`
	// 疫情上报服务
	EpidemicMonitoring string `json:"epidemicMonitoring"`
	// 负载均衡
	LoadBalance string `json:"loadBalance" default:"polling"`
	// 小程序配置id
	AppletConfId string `json:"appletConfId"`
	// 公众号配置id
	WechatConfId string `json:"wechatConfId"`
}

// ServerConfig 服务配置
type ServerConfig struct {
	//nacos服务的scheme
	Scheme string `json:"scheme"`
	//nacos服务的 contextpath
	ContextPath string `json:"contextPath"`
	//nacos服务的 address
	IpAddr string `json:"ipAddr"`
	//nacos服务的 port
	Port uint64 `json:"port"`
}

// ClientConfig nacos客户端配置
type ClientConfig struct {
	// 请求Nacos服务器超时，默认值为10000ms
	TimeoutMs uint64 `json:"timeoutMs"`
	// 废弃不建议使用
	ListenInterval uint64 `json:"listenInterval"`
	// 向服务器发送节拍的时间间隔，默认值为5000ms
	BeatInterval int64 `json:"beatInterval"`
	// Nacos的名称空间ID。当命名空间是公共的时，在这里填入空白字符串。
	NamespaceId string `json:"namespaceId"`
	// 应用名称
	AppName string `json:"appAame"`
	// 服务端地址
	Endpoint string `json:"endpoint"`
	// RegionId
	RegionId string `json:"regionId"`
	// 授权秘钥
	AccessKey string `json:"accessKey"`
	// 秘钥
	SecretKey string `json:"secretKey"`
	// 这是打开kms，默认值为false. https://help.aliyun.com/product/28933.html
	OpenKMS bool `json:"openKms"`
	// 持久化nacos服务信息的目录，默认值为当前路径
	CacheDir string `json:"cacheDir"`
	// 更新nacos服务信息的gorutine数，默认值为20
	UpdateThreadNum int `json:"updateThreadNum"`
	// 不在开始时在CacheDir中加载持久nacos服务信息
	NotLoadCacheAtStart bool `json:"notLoadCacheAtStart"`
	// 从服务器获取空服务实例时更新缓存
	UpdateCacheWhenEmpty bool `json:"updateCacheWhenEmpty"`
	// 用户名
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
	// 工作目录
	Dir string `json:"dir"`
	// 轮询时间 eg: 30m, 1h, 24h, default is 24h
	RotateTime string `json:"rotateTime"`
	// 设置资源最大值
	MaxAge int64 `json:"maxAge"`
	// 级别: 必须是调试、信息、警告、错误，默认值为信息
	Level string `json:"level"`
	// sampling 配置
	Sampling *ClientSamplingConfig `json:"sampling"`
	// nacos服务上下文
	ContextPath string `json:"contextPath"`
}

// ClientSamplingConfig sampling 配置
type ClientSamplingConfig struct {
	// 初始值
	Initial int `json:"initial"`
	// sampling 健康状态
	Thereafter int `json:"thereafter"`
	// 备忘
	Tick time.Duration `json:"tick"`
}

// Datasource 数据连接配置
type Datasource struct {
	// 是否启用配置
	Enabled bool `json:"enabled"`
	// 驱动名称
	DriverName string `json:"driverName"`
	// 主机地址
	Host string `json:"host"`
	// 端口
	Port uint64 `json:"port"`
	// 数据库名
	Database string `json:"database"`
	// 用户名
	UserName string `json:"userName"`
	// 密码
	Password string `json:"password"`
	// 字符集
	Charset string `json:"charset"`
}

// RedisConfig Redis数据连接配置
type RedisConfig struct {
	// 是否启用配置
	Enabled bool `json:"enabled"`
	// 主机地址
	Host string `json:"host"`
	// 端口
	Port string `json:"port"`
	// 密码
	Password string `json:"password"`
	// 连接数据库编号
	Database int `json:"database"`
	// 连接池
	Pool RedisConfigPool `json:"pool"`
	// 存储过期时间 // 默认时间2h
	Expire time.Duration `json:"expire" default:"7200"`
	// 微信存储过期时间 // 默认时间7 * 24h
	WeChatExpire time.Duration `json:"weChatExpire" default:"604800"`
}

// SmsTemplate 短息模版
type SmsTemplate struct {
	// 短信模板编码
	TemplateCode *string `json:"templateCode"`
	// 签名名称
	SignName *string `json:"signName"`
}

// Sms 短信发送配置
type Sms struct {
	// 您的AccessKey ID
	AccessKeyId *string `json:"accessKeyId"`
	// 您的AccessKey Secret
	AccessKeySecret *string `json:"accessKeySecret"`
	// 有效期
	Expire time.Duration `json:"expire"`
	// 设置访问域名
	Host *string `json:"host"`
	// 配置
	Templates map[string]SmsTemplate `json:"templates"`
}

// RedisConfigPool Redis数据连接池
type RedisConfigPool struct {
	// 连接池的最大数据库连接数。设为0表示无限制。取值为20，表示同时最多有20个数据库连
	MaxActive int `json:"maxActive"`
	// 最大空闲数，数据库连接的最大空闲时间。超过空闲时间，数据库连接将被标记为不可用，然后被释放。设为0表示无限制。
	MaxIdle int `json:"maxIdle"`
	// 最大等待毫秒数, 单位为 ms, 超过时间会出错误信息
	MaxWait int64 `json:"maxWait"`
	// 超时时间
	Timeout int64 `json:"timeout"`
	// 等待时间
	Wait bool `json:"wait"`
}

// Dingtalk 钉钉配置
type Dingtalk struct {
	// 是否启用配置
	Enabled bool `json:"enabled"`
	// token
	Token string `json:"token"`
	// 秘钥
	Secret string `json:"secret"`
}

// Configs 总配置
type Configs struct {
	// 数据库配置
	Datasource Datasource `json:"datasource"`
	// redis配置
	RedisConfig RedisConfig `json:"redisConfig"`
	// 服务配置
	Server Server `json:"server"`
	// 短信配置
	Sms Sms `json:"sms"`
	// 钉钉配置
	Dingtalk Dingtalk `json:"dingtalk"`
}

type config struct{}

// Global 全局对象
var Global = config{}

// cfgs 初始化对象存储
var cfgs *Configs

// Init 初始化全局参数
func (g *config) Init(configStr string) *Configs {
	initConfigs := &Configs{}
	err := json.Unmarshal([]byte(configStr), &initConfigs)
	if err != nil {
		fmt.Println("配置解析失败...")
		panic(err)
	}
	cfgs = initConfigs
	return cfgs
}

// All 全部服务配置
func (g *config) All() *Configs {
	return cfgs
}

// ServerConfig 服务配置
func (g *config) ServerConfig() Server {
	return cfgs.Server
}

// DatasourceConfig 服务配置
func (g *config) DatasourceConfig() Datasource {
	return cfgs.Datasource
}

// RedisConfig 服务配置
func (g *config) RedisConfig() RedisConfig {
	return cfgs.RedisConfig
}

// SmsConfig 服务配置
func (g *config) SmsConfig() Sms {
	return cfgs.Sms
}
