package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	. "GMarkdown/apputil"

	sr "runtime"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/ini.v1"
)

var (
	nowFile     string
	clogger     *CustomLogger
	iniFilePath string
)

// App struct
type App struct {
	ctx         context.Context
	iWidth      int
	iHeight     int
	iTop        int
	iLeft       int
	isMaxWindow bool
}

// NewApp creates a new App application struct
func NewApp(sIniFilePath string) *App {
	result := &App{}
	clogger = &CustomLogger{}
	clogger.Init("")                    //初始化日志
	result.createInitFile(sIniFilePath) //创建配置文件
	result.ReadInit()                   //读取配置文件

	return result
}

// 创建配置文件
func (a *App) createInitFile(sPath string) {
	if strings.HasPrefix(sPath, "./") {
		if sr.GOOS == "windows" {
			sPath = filepath.Dir(os.Args[0]) + "/" + sPath[2:] //获取程序真实目录
		} else {
			exePath, _ := os.Executable()
			sPath = exePath + "/" + sPath[2:]
		}
	}

	// 获取文件的完整路径
	var err error
	iniFilePath, err = filepath.Abs(sPath)
	if err != nil {
		clogger.Error("获取绝对路径失败:" + err.Error())
	}

	_, err = os.Stat(iniFilePath) // 尝试获取文件信息
	if err == nil {               // 如果文件存在并且没有错误发生，则说明存在配置文件
		return
	}

	// 检查文件夹是否存在，如果不存在则创建
	dir := filepath.Dir(iniFilePath)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		clogger.Error("创建目录失败:" + err.Error())
	}

	// 打开文件，如果文件不存在，os.OpenFile会创建它
	file, err := os.OpenFile(iniFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		clogger.Error("打开或创建文件失败:" + err.Error())
		return
	}

	file.WriteString("")
	file.Close()

	cfg, err := ini.Load(iniFilePath)
	if err != nil {
		clogger.Error("读取配置文件异常:" + err.Error())
		return
	}

	cfg.Section("App").Key("iWidth").SetValue("960")
	cfg.Section("App").Key("iHeight").SetValue("600")
	cfg.Section("App").Key("iLeft").SetValue("-1024")
	cfg.Section("App").Key("iTop").SetValue("-1024")
	cfg.Section("App").Key("isMaxWindow").SetValue("0")
	cfg.SaveTo(iniFilePath)
}

func (a *App) ReadInit() {
	cfg, err := ini.Load(iniFilePath)
	if err != nil {
		clogger.Error("读取配置文件异常:" + err.Error())
		return
	}

	section := cfg.Section("App")
	a.iWidth, err = section.Key("iWidth").Int()
	if err != nil {
		a.iWidth = 900
	}

	a.iHeight, err = section.Key("iHeight").Int()
	if err != nil {
		a.iHeight = 600
	}

	a.iTop, err = section.Key("iTop").Int()
	if err != nil {
		a.iTop = -1024
	}

	a.iLeft, err = section.Key("iLeft").Int()
	if err != nil {
		a.iLeft = -1024
	}

	temp, _ := section.Key("isMaxWindow").Int()
	a.isMaxWindow = temp == 1
}

// 一旦 Wails 分配了它需要的资源，就会调用 startup 方法，它是创建资源、设置事件侦听器以及应用程序在启动时需要的任何其他东西的好地方
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	if (a.iTop == -1024) || (a.iLeft == -1024) {
		runtime.WindowCenter(a.ctx) //居中
		return
	}

	runtime.WindowSetPosition(a.ctx, a.iLeft, a.iTop) //指定位置
}

// 在关闭过程结束时由 Wails 调用。 这是释放内存和执行关闭任务的好地方
func (a *App) shutdown(ctx context.Context) {
	// saveWindowInfo(ctx)//无法获得window信息,不能在这里调用
}

func saveWindowInfo(ctx context.Context) {
	cfg, err := ini.Load(iniFilePath)
	if err != nil {
		clogger.Error("读取配置文件异常:" + err.Error())
		return
	}

	if runtime.WindowIsMaximised(ctx) { //窗口最大化情况下不能保存窗口位置和大小,否则会都是最大化
		cfg.Section("App").Key("isMaxWindow").SetValue("1")
		cfg.SaveTo(iniFilePath)
		return
	}

	iWidth, iHeight := runtime.WindowGetSize(ctx)
	iLeft, iTop := runtime.WindowGetPosition(ctx)

	cfg.Section("App").Key("iWidth").SetValue(strconv.Itoa(iWidth))
	cfg.Section("App").Key("iHeight").SetValue(strconv.Itoa(iHeight))
	cfg.Section("App").Key("iLeft").SetValue(strconv.Itoa(iLeft))
	cfg.Section("App").Key("iTop").SetValue(strconv.Itoa(iTop))
	cfg.Section("App").Key("isMaxWindow").SetValue("0")
	cfg.SaveTo(iniFilePath)
}

// 应用关闭前回调
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	saveWindowInfo(ctx) //保存窗口信息

	return false //true 将导致应用程序继续;false 将继续正常关闭
}

// 读取当前文件内容
func (a *App) ReadNowFile() *MsgEmity {
	if nowFile == "" {
		return MsgEmity{}.Err(8001, "路径参数缺失")
	}

	return a.ReadFile(nowFile)
}

// 读取指定文件内容
func (a *App) ReadFile(path string) *MsgEmity {
	if path == "" {
		return MsgEmity{}.Err(8001, "路径参数缺失")
	}

	//读取文件内容
	file, err := os.Open(path)
	if err != nil {
		return MsgEmity{}.Err(8002, "打开文件异常:"+err.Error())
	}

	defer file.Close()

	//读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		return MsgEmity{}.Err(8003, "读取文件异常:"+err.Error())
	}

	return MsgEmity{}.Success(string(content), path)
}

// 保存文件内容
func (a *App) SaveFile(path string, text string) *MsgEmity {
	if path == "" {
		return MsgEmity{}.Err(8001, "路径参数缺失")
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return MsgEmity{}.Err(8002, "打开文件异常:"+err.Error())
	}

	defer file.Close()
	_, err = file.WriteString(text)
	if err != nil {
		return MsgEmity{}.Err(8003, "保存文件异常:"+err.Error())
	}

	return MsgEmity{}.Success(path, "保存文件成功")
}

// 另存文件
func (a *App) SaveAsFile(defaultPath string, text string) *MsgEmity {
	dialogOptions := runtime.SaveDialogOptions{
		DefaultDirectory: defaultPath,
		Title:            "选择保存位置",
		Filters: []runtime.FileFilter{{
			DisplayName: "Markdown",
			Pattern:     "*.md",
		}},
		CanCreateDirectories: true,
	}

	sFilePath, err := runtime.SaveFileDialog(a.ctx, dialogOptions)
	if err != nil {
		return MsgEmity{}.Err(8001, "选择存储位置异常:"+err.Error())
	}

	file, err := os.OpenFile(sFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return MsgEmity{}.Err(8003, "打开文件异常:"+err.Error())
	}

	defer file.Close()
	_, err = file.WriteString(text)
	if err != nil {
		return MsgEmity{}.Err(8004, "保存文件异常:"+err.Error())
	}

	nowFile = sFilePath
	runtime.WindowSetTitle(a.ctx, "Markdown编辑器 "+nowFile)

	return MsgEmity{}.Success(sFilePath, "保存文件成功")
}

// 打开文件
func (a *App) OpenFile(defaultPath string) *MsgEmity {
	dialogOptions := runtime.OpenDialogOptions{
		DefaultDirectory: defaultPath,
		Title:            "选择文件",
		Filters: []runtime.FileFilter{{
			DisplayName: "Markdown",
			Pattern:     "*.md",
		}, {
			DisplayName: "所有文件",
			Pattern:     "*.*",
		}},
	}

	sFilePath, err := runtime.OpenFileDialog(a.ctx, dialogOptions)
	if err != nil {
		return MsgEmity{}.Err(8001, "打开文件异常:"+err.Error())
	}

	if sFilePath == "" {
		return MsgEmity{}.Err(8002, "文件路径为空")
	}

	nowFile = sFilePath

	//读取文件内容
	file, err := os.Open(sFilePath)
	if err != nil {
		return MsgEmity{}.Err(8003, "试图打开文件异常:"+err.Error())
	}

	defer file.Close()

	//读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		return MsgEmity{}.Err(8001, "读取文件异常:"+err.Error())
	}

	nowFile = sFilePath
	runtime.WindowSetTitle(a.ctx, "Markdown编辑器 "+nowFile)

	return MsgEmity{}.Success(string(content), sFilePath)
}

// Info消息
func (a *App) InfoMsg(sTitle, sMessage string) {
	dialogOptions := runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         sTitle,
		Message:       sMessage,
		DefaultButton: "确定",
		//CancelButton:  string
	}

	runtime.MessageDialog(a.ctx, dialogOptions)
}

// Warning消息
func (a *App) WarningMsg(sTitle, sMessage string) {
	dialogOptions := runtime.MessageDialogOptions{
		Type:          runtime.WarningDialog,
		Title:         sTitle,
		Message:       sMessage,
		DefaultButton: "确定",
		//CancelButton:  string
	}

	runtime.MessageDialog(a.ctx, dialogOptions)
}

// Error消息
func (a *App) ErrorMsg(sTitle, sMessage string) {
	dialogOptions := runtime.MessageDialogOptions{
		Type:          runtime.ErrorDialog,
		Title:         sTitle,
		Message:       sMessage,
		DefaultButton: "确定",
		//CancelButton:  string
	}

	runtime.MessageDialog(a.ctx, dialogOptions)
}

// Question消息
func (a *App) QuestionMsg(sTitle, sMessage string) string {
	dialogOptions := runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         sTitle,
		Message:       sMessage,
		DefaultButton: "确定",
		CancelButton:  "取消",
	}

	result, _ := runtime.MessageDialog(a.ctx, dialogOptions)

	return result
}

// Question消息
func (a *App) WindowPrint() string {
	runtime.WindowPrint(a.ctx)

	return "响应"
}

func (a *App) SecondInstanceLaunch() {
	fmt.Println("onSecondInstanceLaunch")
}
