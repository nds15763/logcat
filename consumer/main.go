package main

import (
	"logcat/esserver/service"
)

func main() {

	s := service.NewService()

	s.Start()

	//1.启动service，初始化es和kafka

	//2.消费kafka 获取clickid

	//3.查询ES

	//4.返回查询结果

}
