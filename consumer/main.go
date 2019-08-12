package main

import (
	"logcat/consumer/service"
)

func main() {

	// cfg := flag.String("c", "../config-local.toml", "配置文件路径")

	// flag.Parse()

	// params := config.NewConfiguration()

	// if len(*cfg) != 0 {
	// 	params.InitFromFile(*cfg)
	// } else {
	// 	fmt.Println("注意: 使用默认配置!")
	// }

	s := service.NewService()

	s.Start()

}
