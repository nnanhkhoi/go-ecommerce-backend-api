package main

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age:%d ", "TipGo", 40)

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipGo"), zap.Int("age", 40))

	encoder := getEncoderLog()
	sync := getWriterSync()

	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Info log", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	encodeConfig.TimeKey = "time"

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)

}

func getWriterSync() zapcore.WriteSyncer {
	logDir := ".log"
	logFile := logDir + "/log.txt"

	// Ensure the directory exists
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, os.ModePerm); err != nil {
			log.Printf("Failed to create log directory: %v", err)
			return nil
		}
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
		return nil
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)

}
