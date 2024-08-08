package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "logs/blog"
	linkName := "blog-log"
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(
		filePath+".%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),     //设置最大存储时间
		rotatelogs.WithRotationTime(24*time.Hour), //设置分割时间
		rotatelogs.WithLinkName(linkName),         //设置软连接，类似快捷方式
	)
	if err != nil {
		logger.Errorf("logrus出错. %+v", errors.WithStack(err))
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		spendTime := fmt.Sprintf("%.2fms", float32(stop)/float32(time.Millisecond))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.URL.Path

		entry := logger.WithFields(logrus.Fields{
			"hostName":   hostName,
			"statusCode": statusCode,
			"spendTime":  spendTime,
			"clientIP":   clientIP,
			"userAgent":  userAgent,
			"dataSize":   dataSize,
			"method":     method,
			"path":       path,
		})
		//检测错误，如果gin内部的Errors列表不为0，即存在error
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			//在entry内写入错误信息
			//ByType会根据类型筛选出gin.ErrorTypePrivate的错误类型
			//该错误类型是内部服务器错误、数据库错误等内容
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
