package pool

import (
	"core/config"
	"core/logger"
	"core/utils"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GrpClient struct {
	Ip           string           // 连接池ip
	GrpcClient   *grpc.ClientConn // 连接池实列
	StrategyName string           // 策略名称
	Weight       uint             // 权重值
}

// 连接池信息数据
type poolMaps map[string][]*GrpClient

// 客户端连接池
var GrpcClientPool = make(poolMaps, 0)

// 策略执行标记
var strategyMark = make(map[string]int, 0)

/*
 * 从链接池获取
 */
func (maps poolMaps) Get(serverName string) (*grpc.ClientConn, error) {
	if serverName == "" {
		return nil, errors.New("服务名称不能为空")
	}
	list := maps[serverName]
	logger.SugarLogger.Infof("负载均衡策略为%s,连接池数:%d, 连接池为: %v", config.Global.ServerConfig().LoadBalance, len(list), list)
	if list != nil {
		index := 0
		switch config.Global.ServerConfig().LoadBalance {
		case "polling": // 轮询
			index = maps.getPollingMark(serverName)
			logger.SugarLogger.Infof("当前服务数量为 [%v] 轮询数为：%v", len(list), index)
		case "random": // 随机
			index = utils.GetRandom(0, len(list))
			logger.SugarLogger.Infof("当前服务数量为 [%v] 随机数为：%v", len(list), index)
		case "weight":
			logger.SugarLogger.Info("权重")
		case "hash":
			logger.SugarLogger.Info("哈希")
		default:
			logger.SugarLogger.Info("权重")

		}
		return list[index].GrpcClient, nil
	}
	return nil, errors.New("连接池不存在初始化的数据或服务提供者不存在")
}

/*
 * 连接池初始化及管理
 */
func (maps poolMaps) Put(serverName string, providerIpPorts []string, appCrt credentials.TransportCredentials) error {
	if serverName == "" {
		return errors.New("服务名称不能为空")
	}
	if len(providerIpPorts) == 0 {
		return errors.New("需要创建的服务及端口不能为空")
	}
	println(fmt.Sprintf("providerIpPorts = %v", len(providerIpPorts)))
	// 关闭原来的链接重新创建新的链接保持一致性
	if len(maps[serverName]) > 0 {
		for _, v := range maps[serverName] {
			v.GrpcClient.Close()
		}
	}
	// 创建新的客户端连接池
	pools := make([]*GrpClient, 0)
	for _, ipPort := range providerIpPorts {
		if grpcClient, connectErr := grpc.Dial(
			ipPort,
			grpc.WithTransportCredentials(appCrt),
		); connectErr != nil {
			return errors.New("初始化grpc失败")
		} else {
			pools = append(pools, &GrpClient{
				Ip:         ipPort,
				GrpcClient: grpcClient,
			})
		}
	}
	maps[serverName] = pools
	println(fmt.Sprintf("当前服务 [%v] 初始化了 [%v] 个grpc客户端\n", serverName, len(maps[serverName])))
	if len(maps[serverName]) > 0 {
		for k, v := range maps[serverName] {
			println(fmt.Sprintf("第[%v]个: %v \n", k+1, v))
		}
	}
	return nil
}

// 删除连接池
func (pool poolMaps) Delete(serverName string) error {
	if serverName == "" {
		return errors.New("服务名称不能为空")
	}
	delete(pool, serverName)
	return nil
}

/*
 * 轮询标记
 */
func (pool poolMaps) getPollingMark(serverName string) int {
	if strategyMark[serverName] >= len(pool[serverName])-1 {
		strategyMark[serverName] = 0
	} else {
		strategyMark[serverName]++
	}
	return strategyMark[serverName]
}
