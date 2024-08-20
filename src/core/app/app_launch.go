package app

import (
	"context"
	"core/config"
	"core/logger"
	coreModel "core/model"
	"core/utils"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type AppLaunch struct {
	Active           string                           // 当前环境
	Port             uint64                           //运行端口
	ConfigPath       string                           // 配置文件路径
	Ip               string                           // 服务ip地址
	Credentials      credentials.TransportCredentials // 证书
	GrpcServer       *grpc.Server                     // grpc服务
	Server           *http.Server                     // http服务
	shutdownFinished chan struct{}
}

var ENV string

// 初始化环境变量
func (app *AppLaunch) initEnv() *AppLaunch {
	workDir, err := os.Getwd()
	if err != nil {
		panic("操作系统导致启动失败")
	}
	active := flag.String("active", "dev", "请传入端口参数[dev|test|prod],默认dev")
	conf := flag.String("conf", workDir, "请传入启动配置文件参数")
	port := flag.Uint64("port", 8081, "请传入启动端口号")
	flag.Parse()
	app.Active = *active
	app.Port = *port
	viper.SetConfigName("application." + app.Active)
	viper.SetConfigType("yml")
	confPath := *conf + "/conf"
	app.ConfigPath = confPath
	viper.AddConfigPath(confPath)
	if err := viper.ReadInConfig(); err != nil {
		panic("配置加载失败")
	}
	ENV = *active
	return app
}

// StartHttpServer 启动http服务
func (app *AppLaunch) StartHttpServer(route *gin.Engine, baseApiPath string) *AppLaunch {
	app.initEnv()
	// 开启日志
	logger.InitLog(viper.GetString("zap.log-dir"), "zap.log-level")
	// 客户端配置 如果需要支持多namespace，我们可以场景多个client,
	if app.Active == "dev" {
		localIp, _ := utils.GetTrueIp()
		app.Ip = cast.ToString(localIp)
		// port, _ := utils.GetAvailablePort()
		// app.Port = port
		app.Port = viper.GetUint64("application.port")
	} else {
		app.Ip = utils.GetIpV4()
	}
	logger.SugarLogger.Infof("启动服务....., 启动环境:[%v], 配置路径:[%v], 运行端口: [%v]", app.Active, app.ConfigPath, app.Port)
	initParam := coreModel.InitParam{
		NamespaceId: viper.GetString("nacos.config.namespace"),
		DataId:      viper.GetString("nacos.config.data-id"),
		GroupName:   viper.GetString("nacos.config.group"),
		ServiceName: viper.GetString("application.name"),
		IpAddr:      viper.GetString("nacos.config.server-addr"),
		Port:        viper.GetUint64("nacos.config.port"),
		LogDir:      viper.GetString("nacos.log-dir"),
		CacheDir:    viper.GetString("nacos.cache-dir"),
		LogLevel:    viper.GetString("nacos.log-level"),
		UserName:    viper.GetString("nacos.username"),
		Password:    viper.GetString("nacos.password"),
		RegisterIp:  app.Ip,
		ServerPort:  app.Port,
		IsClient:    true,
	}

	// 初始化证书调用
	app.Credentials = config.InitClientCredentials(app.ConfigPath)

	// 初始化nacos
	if _, err := Nacos.InitNacos(initParam, app.Credentials); err != nil {
		panic(err.Error())
	}

	srv := &AppLaunch{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", app.Port),
			Handler: route,
		},
	}

	go srv.watchServerExit(5 * time.Second)

	// API 文档
	utils.PrintLogo(true, fmt.Sprintf("http://%s:%d%s/swagger/index.html", app.Ip, app.Port, baseApiPath))
	err := srv.listenAndServe()
	if err != nil {
		logger.SugarLogger.Fatalf("来自ListenAndServe的意外错误: %w", err)
	}
	return app
}

// StartGrpcServer 启动Grpc服务
func (app *AppLaunch) StartGrpcServer() *AppLaunch {
	app.initEnv()
	// 开启日志
	logger.InitLog(viper.GetString("zap.log-dir"), "zap.log-level")
	// 客户端配置 如果需要支持多namespace，我们可以场景多个client,
	// 如果开发环境使用本地配置固定端口，如果去非开发环境，则由shell脚本自动获取及某一个ip段的端口自动分配
	if app.Active == "dev" {
		// 获取本机
		localIp, _ := utils.GetTrueIp()
		app.Ip = cast.ToString(localIp)
		// port, _ := utils.GetAvailablePort()
		// app.Port = port
		app.Port = viper.GetUint64("application.port")
	} else {
		app.Ip = utils.GetIpV4()
		// 如果去非开发环境，则由shell脚本自动获取及某一个ip段的端口自动分配
		// 	port, _ := utils.GetAvailablePort()
		// 	app.Port = port
	}
	logger.SugarLogger.Infof("启动服务....., 启动环境:[%v], 配置路径:[%v], 运行端口: [%v]", app.Active, app.ConfigPath, app.Port)
	// 初始化证书信息
	logger.SugarLogger.Infof("证书加载地址：%v", app.ConfigPath)
	app.Credentials = config.InitServerCredentials(app.ConfigPath)
	app.GrpcServer = grpc.NewServer(grpc.Creds(app.Credentials))
	return app
}

// StartTcp 开启tcp服务
func (app *AppLaunch) StartTcp() {
	// 启动端口
	host := fmt.Sprintf(":%v", app.Port)
	// 指定端口
	listen, err := net.Listen("tcp", host)
	logger.SugarLogger.Infof(">>>>>>>>>> 启动TCP服务,端口为%v", host)
	if err != nil {
		logger.SugarLogger.Errorf("监听端口出错,错误信息：%v", err.Error())
		panic("监听端口出错")
	}

	initParam := coreModel.InitParam{
		NamespaceId: viper.GetString("nacos.config.namespace"),
		DataId:      viper.GetString("nacos.config.data-id"),
		GroupName:   viper.GetString("nacos.config.group"),
		ServiceName: viper.GetString("application.name"),
		IpAddr:      viper.GetString("nacos.config.server-addr"),
		Port:        viper.GetUint64("nacos.config.port"),
		LogDir:      viper.GetString("nacos.log-dir"),
		CacheDir:    viper.GetString("nacos.cache-dir"),
		LogLevel:    viper.GetString("nacos.log-level"),
		UserName:    viper.GetString("nacos.username"),
		Password:    viper.GetString("nacos.password"),
		RegisterIp:  app.Ip,
		ServerPort:  app.Port,
	}
	nacos, err := Nacos.InitNacos(initParam, app.Credentials)
	if err != nil {
		panic(err.Error())
	}
	// 启动服务
	err = app.GrpcServer.Serve(listen)
	if err != nil {
		logger.SugarLogger.Errorf(">>>>>> 启动服务出错,错误信息：%v", err.Error())
	}
	logger.SugarLogger.Infof(">>>>>> 启动启动成功,端口号：[ %v ]", app.Port)
	// 打印logo
	utils.PrintLogo(true, "")
	ch := make(chan os.Signal, 1)
	// SIGHUP	1	Term	终端控制进程结束(终端连接断开)
	// SIGINT	2	Term	用户发送INTR字符(Ctrl+C)触发
	// SIGQUIT	3	Core	用户发送QUIT字符(Ctrl+/)触发
	// SIGILL	4	Core	非法指令(程序错误、试图执行数据段、栈溢出等)
	// SIGABRT	6	Core	调用abort函数触发
	// SIGFPE	8	Core	算术运行错误(浮点运算错误、除数为零等)
	// SIGKILL	9	Term	无条件结束程序(不能被捕获、阻塞或忽略)
	// SIGSEGV	11	Core	无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)
	// SIGPIPE	13	Term	消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
	// SIGALRM	14	Term	时钟定时信号
	// SIGTERM	15	Term	结束程序(可以被捕获、阻塞或忽略)
	// SIGUSR1	30,10,16	Term	用户保留
	// SIGUSR2	31,12,17	Term	用户保留
	// SIGCHLD	20,17,18	Ign	子进程结束(由父进程接收)
	// SIGCONT	19,18,25	Cont	继续执行已经停止的进程(不能被阻塞)
	// SIGSTOP	17,19,23	Stop	停止进程(不能被捕获、阻塞或忽略)
	// SIGTSTP	18,20,24	Stop	停止进程(可以被捕获、阻塞或忽略)
	// SIGTTIN	21,21,26	Stop	后台程序从终端中读取数据时触发
	// SIGTTOU	22,22,27	Stop	后台
	signal.Notify(ch, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	s := <-ch
	logger.SugarLogger.Warnf(">>>>>> 监听退出信号: %+v", s.Signal)
	// 注销nacos实例
	nacos.deregisterInstance()
	// 先关闭监听socket,阻止新链接
	if err := listen.Close(); err != nil {
		logger.SugarLogger.Warnf(">>>> 关闭监听socket失败: %+v", err.Error())
	}
	// 阻止呼叫
	app.GrpcServer.GracefulStop()
	// 准备退出系统
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println(">>>>>> 5秒后服务退出")
	select {
	case <-ctx.Done():
		app.GrpcServer.Stop()
		log.Println(">>>>>> 服务退出")
	default:
		app.GrpcServer.Stop()
		log.Println(">>>>>> 服务退出")
	}
}

// listenAndServe 监听服务状态
func (app *AppLaunch) listenAndServe() (err error) {
	if app.shutdownFinished == nil {
		app.shutdownFinished = make(chan struct{})
	}

	err = app.Server.ListenAndServe()
	if err == http.ErrServerClosed {
		err = nil
	} else if err != nil {
		logger.SugarLogger.Fatalf("来自ListenAndServe的意外错误: %w", err)
		return
	}

	log.Println("等待服务停止完成...")
	<-app.shutdownFinished
	log.Println("服务停止成功")
	return
}

// watchServerExit 监听服务退出
func (app *AppLaunch) watchServerExit(timeout time.Duration) {
	var ch = make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	log.Printf("%d秒后程序将停止", timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := app.Server.Shutdown(ctx)
	if err != nil {
		log.Println("停止服务: " + err.Error())
	} else {
		log.Println("服务停止成功")
		close(app.shutdownFinished)
	}
}
