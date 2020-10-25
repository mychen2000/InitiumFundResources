package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/initiumfund/artis-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func DefaultEncoderCfg() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "ts",
		NameKey:          "logger",
		CallerKey:        "caller",
		FunctionKey:      "",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeName:       nil,
		ConsoleSeparator: "",
	}
}

func SetupLogger(cfg *config.Config) (*zap.SugaredLogger, error) {
	var cores []zapcore.Core

	if cfg.Log.FileName != "" {
		if cfg.Log.LogLevel == "" {
			cfg.Log.LogLevel = "./logs/"
		}

		logFilePath, err := filepath.Abs(path.Join(cfg.Log.Directory, cfg.Log.FileName+".log"))
		if err != nil {
			return nil, err
		}

		logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}

		// 本地日志文件
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(DefaultEncoderCfg()), logFile, ParseLogLevel(cfg)))
	}

	// 终端 stdout
	cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(DefaultEncoderCfg()), os.Stdout, ParseLogLevel(cfg)))

	logCores := zapcore.NewTee(cores...)
	logger := zap.New(logCores, zap.AddStacktrace(zapcore.ErrorLevel))

	if cfg.IsDevelopment() {
		logger.WithOptions(zap.Development())
	}

	sugaredLogger := logger.Sugar()
	return sugaredLogger, nil
}

func ParseLogLevel(cfg *config.Config) zapcore.Level {
	switch cfg.Log.LogLevel {
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	case "error":
		return zapcore.ErrorLevel
	case "warn":
		return zapcore.WarnLevel
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

func StatusColor(statusCode int) string {
	switch {
	case statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices:
		return green
	case statusCode >= http.StatusMultipleChoices && statusCode < http.StatusBadRequest:
		return white
	case statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

func MethodColor(method string) string {
	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

func Logger(log *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqPath := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		method := c.Request.Method
		start := time.Now()

		c.Next()

		stop := time.Since(start)
		cost := stop.Milliseconds()
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()

		entry := log.With(
			"statusCode", statusCode,
			"cost", cost,
			"clientIP", clientIP,
			"method", method,
			"path", reqPath,
			"query", query,
			"userAgent", clientUserAgent,
			"handler", c.HandlerName(),
		)

		// Print errors if it is an error
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.String())
		}

		// Prettify outputs
		statusColor := StatusColor(statusCode)
		methodColor := MethodColor(method)
		resetColor := reset

		queryUrl := query
		if queryUrl != "" {
			queryUrl = "?" + queryUrl
		}

		msg := fmt.Sprintf("|%s %3d %s| %13v |%s %-7s %s %s",
			statusColor, statusCode, resetColor,
			stop,
			methodColor, method, resetColor,
			reqPath+queryUrl,
		)

		if statusCode >= 500 {
			entry.Error(msg)
		} else if statusCode >= 400 {
			entry.Warn(msg)
		} else {
			entry.Info(msg)
		}
	}
}
