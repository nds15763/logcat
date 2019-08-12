package httpsvr

import (
	"logcat/consumer/handler"
	"logcat/consumer/httpctx"

	"logcat/consumer/service/elastic"
	"logcat/tools/logrus"
	"os"

	"github.com/kataras/iris"
)

type HttpServer struct {
	Es      *elastic.ElasticManager
	App     *iris.Application
	Context *httpctx.ServerContext
	Logger  *logrus.Logger
}

//HTTPserver包含几大部分功能
//1：路由
//2：保存session，因为session本质上也是存储在内存之中，golang也并没有原生支持session，所以可以直接将所有信息直接保存在系统里
//	也可以存储静态表，也可以存储用户信息
//3：保存service层及其他信息
func NewHttpServer(es *elastic.ElasticManager) *HttpServer {

	return &HttpServer{
		Es:  es,
		App: iris.New(),
		Context: &httpctx.ServerContext{
			EsManager: es,
		},
	}
}

func (this *HttpServer) Start() {

	this.LoadData()

	kaMap := handler.NewKafkaMap([]string{"android", "ios", "stream"})
	kaMap.Run()

	this.App.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("Hello world!")
	})

	this.App.Post("/findes", this.Wrap(handler.FindEs))

	err := this.App.Run(iris.Addr(":8090"), iris.WithoutServerError(iris.ErrServerClosed))

	if err != nil {

		this.Logger.Error("[http server] iris app run error " + err.Error())

		os.Exit(1)
	}

}

func (this *HttpServer) LoadData() {
	this.Context.EsManager = this.Es
}

func (this *HttpServer) Wrap(handler func(ctx iris.Context, logger *logrus.Logger, context *httpctx.ServerContext)) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		handler(ctx, this.Logger, this.Context)
	}
}
