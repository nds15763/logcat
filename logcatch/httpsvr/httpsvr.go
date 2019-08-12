package httpsvr

import (
	stdContext "context"
	"logcat/logcatch/handler"
	"logcat/logcatch/httpctx"
	"logcat/logcatch/service/elastic"
	"logcat/logcatch/service/kafka"
	"logcat/tools/logrus"
	"os"
	"time"

	"github.com/kataras/iris"
)

type HttpServer struct {
	Es      *elastic.ElasticManager
	Ka      *kafka.Producer
	App     *iris.Application
	Context *httpctx.ServerContext
	Logger  *logrus.Logger
}

// 滑动窗口限流
type RateLimit struct {
	count       int64
	counterList []int64
}

//NewHttpServer构造方法
func NewHttpServer(es *elastic.ElasticManager, ka *kafka.Producer) *HttpServer {

	return &HttpServer{
		Es:  es,
		Ka:  ka,
		App: iris.New(),
		Context: &httpctx.ServerContext{
			EsManager: es,
		},
	}
}

//Start初始化
//注册handler及加载公共信息
func (this *HttpServer) Start() {

	this.LoadData()

	dispatcher := handler.NewDispatcher(handler.MaxWorker)
	dispatcher.Run()

	this.App.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("LogCat!")
	})

	// uploader 根据上报获取日志，日志中分出各种类型 （HTTP请求很耗时，可能会改成其他方式）
	this.App.Post("/addlog", this.Wrap(handler.AddLogJob))

	// pusher 通过APP端主动推送，进行日志回捞
	//this.App.Post("/pusher", this.Wrap(handler.StreamLog))

	err := this.App.Run(iris.Addr(":8093"), iris.WithoutServerError(iris.ErrServerClosed))

	if err != nil {

		this.Logger.Error("[http server] iris app run error " + err.Error())

		os.Exit(1)
	}

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		// 关闭所有的 host
		this.App.Shutdown(ctx)
	})

}

//LoadData加载公共参数
//加载ES及kafka指针到context，随参数传入handler
func (this *HttpServer) LoadData() {
	this.Context.EsManager = this.Es
	this.Context.Kafka = this.Ka
}

// Wrap中间件
//添加log及公共参数，传入handler
func (this *HttpServer) Wrap(handler func(ctx iris.Context, logger *logrus.Logger, context *httpctx.ServerContext)) func(ctx iris.Context) {

	return func(ctx iris.Context) {
		handler(ctx, this.Logger, this.Context)
	}
}

// API统一安全过滤
// func (this *HttpServer) RLimit(handler func(ctx iris.Context)) {

// 	// go func(){

// 	// 	//根据滑动窗口判断压力

// 	// 	//如果服务器压力很大或者业务并不需要实时返回值，则根据服务器状态判断直接返回true
// 	// }

// 	return func(ctx iris.Context) {
// 		handler(ctx, this.Logger, this.Context)
// 	}
// }
