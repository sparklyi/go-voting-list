package logger

import "github.com/sirupsen/logrus"

func init() {
	//todo 初始化设置日志格式
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2000-01-01 12:34:56"})
	//todo logrus.SetReportCaller(true)设置在输出日志中添加文件名和方法信息
	logrus.SetReportCaller(false)
}
