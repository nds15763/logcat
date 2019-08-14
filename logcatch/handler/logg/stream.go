package logg

import (
	"bufio"
	"io"
	"log"
	"logcat/logcatch/httpctx"
	"logcat/tools/logrus"

	"github.com/kataras/iris"
)

func StreamLog(ctx iris.Context, logger *logrus.Logger, context *httpctx.ServerContext) {

	//

	reader := bufio.NewReader(ctx.Request.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if io.EOF == err {
			break
		}

		log.Println(string(line))
	}
}
