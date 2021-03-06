package main

import (
    "github.com/labstack/echo/v4"
    "fmt"
	"github.com/ipfans/go-common/echoext"
	mlog "github.com/ipfans/go-common/logext"
 )
func init(){
    mlog.InitLog("./test.log", "debug")
}

func main(){
    echo:= echo.New()
    echo.Use(echoext.ExtStdContext())
    echo.GET("/", indexController)
    echo.Start(":1323")
}


func indexController(c echo.Context) error{
	fmt.Println("--------------mlog----start-------------------")

	mlog.Logger.Infow(c.Request().Context(), "测试日志", "err_key_1", "value_1")
	mlog.Logger.Info(c.Request().Context(),  "err_key_1", "value_1")
	mlog.Logger.Errorw(c.Request().Context(), "测试日志", "err_key_1", "value_1")
	mlog.Logger.Error(c.Request().Context(),  "err_key_1", "value_1")
    fmt.Println("--------------mlog------end-----------------")
    return nil
}