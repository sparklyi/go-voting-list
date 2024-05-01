package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"
)

func init() {
	//todo 初始化设置日志格式
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	//todo logrus.SetReportCaller(true)设置在输出日志中添加文件名和方法信息
	logrus.SetReportCaller(false)
}

// setOutPutFilePath 设置log文件存放位置
func setOutPutFilePath(filename string, level logrus.Level) {

	//查询存不存在此文件
	//err != nil
	if _, err := os.Stat("./runtime/log"); os.IsNotExist(err) {
		//创建此文件夹
		err = os.MkdirAll("./runtime/log", 0777)
		if err != nil {
			panic("create log folder failed" + err.Error())
		}
	}
	//设置时间格式
	t := time.Now().Format("2006-01-02")
	//拼接文件名
	fileName := path.Join("./runtime/log", filename+"_"+t+".log")

	//写入日志
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//获取当前所在路径
		p, _ := os.Getwd()
		panic("open file failed : " + p)
	}
	//没有打开权限，关闭后后面无法写入
	//defer file.Close()
	//设置日志输出位置
	logrus.SetOutput(file)
	//设置日志
	logrus.SetLevel(level)

}

// ToFile 将日志写入到文件中
func ToFile() gin.LoggerConfig {

	//判断文件夹是否存在

	if _, err := os.Stat("./runtime/log"); err != nil {
		//create folder
		err = os.MkdirAll("./runtime/log", 0777)
		if err != nil {
			panic("create log folder failed" + err.Error())
		}
	}
	t := time.Now().Format("2006-01-02")
	filename := path.Join("./runtime/log", "visit_"+t+".log")

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("file open failed" + err.Error())
	}

	conf := gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			//将主机信息拼接成字符串返回
			return fmt.Sprintln(
				params.TimeStamp.Format("2006-01-02 15:04:05"),
				params.ClientIP,
				params.Request.Proto,
				params.Request.UserAgent(),
				params.Request.Header,
				params.ErrorMessage,
				params.Method,
				params.StatusCode,
			)
		},
		//写入到file文件所表示的文件中
		Output: file,
	}

	return conf
}

// Recover 恢复panic异常 记录错误日志
func Recover(c *gin.Context) {

	defer func() {
		if panic_err := recover(); panic_err != nil {
			//panic时处理
			if _, err := os.Stat("./runtime/log"); err != nil {
				//不存在创建文件夹
				err = os.MkdirAll("./runtime/log", 0777)
				if err != nil {
					fmt.Println("create folder failed")
				}
			}
			t := time.Now().Format("2006-01-02")
			filename := path.Join("./runtime/log", "error_"+t+".log")
			file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			defer file.Close()
			//写入数据
			//todo 记录panic时间
			file.WriteString("panic error time:" + time.Now().Format("2006-01-02 15:04:05") + "=================\n")

			//todo 记录错误信息panic_eer
			file.WriteString(fmt.Sprintln(panic_err))
			file.WriteString("stacktrace from panic:" + string(debug.Stack()) + "\n")
			//向请求返回服务器错误
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": panic_err,
			})
			//todo 终止本次请求的后续处理
			c.Abort()
		}
	}()
	//将控制权传递给下一个中间件或处理程序
	c.Next()
}

// Write 自使用日志，记录msg
func Write(msg string, filename string) {

	//设置输出文件位置，等级为info
	setOutPutFilePath(filename, logrus.InfoLevel)
	//写入日志
	logrus.Info(msg)

}

//todo 记录不同等级的日志，严重程度从高到低，高级触发后低级不会触发
/*
	panic 恐慌日志,最高级别
	fatal 致命错误
	error 错误
	warn  警告
	info  运行信息
	debug 调试信息
	trace 追踪日志
*/

// Panic 记录panic等级的日志
func Panic(field logrus.Fields, msg ...any) {

	//设置输出位置
	setOutPutFilePath("panic", logrus.PanicLevel)
	//将日志字段和msg全部记录到日志中
	logrus.WithFields(field).Panic(msg)

}

// Fatal 记录fatal等级的日志
func Fatal(field logrus.Fields, msg ...any) {

	//设置输出位置
	setOutPutFilePath("Fatal", logrus.FatalLevel)
	//将日志字段和msg全部记录到日志中
	logrus.WithFields(field).Fatal(msg)

}

// Error 记录Error等级的日志
func Error(fields logrus.Fields, msg ...any) {

	setOutPutFilePath("Error", logrus.ErrorLevel)
	logrus.WithFields(fields).Error(msg)
}

// Warn 记录warning等级的日志
func Warn(fields logrus.Fields, msg ...any) {

	setOutPutFilePath("Warn", logrus.WarnLevel)
	logrus.WithFields(fields).Warn(msg)

}

// Info  记录Info等级的日志
func Info(fields logrus.Fields, msg ...any) {
	setOutPutFilePath("info", logrus.InfoLevel)
	logrus.WithFields(fields).Info(msg)
}

// Debug 记录Debug等级的日志
func Debug(fields logrus.Fields, msg ...any) {

	setOutPutFilePath("Debug", logrus.DebugLevel)
	logrus.WithFields(fields).Debug(msg)
}

// Trace 记录Trace等级的日志
func Trace(fields logrus.Fields, msg ...any) {

	setOutPutFilePath("Trace", logrus.TraceLevel)
	logrus.WithFields(fields).Trace(msg)

}
