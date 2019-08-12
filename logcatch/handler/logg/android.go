package logg

import (
	"encoding/json"
	"fmt"
	"logcat/logcatch/httpctx"
	"logcat/logcatch/model"
	"logcat/tools/logrus"

	"github.com/kataras/iris"
)

func AndroidLog(ctx iris.Context, logger *logrus.Logger, context *httpctx.ServerContext) {

	jsons, errs := json.Marshal(&model.InLog{
		LogLevel:   1,
		LogHeader:  "",
		LogContent: "",
		Platform:   "",
	})
	if errs != nil {
		fmt.Println(errs.Error())
	}

	context.Kafka.Log("android", string(jsons), 1)

}
