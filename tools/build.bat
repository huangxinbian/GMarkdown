cd /d %~dp0

cd ../

:: wails build -ldflags="-s -w" 压缩
wails build

pause