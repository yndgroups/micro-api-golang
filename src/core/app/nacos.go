package app

import (
	"core/config"
	"core/db"
	"core/logger"
	coreModel "core/model"
	"core/pool"
	"core/redis"
	"errors"
	"fmt"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Nacos 初始化配置中心
var Nacos = &nacos{}

// nacos 配置中心
type nacos struct {
	// 客户端
	namingClient naming_client.INamingClient
	// 客户端
	cfgClient config_client.IConfigClient
	// 初始化参数
	initParam *coreModel.InitParam
	// 配置参数
	cfg *config.Configs
	// 应用证书
	appCrt credentials.TransportCredentials
}

// clientPrefix 客户端前缀
var clientPrefix = "consumer"

// serverPrefix 服务端前缀
var serverPrefix = "provider"

// InitNacos 初始化nacos
func (n *nacos) InitNacos(param coreModel.InitParam, credentials credentials.TransportCredentials) (*nacos, error) {
	// logger.SugarLogger.Infof("init nacos config = %v", param)
	n.initParam = &param
	n.appCrt = credentials
	cc := constant.ClientConfig{
		NamespaceId:          param.NamespaceId, // 命名空间id
		TimeoutMs:            5 * 1000,          // http请求超时时间，单位毫秒
		BeatInterval:         5 * 1000,          //心跳间隔时间，单位毫秒（仅在ServiceClient中有效）
		UpdateThreadNum:      20,                //更新服务的线程数
		LogDir:               param.LogDir,
		CacheDir:             param.CacheDir,
		LogLevel:             param.LogLevel,
		NotLoadCacheAtStart:  true, //在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: true, //当服务列表为空时是否更新本地缓存，true--更新,false--不更新
		Username:             param.UserName,
		Password:             param.Password,
	}
	sc := []constant.ServerConfig{
		{
			IpAddr: param.IpAddr,
			Port:   param.Port,
		},
		// *constant.NewServerConfig(param.IpAddr, param.Port),
	}
	logger.SugarLogger.Infof(">>>>>> 初始化服务:[%s]到组[%s]", param.ServiceName, param.GroupName)
	if client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	); err != nil {
		logger.SugarLogger.Errorf("nacos初始化失败,错误信息：%v", err.Error())
		return nil, errors.New("nacos初始化失败")
	} else {
		n.namingClient = client
	}

	// 创建客户端连接
	namingClient, namingClientErr := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if namingClientErr != nil {
		logger.SugarLogger.Errorf("nacos的client初始化失:%v", namingClientErr.Error())
		return nil, errors.New("nacos的client初始化失")
	}
	// 获取nacos远程配置
	if content, configErr := namingClient.GetConfig(vo.ConfigParam{
		DataId: n.initParam.DataId,
		Group:  n.initParam.GroupName,
	}); configErr != nil {
		logger.SugarLogger.Errorf("nacos配置文件获取失败: %v", configErr.Error())
		return nil, errors.New("nacos配置文件获取失败")
	} else {
		if content == "" {
			logger.SugarLogger.Error("nacos获取的配置文件内容为空")
			return nil, errors.New("nacos获取的配置文件内容为空")
		}
		n.cfgClient = namingClient
		n.cfg = config.Global.Init(content)       // 初始化配置
		n.registerInstance(n.namingClient, param) // 注册服务
		n.initApp(n.cfg)                          // 初始化应用
		n.listenConfig(namingClient)
		if n.initParam.IsClient {
			n.subscribe(param.ServiceName) // 监听服务
		}
		return n, nil
	}
}

// getNamingClient 获取NamingClient实例
func (n *nacos) getNamingClient() naming_client.INamingClient {
	return n.namingClient
}

// getGrpcPoolByServerName 获取服务实例ip信息
func (n *nacos) getGrpcPoolByServerName(serverName string, tlsCredentials credentials.TransportCredentials) (*grpc.ClientConn, error) {
	if serverName == "" {
		return nil, errors.New("服务名称不能为空")
	}
	if grpcClient, err := pool.GrpcClientPool.Get(serverName); err != nil {
		return nil, err
	} else {
		return grpcClient, nil
	}

}

// registerInstance 注册实例
func (n *nacos) registerInstance(client naming_client.INamingClient, param coreModel.InitParam) {
	// logger.SugarLogger.Infof(">>> 服务: [ %s ] 注册到组: [ %s ]", param.ServiceName, param.GroupName)
	if bool, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          param.RegisterIp,
		Port:        param.ServerPort,
		ServiceName: n.initParam.ServiceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "DEFAULT",
		GroupName:   n.initParam.GroupName,
	}); bool {
		logger.SugarLogger.Infof(">>>>>> 服务[ %s ]注册nacos实例成功", param.ServiceName)
		n.subscribe(n.initParam.ServiceName)
	} else {
		logger.SugarLogger.Error(">>>>>> 注册实例到nacos失败,失败原因:%s", err.Error())
	}
}

// getURLByServiceInfo 根据服务名称获取完整的请求地址
// serverName  服务名称
// apiPath  api地址
// author 杨大琼
// date 2022-02-20 17:51
func (n *nacos) getURLByServiceInfo(serverName string, httpMethod string, apiPath string) (string, error) {
	services, err := n.getNamingClient().GetService(vo.GetServiceParam{
		ServiceName: serverName,
		GroupName:   n.initParam.GroupName,
		//Clusters:    []string{"DEFAULT"},
	})
	if err != nil {
		return "", errors.New("获取服务实例失败")
	}
	hosts := services.Hosts
	hostInfo := ""
	for _, v := range hosts {
		if v.Healthy {
			hostInfo = fmt.Sprintf("%v:%v", v.Ip, v.Port)
			break
		}
	}
	if len(hostInfo) > 0 {
		return fmt.Sprintf(`%v%v%v`, httpMethod, hostInfo, apiPath), nil
	} else {
		// 当前商品服务没有可调用的健康实例
		return "", errors.New(" 当前商品服务没有可调用的健康实例")
	}
}

// selectInstances 获取实例列表
// 只返回满足这些条件的实例列表：healthy=${HealthyOnly},enable=true 和 weight > 0
func (n *nacos) selectInstances(serviceName string) error {
	if serviceName == "" {
		return errors.New("服务，名称不能为空")
	}

	// 服务提供者实例
	if strings.Contains(serviceName, serverPrefix) {
		logger.SugarLogger.Infof(">>>>>> selectInstances: 获取 [%v, %v] 服务实例", serviceName, n.initParam.GroupName)
		if providerInstances, selectErr := n.namingClient.SelectInstances(vo.SelectInstancesParam{
			ServiceName: serviceName,
			GroupName:   n.initParam.GroupName, // 默认值DEFAULT_GROUP
			//Clusters:    []string{"DEFAULT"}, // 默认值DEFAULT
			HealthyOnly: true,
		}); selectErr != nil {
			return selectErr
		} else {
			logger.SugarLogger.Infof(">>>>>> 服务实例", providerInstances)
			println(len(providerInstances))
			if len(providerInstances) == 0 {
				return errors.New("该服务" + serviceName + "未成功对应的提供者信息,请启动服务提供者或者检查服务提供者是否正常启动")
			}
			providerIpPorts := make([]string, 0)
			for _, v := range providerInstances {
				if v.Healthy {
					providerIpPorts = append(providerIpPorts, fmt.Sprintf("%v:%v", v.Ip, v.Port))
				}
			}
			logger.SugarLogger.Infof(">>>>>> 初始化实例到 [ %v ] 服务的连接池: %+v", serviceName, providerIpPorts)
			pool.GrpcClientPool.Put(serviceName, providerIpPorts, n.appCrt)
			n.subscribe(serviceName)
			return nil
		}

	}

	// 服务消费者实例
	if strings.Contains(serviceName, clientPrefix) {
		logger.SugarLogger.Infof(">>>>>> selectInstances: 获取 [%v] 服务实例", serviceName)
		consumerInstances, err := n.namingClient.SelectInstances(vo.SelectInstancesParam{
			ServiceName: serviceName,
			GroupName:   n.initParam.GroupName, // 默认值DEFAULT_GROUP
			//Clusters:    []string{"DEFAULT"}, // 默认值DEFAULT
			HealthyOnly: true,
		})
		if len(consumerInstances) > 0 {
			ipPorts := make([]string, 0)
			for _, v := range consumerInstances {
				if v.Healthy {
					ipPorts = append(ipPorts, fmt.Sprintf("%v:%v", v.Ip, v.Port))
				}
			}
			logger.SugarLogger.Infof("serviceName = %v, len = %v,instances = %v", serviceName, len(consumerInstances), consumerInstances)
			if len(ipPorts) > 0 {
				return nil
			} else {
				return errors.New("没有健康的实例可调用")
			}
		} else {
			pool.GrpcClientPool.Delete(serviceName)
			return err
		}
	}
	// 不符合约定的服务名称
	return errors.New("非法的服务名称")
}

// subscribe 监听服务变化
func (n *nacos) subscribe(serviceName string) {
	// println(fmt.Printf(" %+v", param))
	logger.SugarLogger.Infof(">>>> 开启服务监听 subscribe:" + serviceName)
	// Subscribe key=serviceName+groupName+cluster
	// 注意:我们可以在相同的key添加多个SubscribeCallback.
	if err := n.namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName: serviceName,
		GroupName:   n.initParam.GroupName, // 默认值DEFAULT_GROUP
		//Clusters:    []string{"DEFAULT"}, // 默认值DEFAULT
		SubscribeCallback: func(services []model.Instance, err error) {
			logger.SugarLogger.Infof(">>>> 服务监听数据更新:" + serviceName)
			if err != nil {
				logger.SugarLogger.Warnf("服务监听出错,原因为:%v", err.Error())
				if n.initParam.IsClient {
					sName := n.initParam.ServiceName
					pool.GrpcClientPool.Delete(strings.Replace(sName, clientPrefix, serverPrefix, 1))
				}
				return
			}
			if n.initParam.IsClient {
				sName := strings.Replace(serviceName, clientPrefix, serverPrefix, 1)
				logger.SugarLogger.Infof(">>>> 从新获取[%v]服务监听数据更新", sName)
				n.selectInstances(sName)
			}
		},
	}); err != nil {
		logger.SugarLogger.Warnf("服务监听失败：%v", err.Error())
	}
}

// listenConfig 监听配置变化
func (n *nacos) listenConfig(client config_client.IConfigClient) {
	logger.SugarLogger.Infof("开启开启配置监听...... ListenConfig")
	if err := client.ListenConfig(vo.ConfigParam{
		DataId: n.initParam.DataId,
		Group:  n.initParam.GroupName,
		OnChange: func(namespace, group, dataId, data string) {
			logger.SugarLogger.Infof("远程配置文件已经更新group: %v,dataId: %v,config: %v namespace: %v", group, dataId, data, namespace)
			n.cfg = config.Global.Init(data)
			n.initApp(n.cfg)
		},
	}); err != nil {
		logger.SugarLogger.Infof("配置监听失败：%v", err.Error())
	}
}

// selectOneHealthyInstance 获取一个健康的实例（加权随机轮询）
// 将会按加权随机轮询的负载均衡策略返回一个健康的实例
// 实例必须满足的条件：health=true,enable=true and weight>0
func (n *nacos) selectOneHealthyInstance() {
	instance, err := n.namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: n.initParam.ServiceName,
		GroupName:   n.initParam.GroupName, // 默认值DEFAULT_GROUP
		Clusters:    []string{"DEFAULT"},   // 默认值DEFAULT
	})
	if err != nil {
		logger.SugarLogger.Warnf("获取一个健康的例列表失败,原因:%S", err.Error())
	}
	println(fmt.Sprintf("%v", instance))
}

// 注销实例
func (n *nacos) deregisterInstance() {
	success, err := n.namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          n.initParam.IpAddr,
		Port:        n.initParam.Port,
		ServiceName: n.initParam.ServiceName,
		Ephemeral:   true,
		GroupName:   n.initParam.GroupName, // 默认值DEFAULT_GROUP
		//Cluster:     "DEFAULT",           // 默认值DEFAULT
	})
	println(success, err)
}

// 获取服务信息
func (n *nacos) getService() {
	_, err := n.namingClient.GetService(vo.GetServiceParam{
		ServiceName: n.initParam.ServiceName,
		GroupName:   n.initParam.GroupName, // 默认值DEFAULT_GROUP
		//Clusters:    []string{"默认值DEFAULT"}, // 默认值DEFAULT
	})
	println(err)
}

/*
 * 获取所有的实例列表
 * 可以返回全部实例列表,包括healthy:false,enable:false,weight<=0
 */
func (n *nacos) selectAllInstances(param coreModel.InitParam) []model.Instance {
	sName := param.ServiceName
	if n.initParam.IsClient {
		sName = strings.Replace(sName, clientPrefix, serverPrefix, 1)
	}
	if instances, err := n.namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: param.ServiceName,
		GroupName:   param.GroupName, // 默认值DEFAULT_GROUP
		//Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
	}); err != nil {
		logger.SugarLogger.Warnf("获取服务[%v]所有的实例列表失败,原因:%v", sName, err.Error())
		return nil
	} else {
		println(fmt.Sprintf("%v", instances))
		return instances
	}
}

// 取消服务监听
func (n *nacos) unsubscribe() {
	if err := n.namingClient.Unsubscribe(&vo.SubscribeParam{
		ServiceName: n.initParam.ServiceName,
		GroupName:   n.initParam.GroupName, // 默认值DEFAULT_GROUP
		Clusters:    []string{"DEFAULT"},   // 默认值DEFAULT
		SubscribeCallback: func(services []model.Instance, err error) {
			//log.Printf("\n\n callback return services:%s \n\n", utils.ToJsonString(services))
		},
	}); err != nil {
		logger.SugarLogger.Warnf("取消服务监听失败,原因:%S", err.Error())
	}
}

// 获取服务名列表
func (n *nacos) getAllServicesInfo() {
	_, err := n.namingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		NameSpace: n.initParam.NamespaceId,
		PageNo:    1,
		PageSize:  10,
	})
	println(err)
}

// 发布配置
func (n *nacos) publishConfig() {
	success, err := n.cfgClient.PublishConfig(vo.ConfigParam{
		DataId:  n.initParam.DataId,
		Group:   n.initParam.GroupName,
		Content: "test config"})
	println(success, err)
}

// 删除配置
func (n *nacos) deleteConfig() {
	success, err := n.cfgClient.DeleteConfig(vo.ConfigParam{
		DataId: n.initParam.DataId,
		Group:  n.initParam.GroupName})
	println(success, err)
}

// 获取配置
func (n *nacos) getConfig() {
	content, err := n.cfgClient.GetConfig(vo.ConfigParam{
		DataId: n.initParam.DataId,
		Group:  n.initParam.GroupName})
	println(content, err)
}

// 取消配置监听
func (n *nacos) cancelListenConfig() {
	err := n.cfgClient.CancelListenConfig(vo.ConfigParam{
		DataId: n.initParam.DataId,
		Group:  n.initParam.GroupName,
	})
	println(err)
}

// 搜索配置
func (n *nacos) searchConfig() {
	err := n.cfgClient.CancelListenConfig(vo.ConfigParam{
		DataId: n.initParam.DataId,
		Group:  n.initParam.GroupName,
	})
	println(err)
}

// initApp 初始化数应用所需的环境
func (n *nacos) initApp(cfg *config.Configs) {
	// 初始化数据库
	if cfg.Datasource.Enabled {
		db.InitDb(cfg.Datasource)
	}
	// 初始化Redis
	if cfg.RedisConfig.Enabled {
		redis.InitConnectRedis(cfg.RedisConfig)
	}
	// 如果配置初始化钉钉
	if cfg.Dingtalk.Enabled {
		InitDingtalk(cfg.Dingtalk)
	}
}

// 公共服务
func (n *nacos) CommonGrpcProvider() *grpc.ClientConn {
	if con, conErr := n.getGrpcConByServerName(config.Global.ServerConfig().CommonProvider); conErr != nil {
		// "生产者服务获取异常！"
		logger.SugarLogger.Errorf(">>>>>> 生产者 [ %s ] 服务获取异常！", config.Global.ServerConfig().CommonProvider)
		panic(conErr.Error())
	} else {
		return con
	}
}

// 商城后台服务提供者
func (n *nacos) ShoppingGrpcAdminProvider() *grpc.ClientConn {
	if con, conErr := n.getGrpcConByServerName(config.Global.ServerConfig().ShoppingAdminProvider); conErr != nil {
		// "生产者服务获取异常！"
		logger.SugarLogger.Errorf(">>>>>> 生产者 [ %s ] 服务获取异常！", config.Global.ServerConfig().ShoppingAdminProvider)
		panic(conErr.Error())
	} else {
		return con
	}
}

// 商城客户端服务提供者
func (n *nacos) ShoppingClientProvider() *grpc.ClientConn {
	if con, conErr := n.getGrpcConByServerName(config.Global.ServerConfig().ShoppingClientProvider); conErr != nil {
		// "生产者服务获取异常！"
		logger.SugarLogger.Errorf(">>>>>> 生产者 [ %s ] 服务获取异常！", config.Global.ServerConfig().ShoppingClientProvider)
		panic(conErr.Error())
	} else {
		return con
	}
}

// OrderGrpcClient
func (n *nacos) OrderGrpcProvider() *grpc.ClientConn {
	if con, conErr := n.getGrpcConByServerName(config.Global.ServerConfig().OrderProvider); conErr != nil {
		// "生产者服务获取异常！"
		logger.SugarLogger.Errorf(">>>>>> 生产者 [ %s ] 服务获取异常！", config.Global.ServerConfig().OrderProvider)
		panic(conErr.Error())
	} else {
		return con
	}
}

/*
 *根据服务名称获取对应的服务
 */
func (n *nacos) getGrpcConByServerName(serverName string) (*grpc.ClientConn, error) {
	logger.SugarLogger.Infof(">>>>>> getGrpcConByServerName 获取[ %v ] 的服务客户端链接", serverName)
	if grpcClient, err := pool.GrpcClientPool.Get(serverName); err != nil {
		if err := n.selectInstances(serverName); err != nil {
			logger.SugarLogger.Errorf(">>>>>> 服务[ %v ]实例查询失败", serverName)
			return nil, err
		} else {
			logger.SugarLogger.Infof(">>>>>> 递归获取服务初始化的客户端")
			return n.getGrpcConByServerName(serverName)
		}
	} else {
		return grpcClient, nil
	}
}
