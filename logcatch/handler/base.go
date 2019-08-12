package handler

import (
	"context"
	"fmt"
	"logcat/logcatch/handler/logg"
	"logcat/logcatch/httpctx"
	"logcat/tools/logrus"
	"time"

	"github.com/kataras/iris"
)

var (
	MaxWorker = 10
)

type LogJob struct {
	LogLevel   int64  `json:"log_level"`   //日志级别，error，warning，info
	LogHeader  string `json:"log_header"`  //标题
	LogContent string `json:"log_content"` //内容
	Platform   string `json:"platform"`
	request    *httpctx.IrisRequest
}

//待执行的工作
type Job struct {
	LogJob LogJob
}

//任务channal
var JobQueue chan Job

//执行任务的工作者单元
type Worker struct {
	WorkerPool chan chan Job //工作者池--每个元素是一个工作者的私有任务channal
	JobChannel chan Job      //每个工作者单元包含一个任务管道 用于获取任务
	quit       chan bool     //退出信号
	no         int           //编号
}

//创建一个新工作者单元
func NewWorker(workerPool chan chan Job, no int) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		no:         no,
	}
}

//循环  监听任务和结束信号
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:

				switch job.LogJob.Platform {
				case "android":
					go w.Wrap(job.LogJob.request, logg.AndroidLog)
				case "ios":
					go w.Wrap(job.LogJob.request, logg.IosLog)
				case "stream":
					go w.Wrap(job.LogJob.request, logg.StreamLog)
				default:
				}

			case <-w.quit:
				// 收到退出信号
				return
			}
		}
	}()
}

// 停止信号
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func (w Worker) Wrap(request *httpctx.IrisRequest, handler func(ctx iris.Context, logger *logrus.Logger, context *httpctx.ServerContext)) {
	//验证或者做其他input转换
	timeout, cancel := context.WithTimeout(request.Ctx.(context.Context), 20*time.Second)

	go handler(request.Ctx, request.Logger, request.Context)

	//
	select {
	case <-timeout.Done():
		cancel()
	}
	//

}

//调度中心
type Dispatcher struct {
	//工作者池
	WorkerPool chan chan Job
	//工作者数量
	MaxWorkers int
}

//创建调度中心
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}

//工作者池的初始化
func (d *Dispatcher) Run() {
	for i := 1; i < d.MaxWorkers+1; i++ {
		worker := NewWorker(d.WorkerPool, i)
		worker.Start()
	}
	go d.dispatch()
}

//调度
func (d *Dispatcher) dispatch() {
	JobQueue = make(chan Job, 10)
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("job := <-JobQueue:")
			go func(job Job) {
				fmt.Println("等待空闲worker (任务多的时候会阻塞这里")
				//等待空闲worker (任务多的时候会阻塞这里)
				jobChannel := <-d.WorkerPool
				// 将任务放到上述woker的私有任务channal中
				jobChannel <- job

			}(job)
		}
	}
}
