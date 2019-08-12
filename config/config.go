package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	// 配置参数(通过配置文件)

	Port   string `toml:"port"`   // 服务器监听端口
	Mode   string `toml:"mode"`   // 服务器开发模式
	Daemon bool   `toml:"daemon"` // 后端化运行

	// 数据库相关参数
	DbUrl    string `toml:"database_url"`    // 数据库URL
	DbPort   string `toml:"database_port"`   // 数据库端口
	DbName   string `toml:"database_name"`   // 数据库名称
	DbUser   string `toml:"database_user"`   // 数据库用户名
	DbPasswd string `toml:"database_passwd"` // 数据库密码

	// Redis数据库相关参数
	RedisUrl      string `toml:"redis_url"`       // 数据库URL
	RedisPassword string `toml:"redis_password"`  //
	RedisPoolSize int    `toml:"redis_pool_size"` // 连接池大小

	// ZK相关
	ZookeeperServerList string `toml:"zookeeper_server_list"` // ZK服务器列表
	ZookeeperMasterNode string `toml:"zookeeper_master_node"` // ZK主服务器路径名称

	// 日志相关参数
	LogLevel string `toml:"log_level"`
	LogDest  string `toml:"log_dest"`
	LogDir   string `toml:"log_dir"`

	// 性能监控参数
	GraphiteEnabled bool   `toml:"graphite_enabled"`
	GraphiteHost    string `toml:"graphite_host"`
	GraphitePort    int    `toml:"graphite_port"`
	GraphitePrefix  string `toml:"graphite_prefix"`

	HttpListenMode bool `toml:"httplisten_mode"`

	ProfilePath    string `toml:"profile_path"`
	BatchServerUrl string `toml:"batchserver_url"`

	LogType []string `toml:"log_type"`
}

func NewConfiguration() *Configuration {

	return &Configuration{ // 在此提供默认值(为所有参数提供默认值)
		Port: "3111",
		Mode: "develop",

		LogDir:   "./log",
		LogLevel: "debug",
		LogDest:  "file",

		DbUrl:    "180.76.55.12",
		DbPort:   "3306",
		DbName:   "adwetec",
		DbUser:   "root",
		DbPasswd: "xieming4243054",

		ZookeeperServerList: "192.168.0.20:2181,192.168.0.25:2181,192.168.0.27:2181",
		ZookeeperMasterNode: "192.168.0.20",

		GraphiteEnabled: false,
		GraphiteHost:    "",
		GraphitePort:    0,
		GraphitePrefix:  "api.",

		HttpListenMode: false,

		ProfilePath: "./profile",

		BatchServerUrl: "127.0.0.1:3117",
	}

}
func (this *Configuration) InitFromFile(path string) { // 用于启动时加载配置文件

	if _, err := toml.DecodeFile(path, this); err != nil {
		panic(fmt.Sprintf("can't decode conf file: [%s]", path))
	}

}
func (this *Configuration) ReloadFromFile(path string) error { // 用于配置文件的重新加载

	if _, err := toml.DecodeFile(path, this); err != nil {
		return err
	}

	return nil
}
