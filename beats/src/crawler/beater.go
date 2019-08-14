package crawler

import (
	"beats/src/config"
	"beats/src/model"
)

// //对于每一个类型进行fetch
// type Beater interface {
// 	Config() error
// 	Setup() error
// 	Run() error
// 	Stop()
// }

type Beater struct {
	cfg      config.LogType
	FileName string
	Wf       WatchFlags //监控文件变化
	PushChan chan model.Beat
}

func NewBeater(cfg config.LogType, fileName string) *Beater {
	return &Beater{
		cfg:      cfg,
		FileName: fileName}
}

func (this *Beater) Fetch(fname string) {
	//查找所有文件中符合正则的内容
	//若没有正则，全部发送

}

func (this *Beater) SendPusher() chan model.Beat {
	//
	return this.PushChan
}

func (this *Beater) Run() chan model.Beat {
	//
	this.Fetch(this.FileName)
	return this.SendPusher()
}

func (this *Beater) Stop() {

}
