package handler

import (
	"encoding/json"
	"logcat/logcatch/httpctx"
	"logcat/tools/logrus"

	"adwetec.com/tools/common"
	"github.com/kataras/iris"
)

func AddLogJob(ctx iris.Context, logger *logrus.Logger, context *httpctx.ServerContext) {

	inlog := new(LogJob)
	err := json.NewDecoder(ctx.Request().Body).Decode(&inlog)
	if err != nil {
		common.ResponseErrorCode(ctx, 500)
		//context.Logger.Warn(fmt.Sprintf("%s %d: -- 解析消息错误: %s", reqid, userid, err.Error()))
		return
	}
	logjob := LogJob{
		LogLevel:   inlog.LogLevel,
		LogHeader:  inlog.LogHeader,
		LogContent: inlog.LogContent,
		Platform:   inlog.Platform,
	}
	work := Job{LogJob: logjob}
	// 任务放入任务队列channal
	JobQueue <- work

	common.ResponseSuccess(ctx, 200, "success")

}
