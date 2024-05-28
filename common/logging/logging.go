package logging

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	setupOnce sync.Once
)

type AppLogger struct {
	aLogger *logrus.Logger
}

func NewAppLogger() *AppLogger {
	var appLogger = &AppLogger{}
	setupOnce.Do(func() {
		appLogger.aLogger = logrus.New()
		appLogger.aLogger.SetLevel(logrus.DebugLevel)
		appLogger.aLogger.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
		appLogger.aLogger.SetReportCaller(true)
		appLogger.aLogger.SetOutput(os.Stderr)
	})
	return appLogger
}

// NewLoggingField creates a new logrus.Fields object with the given context and code
func NewLoggingField(url string, code int) *logrus.Fields {
	return &logrus.Fields{
		// url is extracted from context
		"url": url,
		// code is the Error code.
		"code": code,
	}
}

func (l *AppLogger) Info(ctx *gin.Context, code int, msg string) {
	if ctx != nil {
		l.aLogger.WithFields(*NewLoggingField(ctx.Request.URL.Path, code)).Info(msg)
	}
	l.aLogger.Info(msg)
}

func (l *AppLogger) Debug(ctx *gin.Context, code int, msg string) {
	if ctx != nil {
		l.aLogger.WithFields(*NewLoggingField(ctx.Request.URL.Path, code)).Debug(msg)
	}
	l.aLogger.Debug(msg)
}

func (l *AppLogger) Error(ctx *gin.Context, code int, msg string) {
	if ctx != nil {
		l.aLogger.WithFields(*NewLoggingField(ctx.Request.URL.Path, code)).Error(msg)
	}
	l.aLogger.Error(msg)
}
func (l *AppLogger) Warn(ctx *gin.Context, code int, msg string) {
	if ctx != nil {
		l.aLogger.WithFields(*NewLoggingField(ctx.Request.URL.Path, code)).Warn(msg)
	}
	l.aLogger.Warn(msg)
}
func (l *AppLogger) Fatal(ctx *gin.Context, code int, msg string) {
	if ctx != nil {
		l.aLogger.WithFields(*NewLoggingField(ctx.Request.URL.Path, code)).Fatal(msg)
	}
	l.aLogger.Fatal(msg)
}

func (l *AppLogger) Panic(ctx *gin.Context, code int, msg string) {
	if ctx != nil {
		l.aLogger.WithFields(*NewLoggingField(ctx.Request.URL.Path, code)).Panic(msg)
	}
	l.aLogger.Panic(msg)
}
