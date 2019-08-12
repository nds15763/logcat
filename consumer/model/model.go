package model

type BaseLog struct {
	LogLevel   int64  `json:"log_level"`   //日志级别，error，warning，info
	LogHeader  string `json:"log_header"`  //标题
	LogContent string `json:"log_content"` //内容
	Platform   string `json:"platform"`
}
