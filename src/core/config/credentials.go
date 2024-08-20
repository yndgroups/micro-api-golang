package config

import (
	"core/logger"
	"crypto/tls"
	"crypto/x509"
	"os"

	"google.golang.org/grpc/credentials"
)

// InitClientCredentials 初始化客户端证书
func InitClientCredentials(configPath string) credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair(configPath+"/client.pem", configPath+"/client.key")
	certPool := x509.NewCertPool()
	ca, _ := os.ReadFile(configPath + "/ca.crt")
	certPool.AppendCertsFromPEM(ca)
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "*.npwl.shop",
		RootCAs:      certPool,
	})
}

// InitServerCredentials 初始化服务端证书
func InitServerCredentials(configPath string) credentials.TransportCredentials {
	cert, configErr := tls.LoadX509KeyPair(configPath+"/server.pem", configPath+"/server.key")
	if configErr != nil {
		logger.SugarLogger.Infof("证书加载错误：%v", configErr)
	}
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, caErr := os.ReadFile(configPath + "/ca.crt")
	if caErr != nil {
		logger.SugarLogger.Infof("ca证书读取错误: %v", caErr)
	}
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	return credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ClientAuth: tls.RequireAndVerifyClientCert,
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})
}
