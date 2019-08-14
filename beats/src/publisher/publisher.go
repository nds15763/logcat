package publisher

import (
	"beats/src/crawler"
	"beats/src/model"
	"bytes"
	"mime/multipart"
	"net/http"
)

type Publisher struct {
	cra crawler.Crawler
}

func NewPublisher(cra crawler.Crawler) *Publisher {
	return &Publisher{
		cra: cra,
	}
}

func (this *Publisher) Run() {

	crawler := this.cra.DoFetch()

	for beaters := range crawler {
		for beat := range beaters {
			this.pusher(beat)
		}
	}
}

func (this *Publisher) pusher(beat model.Beat) {

	uri := beat.Ip //URL地址
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)

	//取出内容类型
	content_type := w.FormDataContentType()

	//将文件数据写入
	// pa, _ := w.CreateFormFile("file", fn)
	// pa.Write(file_data)
	w.Close()

	//开始提交
	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", content_type)
	resp, _ := http.DefaultClient.Do(req)
	//data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
}
