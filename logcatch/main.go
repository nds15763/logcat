package main

import (
	"fmt"
	"logcat/logcatch/service"
)

func main() {

	//logger

	// location, err := time.LoadLocation("Asia/Shanghai")

	// if err != nil {

	// 	println("本地时区初始化失败! " + err.Error())

	// 	os.Exit(1)
	// }

	// logger := logrus.NewLogger(params.LogLevel, params.LogDest, params.LogDir)

	// logger.Info("启动服务!")

	s, err := service.NewService()
	if err != nil {
		fmt.Println(err.Error())
		//logger.Error("service启动失败:%s",err.Error())
	}

	s.Start()

}
