package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/privatesquare/bkst-go-utils/utils/dateutils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var (
	logger                 *zap.Logger
	DefaultLogLevel        = "INFO"
	debugLogLevel          = "DEBUG"
	AuthorizationHeaderKey = "Authorization"
)

func init() {
	SetLoggerConfig(GetLoggerConfig(DefaultLogLevel))
}

func GetLoggerConfig(logLevel string) zap.Config {
	var zapLogLevel zapcore.Level
	if logLevel == debugLogLevel {
		zapLogLevel = zap.DebugLevel
	} else {
		zapLogLevel = zap.InfoLevel
	}
	return zap.Config{
		Level:    zap.NewAtomicLevelAt(zapLogLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{"stdout"},
	}
}

func SetLoggerConfig(cfg zap.Config) {
	var err error
	if logger, err = cfg.Build(); err != nil {
		log.Fatalln("Unable to initialize logger: ", err)
	}
}

func Info(msg string, tags ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, tags...)
}

func Infof(format string, a ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, a...))
}

func Warn(msg string, tags ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Warn(msg, tags...)
}

func Error(msg string, err error, tags ...zapcore.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, tags...)
}

func Debug(msg string, tags ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(2)).Debug(msg, tags...)
}

func Panic(msg string, tags ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Panic(msg, tags...)
}

// GinZap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
func GinZap() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := dateutils.GetDateTimeNow()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		end := dateutils.GetDateTimeNow()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.WithOptions(zap.AddCallerSkip(1)).Error(e)
			}
		} else {
			logger.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("latency", latency.String()),
			)
		}
	}
}

func RestyDebugLogs(resp *resty.Response) {
	resp.Request.Header[AuthorizationHeaderKey] = []string{}
	Debug(fmt.Sprintf("Request: %v", resp.Request))
	Debug(fmt.Sprintf("Request Url: %v", resp.Request.URL))
	Debug(fmt.Sprintf("Request Header: %v", resp.Request.Header))
	Debug(fmt.Sprintf("Request Body: %v", resp.Request.Body))
	Debug(fmt.Sprintf("Response Body: %v", string(resp.Body())))
}
