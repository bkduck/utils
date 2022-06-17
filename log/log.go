package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	//"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func InitLogger(savepath string, debug bool) {
	var err error
	if debug {
		logger, err = zap.NewDevelopment()
		if err != nil {
			panic(fmt.Sprintf("Got error when init logger,  the error is '%v'", err))
		}
	} else {
		logger, err = zap.NewDevelopment()
		//now := time.Now()
		//hook := &lumberjack.Logger{
		//	Filename:   fmt.Sprintf("%s/%s-%04d%02d%02d%02d.log", savepath, module, now.Year(), now.Month(), now.Day(), now.Hour()), //filePath
		//	MaxSize:    500,
		//	MaxBackups: 100,
		//	MaxAge:     180,   //days
		//	Compress:   false, // disabled by default
		//}
		//defer hook.Close()

		//writer := zapcore.AddSync(hook)
		if savepath == ""{
			savepath = "."
		}
		writer := getLogWriter(savepath)
		core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), writer, zap.InfoLevel)
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		if err != nil {
			panic(fmt.Sprintf("Got error when init logger,  the error is '%v'", err))
		}
	}
	//fmt.Println(logger)
	//sugar := logger.Sugar()
	//sugar.Infow("failed to fetch URL",
	//	"url", url,
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
}

func Sync()  {
	 logger.Sync()
}

func Info(msg string, fields ...zap.Field)  {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field)  {
	logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field)  {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field)  {
	logger.Error(msg, fields...)
}

func Fatal (msg string, fields ...zap.Field)  {
	logger.Fatal(msg, fields...)
}

func getLogWriter(savepath string) zapcore.WriteSyncer {
	file, _ := os.Create(savepath + "/test.log")
	return zapcore.AddSync(file)
}




