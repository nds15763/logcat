日志收集系统
------------------------------------------------------------------------

2019-08-11 添加生产部分，采用线程队列控制并发数量

2019-08-12 修改消费部分，添加kafaka消费线程池

2019-08-13 生产部分修改

2019-08-14 添加服务端日志文件拉取功能，根据配置项查找文件内容，并转化为请求推送给logcatch端

2019-08-17 修改两端链接方式，修改beats端publisher方法

------------------------------------------------------------------------
可实现部分：

参照https://tech.meituan.com/2018/02/11/logan.html

日志在网络传输需要注意安全性，AES加密

需要使用缓存缓解IO压力缓存可以使用mmap代替redis

同时采用流式的压缩和加密避免了集中压缩加密可能产生的CPU峰值，所以CPU平均使用率会降低（？这啥）

·CAT端到端日志
·埋点日志
·用户行为日志
·代码级日志
·网络内部日志
·Push日志

logcatch:多元化接收日志，需要解析stream 定义格式
{"log_level":"","log_header":"","log_content":"","platform":"","thread_id":""}




1.传入采用流式传输，相对于http帧节约资源
https://segmentfault.com/a/1190000013339403

2.日志聚合
现代日志管理和分析解决方案包括以下主要功能：

聚合 - 从多个数据源收集和发送日志的能力。
处理 - 将日志消息转换为有意义的数据以便于分析的能力。
存储 - 能够在延长的时间段内存储数据，以便进行监控，趋势分析和安全用例。
分析 - 通过查询数据并在其上创建可视化和仪表板来分析数据的能力。
https://godoc.org/github.com/suzuki-shunsuke/go-graylog


3.elk整体方案

