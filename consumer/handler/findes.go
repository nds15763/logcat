package handler

import (
	"encoding/json"
	"fmt"
	"logcat/consumer/constant"
	"logcat/consumer/httpctx"
	"logcat/tools/common"
	"logcat/tools/logrus"
	"os"

	"github.com/kataras/iris"
)

type Request struct {
	Cid string `json:"cid"`
}

func FindEs(ctx iris.Context, logger *logrus.Logger, session *httpctx.ServerContext) {

	clickid := new(Request)
	err := json.NewDecoder(ctx.Request().Body).Decode(&clickid)
	if err != nil {
		common.ResponseErrorCode(ctx, 100000)
		logger.Warn(fmt.Sprintf("解析消息错误: %s", err.Error()))
		json.NewEncoder(os.Stdout).Encode(clickid)
		return
	}
	_ = clickid
	// var buf bytes.Buffer

	// item := &model.ClickLog{ClickId: clickid.Cid}
	// match := &model.Match{Match: item}
	// query := &model.Query{Query: match}

	// if err := json.NewEncoder(&buf).Encode(query); err != nil {
	// 	log.Fatalf("Error encoding query: %s", err)
	// }
	// log.Println(fmt.Sprintf("buf: %s", buf.String()))

	// // Perform the search request.
	// es := session.EsManager.ESClient
	// res, err := es.Search(
	// 	es.Search.WithContext(context.Background()),
	// 	es.Search.WithIndex("clicklog"),
	// 	es.Search.WithBody(&buf),
	// 	es.Search.WithTrackTotalHits(true),
	// 	es.Search.WithPretty(),
	// )
	// if err != nil {
	// 	log.Fatalf("Error getting response: %s", err)
	// }
	// defer res.Body.Close()

	// if res.IsError() {
	// 	var e map[string]interface{}
	// 	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
	// 		log.Fatalf("Error parsing the response body: %s", err)
	// 	} else {
	// 		// Print the response status and error information.
	// 		log.Fatalf("[%s] %s: %s",
	// 			res.Status(),
	// 			e["error"].(map[string]interface{})["type"],
	// 			e["error"].(map[string]interface{})["reason"],
	// 		)
	// 	}
	// }

	// items := &model.HitItems{
	// 	Source: &model.ClickLog{},
	// }

	// hits := make([]*model.HitItems, 0)
	// subHits := &model.Hits{
	// 	Hits: append(hits, items),
	// }

	// cResponse := &model.EsResponse{
	// 	Hits: subHits,
	// }

	// if err := json.NewDecoder(res.Body).Decode(cResponse); err != nil {
	// 	log.Fatalf("Error parsing the response body: %s", err)
	// }
	// // Print the response status, number of results, and request duration.
	// log.Printf(
	// 	"[%s] %d hits; took: %dms",
	// 	res.Status(),
	// 	cResponse.Hits.Total,
	// 	cResponse.Took,
	// )
	// // Print the ID and document source for each hit.
	// Response := make([]*model.ClickLog, 0)
	// for _, hit := range cResponse.Hits.Hits {
	// 	source := hit.Source.(*model.ClickLog)
	// 	Response = append(Response, source)
	// }

	// log.Println(strings.Repeat("=", 37))

	common.ResponseSuccess(ctx, constant.CONSTANT_CODE_REQUEST_SUCCESS, nil)

}
