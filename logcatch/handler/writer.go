package handler

// import (
// 	"logcat/logcatch/httpctx"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"regexp"
// 	"time"

// 	"github.com/kataras/iris"
// )

// func WriteHandler(ctx iris.Context, logger *logrus.Logger, context *httpctx.ServerContext) {
// 	s := NewSpider("https://www.d1xz.net/test/gexing/")
// 	s.Parse(ctx, logger, context)
// }

// //定义新的数据类型
// type Spider struct {
// 	url    string
// 	header map[string]string
// }
// type Item struct {
// 	link    string `json:"link"`
// 	title   string
// 	content string
// }

// func NewSpider(url string) *Spider {
// 	return &Spider{
// 		url: url,
// 		header: map[string]string{
// 			"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
// 			"accept-encoding":           "gzip, deflate, br",
// 			"accept-language":           "zh-CN,zh;q=0.9",
// 			"referer":                   "https://www.d1xz.net/",
// 			"upgrade-insecure-requests": "1",
// 			"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.90 Safari/537.36",
// 		},
// 	}
// }

// //定义 Spider get的方法
// func (s *Spider) get_html_header() string {

// 	resp, err := http.Get(s.url)
// 	if err != nil {
// 		fmt.Println("err", err.Error())
// 	}
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("err", err.Error())
// 	}
// 	return string(body)
// }

// func (s *Spider) Parse(ctx iris.Context, logger *logrus.Logger, context *httpctx.ServerContext) {

// 	for i := 0; i < 1; i++ {

// 		spider := &Spider{s.url, s.header}
// 		html := spider.get_html_header()

// 		//各个条目
// 		pattern1 := `<li><a target="_blank"(.*?)</p></div></li>`
// 		rp1 := regexp.MustCompile(pattern1)
// 		find_txt1 := rp1.FindAllStringSubmatch(html, -1)

// 		for k, _ := range find_txt1 {
// 			//各个条目
// 			pattern2 := `href="(.*?)"`
// 			rp2 := regexp.MustCompile(pattern2)
// 			item_link := rp2.FindAllStringSubmatch(find_txt1[k][0], -1)
// 			_ = item_link[0][1]
// 			//各个条目
// 			pattern3 := `class="t">(.*?)</a>`
// 			rp3 := regexp.MustCompile(pattern3)
// 			item_title := rp3.FindAllStringSubmatch(find_txt1[k][0], -1)
// 			_ = item_title[0][1]
// 			//各个简介
// 			pattern4 := `class="txt"><p>(.*?)</p>`
// 			rp4 := regexp.MustCompile(pattern4)
// 			item_context := rp4.FindAllStringSubmatch(find_txt1[k][0], -1)
// 			_ = item_context[0][1]

// 			jsons, errs := json.Marshal(&Item{
// 				link:    item_link[0][1],
// 				title:   item_title[0][1],
// 				content: item_context[0][1],
// 			})
// 			if errs != nil {
// 				fmt.Println(errs.Error())
// 			}

// 			context.Kafka.Log("gexing", string(jsons), 1)

// 			time.Sleep(10)
// 		}

// 		fmt.Println("find_txt1: ", find_txt1)

// 	}
// }
