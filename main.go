package main

import (
	"flag"
	"github.com/xuhengpw/go_sslmng/golang_common/lib"
	"github.com/xuhengpw/go_sslmng/router"
	"os"
	"os/signal"
	"syscall"
)

//endpoint dashboard后台管理  server代理服务器
//config ./conf/prod/ 对应配置文件夹

var (
	config   = flag.String("config", "", "input config file like ./conf/dev/")
)

func main() {
	flag.Parse()
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}

	lib.InitModule(*config)
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

		router.HttpServerStop()

}
