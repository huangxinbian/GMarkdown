package apputil

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CustomLogger 是自定义日志器的结构体
type CustomLogger struct {
	logFile *os.File
}

// Init 初始化日志文件
func (c *CustomLogger) Init(filename string) error {
	if filename == "" {
		filename = filepath.Dir(os.Args[0]) + "/logs/wails.log"
	}

	dirname := filepath.Dir(filename)
	err := os.MkdirAll(dirname, 0755)
	if err != nil {
		return err
	}

	c.logFile, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Debug 实现了 logger.Logger 接口的 Debug 方法
func (c *CustomLogger) Debug(message string) {
	c.writeLog("DEBUG", message)
}

// Info 实现了 logger.Logger 接口的 Info 方法
func (c *CustomLogger) Info(message string) {
	c.writeLog("INFO", message)
}

// Warning 实现了 logger.Logger 接口的 Warning 方法
func (c *CustomLogger) Warning(message string) {
	c.writeLog("WARNING", message)
}

// Error 实现了 logger.Logger 接口的 Error 方法
func (c *CustomLogger) Error(message string) {
	c.writeLog("ERROR", message)
}

// Critical 实现了 logger.Logger 接口的 Critical 方法
func (c *CustomLogger) Critical(message string) {
	c.writeLog("CRITICAL", message)
}

// Fatal 实现了 logger.Logger 接口的 Fatal 方法
func (c *CustomLogger) Fatal(message string) {
	c.writeLog("FATAL", message)
	os.Exit(1)
}

// Panic 实现了 logger.Logger 接口的 Panic 方法
func (c *CustomLogger) Print(message string) {
	c.writeLog("Print", message)
	panic(message)
}

// Panic 实现了 logger.Logger 接口的 Panic 方法
func (c *CustomLogger) Trace(message string) {
	c.writeLog("Trace", message)
	panic(message)
}

// writeLog 写入日志到文件
func (c *CustomLogger) writeLog(level, message string) {
	logMessage := fmt.Sprintf("%s %s: %s", time.Now().Format("2006-01-02 15:04:05"), level, fmt.Sprintf(message))
	c.logFile.WriteString(logMessage + "\n")
}
