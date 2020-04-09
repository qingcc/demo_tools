package zap

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

var sugarLogger *zap.SugaredLogger

func run() {
	//initLogger("./test.log", zapcore.DebugLevel)
	initLevel()
	defer sugarLogger.Sync()

	simpleHttpGet("http://www.baidu.com")
	simpleHttpGet("http://www.google.com")
}

func initLogger(filename string, enab zapcore.LevelEnabler) {
	writeSync := getLogWriter(filename)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSync, enab)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,    //mb
		MaxAge:     30,    //保留旧文件的最大天数
		MaxBackups: 10,    //保留旧文件的最大个数
		Compress:   false, //是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("trying to hit Get request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s Encoder:%s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for url %s", resp.Status, url)
	}
}

//日志分级存储
func initLevel() {
	infoWriter := getLogWriter("./logs/info.log")
	errorWriter := getLogWriter("./logs/error.log")
	encoder := getEncoder()
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}
