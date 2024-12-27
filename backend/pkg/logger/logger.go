package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/alex-guoba/gin-clean-template/pkg/setting"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *lumberjack.Logger

func SetupLogger(logset *setting.LogSettingS) {
	// TODO: use lumberjack.Logger as config
	logger := &lumberjack.Logger{
		Filename:   filepath.Join(logset.LogSavePath, logset.LogFileName),
		MaxSize:    logset.MaxSize, // megabytes
		MaxBackups: logset.MaxBackups,
		MaxAge:     3,               // days
		Compress:   logset.Compress, // disabled by default
	}
	lvl, err := log.ParseLevel(logset.Level)
	if err != nil {
		log.SetLevel(lvl)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// use lumberjack to write to implement rotation.
	mw := io.MultiWriter(os.Stdout, logger)
	log.SetOutput(mw) // set output to file and console at the same time

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func Close() {
	logger.Close()
}

func WithTrace(ctx *gin.Context) *log.Entry {
	fields := log.Fields{}
	if len(ctx.GetString("X-Trace-ID")) > 0 {
		fields["trace_id"] = ctx.GetString("X-Trace-ID")
	}
	if len(ctx.GetString("X-Span-ID")) > 0 {
		fields["span_id"] = ctx.GetString("X-Span-ID")
	}
	return log.WithFields(fields)
}
