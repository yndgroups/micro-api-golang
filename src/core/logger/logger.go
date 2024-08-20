package logger

import (
	"io"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 只能输出结构化日志，但是性能要高于 SugaredLogger
var logger *zap.Logger

// SugarLogger 可以输出 结构化日志、非结构化日志。性能茶语 zap.Logger，具体可见上面的的单元测试
var SugarLogger *zap.SugaredLogger

// InitLog 初始化日志 logger
func InitLog(logPath string, level string) {
	var lv zapcore.Level
	switch level {
	case "DebugLevel":
		lv = zap.DebugLevel
	case "InfoLevel":
		lv = zap.InfoLevel
	case "WarnLevel":
		lv = zap.WarnLevel
	case "ErrorLevel":
		lv = zap.ErrorLevel
	case "DPanicLevel":
		lv = zap.DPanicLevel
	case "PanicLevel":
		lv = zap.PanicLevel
	case "FatalLevel":
		lv = zap.FatalLevel
	default:
		lv = zap.InfoLevel
	}
	if ok, _ := PathExists(logPath); !ok {
		os.Mkdir(logPath, os.ModePerm)
	}
	config := zapcore.EncoderConfig{
		MessageKey:   "msg",                       //结构化（json）输出：msg的key
		LevelKey:     "level",                     //结构化（json）输出：日志级别的key（INFO，WARN，ERROR等）
		TimeKey:      "ts",                        //结构化（json）输出：时间的key（INFO，WARN，ERROR等）
		CallerKey:    "file",                      //结构化（json）输出：打印日志的文件对应的Key
		EncodeLevel:  zapcore.CapitalLevelEncoder, //将日志级别转换成大写（INFO，WARN，ERROR等）
		EncodeCaller: zapcore.ShortCallerEncoder,  //采用短文件路径编码输出（test/testmain.go:14	）
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}, //输出的时间格式
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		}, //
	}
	//自定义日志级别：自定义Info级别
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= lv
	})

	//自定义日志级别：自定义Warn级别
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= lv
	})

	// 获取io.Writer的实现
	infoWriter := getWriter(logPath + "info.log")
	warnWriter := getWriter(logPath + "error.log")

	// 实现多个输出
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(infoWriter), infoLevel),                   //将info及以下写入logPath，NewConsoleEncoder 是非结构化输出
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(warnWriter), warnLevel),                   //warn及以上写入errPath
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), lv), //同时将日志输出到控制台，NewJSONEncoder 是结构化输出
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	SugarLogger = logger.Sugar()
}

//
//func getWriter(filename string) io.Writer {
//	return &lumberjack.Logger{
//		Filename:   filename,
//		MaxSize:    10,  //最大M数，超过则切割
//		MaxBackups: 5,   //最大文件保留数，超过就删除最老的日志文件
//		MaxAge:     30,  //保存30天
//		Compress:   false,//是否压缩
//	}
//}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 filename.YYmmddHH
	// filename是指向最新日志的链接
	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),    // 保存30天
		rotatelogs.WithRotationTime(time.Hour*24), //切割频率 24小时
	)
	if err != nil {
		log.Println("日志启动异常")
		panic(err)
	}
	return hook
}

// PathExists 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	// IsNotExist 来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}
