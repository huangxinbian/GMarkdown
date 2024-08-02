cd /d %~dp0

cd ../

::设置代理
SET GOPROXY=https://goproxy.cn

for %%i in ("%cd%") do set dirName= %%~ni

go mod init %dirName%

go mod tidy

SET GOOS=linux

SET GOARCH=amd64

SET CGO_ENABLED=0

wails build

pause
