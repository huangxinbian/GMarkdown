package main

import (
	"embed"
	"os"

	. "GMarkdown/apputil"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	nowFilePath := os.Args[1:]
	if len(nowFilePath) != 0 {
		nowFile = nowFilePath[0]
	}

	app := NewApp("./app.ini")

	runOptions := &options.App{
		Title:  "Markdown编辑器 " + nowFile,
		Width:  app.iWidth,
		Height: app.iHeight,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(), //静态资源访问器
		},
		Logger:             clogger,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.INFO,
		// SingleInstanceLock: &options.SingleInstanceLock{//单例
		// 	UniqueId: "e3984e08-28dc-4e3d-b70a-45e961589cdc",
		// },
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		OnBeforeClose:    app.beforeClose,
		Bind: []interface{}{
			app,
		},
	}

	if app.isMaxWindow { //最大化
		runOptions.WindowStartState = options.Maximised
	}

	err := wails.Run(runOptions)

	if err != nil {
		println("Error:", err.Error())
	}
}
