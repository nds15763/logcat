package constant

const (
	CONST_REPORT_RETRY_TIMES = 5
)
const (
	CONST_BCE_RESPONSE_EXECEPTION = "BecResponseException"
	CONST_BCE_RATE_LIMIT          = "RateLimit"
	CONST_BCE_EOF                 = "EOF"
	CONST_BCE_REACHED_TIMEOUT     = "Reached timeout"
)
const (
	CONST_RATIO_OVERRATE = 1.01
)
const (
	CONST_ADSTYPE_BC0054 = "BC0054" // 手机百度-列表页-Feed 小图链接
	CONST_ADSTYPE_BC0055 = "BC0055" // 手机百度-列表页-Feed 大图链接
	CONST_ADSTYPE_BC0056 = "BC0056" // 手机百度-列表页-Feed 三图链接
	CONST_ADSTYPE_BC0068 = "BC0068" // 手机百度-列表页-Feed 大图视频
	CONST_ADSTYPE_BC0133 = "BC0133" // 手机百度-详情页-Feed 随心互动
	CONST_ADSTYPE_BC0134 = "BC0134" // 手机百度-详情页-Feed 像素渐变
	CONST_ADSTYPE_BC0092 = "BC0092" // 手机百度-视频列表页-Feed-大图链接
	CONST_ADSTYPE_BC0094 = "BC0094" // 手机百度-视频列表页-Feed-大图视频
	CONST_ADSTYPE_BC0119 = "BC0119" // 手机百度-视频详情页后贴片-大图链接
	CONST_ADSTYPE_BC0130 = "BC0130" // 好看视频-列表页-大图链接
	CONST_ADSTYPE_BC0131 = "BC0131" // 好看视频-列表页-大图视频
	CONST_ADSTYPE_BK001  = "BK001"  // 百度APP静态图片闪屏-全屏样式
	CONST_ADSTYPE_BK002  = "BK002"  // 百度APP静态图片闪屏-半屏样式-安卓
	CONST_ADSTYPE_BK003  = "BK003"  // 百度APP动态视频闪屏-全屏样式
	CONST_ADSTYPE_BK004  = "BK004"  // 百度APP动态视频闪屏-半屏样式
	CONST_ADSTYPE_BK005  = "BK005"  // 百度APP摇一摇闪屏-静态图片
	CONST_ADSTYPE_BK006  = "BK006"  // 百度APP摇一摇闪屏-动态视频
	CONST_ADSTYPE_BK007  = "BK007"  // 百度APP九宫格闪屏-静态图片
	CONST_ADSTYPE_BK008  = "BK008"  // 百度APP九宫格闪屏-动态视频
	CONST_ADSTYPE_BK009  = "BK009"  // 百度地图闪屏-静态图片
	CONST_ADSTYPE_BK010  = "BK010"  // 百度地图闪屏-动态视频
	CONST_ADSTYPE_BK011  = "BK011"  // 百度贴吧闪屏-静态图片
	CONST_ADSTYPE_BK012  = "BK012"  // 百度贴吧闪屏-动态视频
	CONST_ADSTYPE_BK013  = "BK013"  // 好看视频APP静态图片-全屏样式
	CONST_ADSTYPE_BK014  = "BK014"  // 好看视频APP静态图片-半屏样式
	CONST_ADSTYPE_BK015  = "BK015"  // 好看视频APP动态视频-全屏样式
	CONST_ADSTYPE_BK016  = "BK016"  // 百度APP静态图片闪屏-半屏样式-IOS
	CONST_ADSTYPE_BK017  = "BK017"  // 百度APP摇一摇闪屏-静态图片-IOS

	CONST_ADSTYPE_10012 = "10012" //作业帮-开屏视频 id 5
	CONST_ADSTYPE_10014 = "10014" //作业帮-开屏图片 id 6
	CONST_ADSTYPE_10002 = "10002" //作业帮-信息流大图 id 7
	CONST_ADSTYPE_10003 = "10003" //作业帮-信息流小图 id 8
	CONST_ADSTYPE_10004 = "10004" //作业帮-信息流三图 id 9
	CONST_ADSTYPE_10008 = "10008" //作业帮-信息流视频 id 10
)

const (
	CONST_PDB_CREATIVE_NURL   = "http://pdbreq.adwintech.com/baidu/wn?impid={{IMPID}}"
	CONST_PDB_CREATIVE_IMPURL = "http://pdbreq.adwintech.com/baidu/in?impid={{IMPID}}"
	CONST_PDB_CREATIVE_EXPURL = "http://pdbreq.adwintech.com/baidu/cn?impid={{IMPID}}"

	//作业帮
	CONST_ZYB_CREATIVE_NURL   = "https://pdbreq.adwetec.com/zybang/wn?impid={{IMPID}}"
	CONST_ZYB_CREATIVE_IMPURL = "https://pdbreq.adwetec.com/zybang/in?impid={{IMPID}}"
	CONST_ZYB_CREATIVE_EXPURL = "https://pdbreq.adwetec.com/zybang/cn?impid={{IMPID}}"

	//开屏
	CONST_BDKPAD_CREATIVE_NURL   = "http://pdbreq.adwintech.com/bdkpad/wn?impid={{IMPID}}"
	CONST_BDKPAD_CREATIVE_IMPURL = "http://pdbreq.adwintech.com/bdkpad/in?impid={{IMPID}}"
	CONST_BDKPAD_CREATIVE_EXPURL = "http://pdbreq.adwintech.com/bdkpad/cn?impid={{IMPID}}"
)

const (
	CONST_PDB_ADVERTISER_NAME_REPE = 40001
	CONST_PDB_ADVERTISER_URL_REPE  = 40021
	CONST_PDB_ACCOUNT_NAME_REPE    = 40002
	CONST_PDB_CAMPAIGN_NAME_REPE   = 40003
	CONST_PDB_CAMPAIGN_DEALID_REPE = 40023
	CONST_PDB_ADGROUP_NAME_REPE    = 40004
	CONST_PDB_ADGROUP_RATIO_ERROR  = 40006
	CONST_PDB_CREATIVE_NAME_REPE   = 40005

	//CONST_PDB_ADVERTISER_NAME_EMPTY = 40011
	//CONST_PDB_ACCOUNT_NAME_EMPTY = 40012
	//CONST_PDB_CAMPAIGN_NAME_EMPTY = 40013
	//CONST_PDB_ADGROUP_NAME_EMPTY = 40014
	//CONST_PDB_CREATIVE_NAME_EMPTY = 40015

	CONST_ZYB_CREATIVE_TITLE_ERROR        = 50001
	CONST_ZYB_CREATIVE_DESC_ERROR         = 50002
	CONST_ZYB_CREATIVE_NUMS_ERROR         = 50007
	CONST_ZYB_CREATIVE_URL_ERROR          = 50008
	CONST_ZYB_CREATIVE_TITLENUMS_ERROR    = 50009
	CONST_ZYB_MATERIEL_VIDEO_FORMAT_ERROR = 50003
	CONST_ZYB_MATERIEL_VIDEO_SIZE_ERROR   = 50004
	CONST_ZYB_MATERIEL_VIDEO_TIME_ERROR   = 50005
	CONST_ZYB_MATERIEL_VIDEO_CODE_ERROR   = 50006
)

const (
	CONST_PDB_CREATIVE_LOCALSTATUS_ENABLE  = 1
	CONST_PDB_CREATIVE_LOCALSTATUS_DISABLE = 0

	CONST_PDB_ADVGROUP_LOCALSTATUS_ENABLE  = 1
	CONST_PDB_ADVGROUP_LOCALSTATUS_DISABLE = 0
)

const (
	CONSTANT_FEED_LENGTH_CAMPAIGN = 100   // 表示每次获取计划的长度(信息流)
	CONSTANT_FEED_LENGTH_ADGROUP  = 10000 // 表示每次获取单元的长度(信息流)
	CONSTANT_FEED_LENGTH_CREATIVE = 10000 // 表示每次获取创意的长度(信息流)
)

const (
	CONSTANT_PDB_AUDIT_STATUS_INIT       = 0 // 首次插入需要审核
	CONSTANT_PDB_AUDIT_STATUS_PROCESSING = 1 // 审核中
	CONSTANT_PDB_AUDIT_STATUS_SUCCESS    = 2 // 审核成功
	CONSTANT_PDB_AUDIT_STATUS_FAILED     = 3 // 审核失败
	CONSTANT_PDB_AUDIT_STATUS_PAUSE      = 4 // 投放结束
)
const (
	CONST_PDB_MATERIAL_STATUS_DELETE  = "0" // 已删除
	CONST_PDB_MATERIAL_STATUS_DRAFT   = "1" // 草稿
	CONST_PDB_MATERIAL_STATUS_DELIVER = "2" // 投放中
	CONST_PDB_MATERIAL_STATUS_OFFLINE = "3" // 已下线
	CONST_PDB_MATERIAL_STATUS_PAUSE   = "4" // 暂停投放
	CONST_PDB_MATERIAL_STATUS_COMMIT  = "5" // 待提交审核
	CONST_PDB_MATERIAL_STATUS_REVIEW  = "6" // 审核中
	CONST_PDB_MATERIAL_STATUS_REFUSE  = "7" // 审核拒绝
	CONST_PDB_MATERIAL_STATUS_EXPIRE  = "8" // 已过期
	CONST_PDB_MATERIAL_STATUS_READY   = "9" // 准备完毕未开始投放
)
const (
	BCE_ACCESS_KEY = "6e27935d3905430bbc2f224636f25ed4" // access key
	BCE_SECRET_KEY = "982b0fa07d99473d8d9b5c46ab05ba31" // secret key
	BCE_USER       = "111cd79b4ed2441898433f6c8e1ca367" // 百度云账户标识

	SHOUBAI_URL          = "http://api.data.baidu.com/channel/api/checkimei"
	SHOUBAI_CHANNEL_NAME = "zhanmei001"
	SHOUBAI_SECRET_KEY   = "hello_channel"

	JWT_SECRETKEY = "welcome to wetec.com"

	FEED_API_AGENT_ACCOUNT  = "原生-致维8171823"
	FEED_API_AGENT_PASSWORD = "Zwkj161209"

	SEM_API_AGENT_ACCOUNT  = "baidu-致维科技B18KA4052"
	SEM_API_AGENT_PASSWORD = "Wetec1234"

	TOKEN_EXPIRE_TIME   = 5
	TOKEN_EXPIRE_WINDOW = 5

	ALI_HUICHUAN_TOKEN = "e9629aa6-1535-4301-9be4-f22777e0e9c4"
	ALI_SHENMA_TOKEN   = "14e1bc26-f001-4384-82f4-afe5ec643a58"

	//腾讯
	TX_CLIENT_ID          = "1107900179"
	TX_REDIRECT_URI       = "http://feedsapi.adwintech.com/tencent"
	TX_CLIENT_SECRET      = "dFIITQ7BaA0RRwE9"
	TX_REFRESH_TOKEN      = "refresh_token"
	TX_REFRESH_TOKEN_INFO = "8f4b18e4843c41c2dab97ab8c15dfb22"
	TX_ACCESS_TOKEN       = "txaccesstoken"

	//头条
	TT_APP_ID             = 1616992255041559
	TT_SECRET             = "b8a091efe63f4a11bdfa0e05ae89e04401c21156"
	TT_GRANT_TYPE_AUTH    = "auth_code"
	TT_GRANT_TYPE_REFRESH = "refresh_token"
	TT_ACCESS_TOKEN       = "ttaccesstoken"
)

var AES_SECRET_KEY = []byte{2, 0, 1, 5, 1, 2, 0, 9, 2, 0, 1, 5, 1, 2, 0, 9}

const (
	CONSTANT_ROLE_SYSTEM     int64 = 1
	CONSTANT_ROLE_MANAGER          = 2
	CONSTANT_ROLE_OPERATE          = 3
	CONSTANT_ROLE_ADVERTISER       = 4
	CONSTANT_ROLE_SELLER           = 5
)

const (
	TIME_UNIT_ALLS = "all"
	TIME_UNIT_HOUR = "hour"
	TIME_UNIT_DATE = "day"
	TIME_UNIT_WEEK = "week"
	TIME_UNIT_MONT = "month"
)
const (
	KPI_TYPE_IMP         = "impression"
	KPI_TYPE_CLK         = "click"
	KPI_TYPE_COS         = "cost"
	KPI_TYPE_CPC         = "cpc"
	KPI_TYPE_CPM         = "cpm"
	KPI_TYPE_CTR         = "ctr"
	KPI_TYPE_PV          = "pv"
	KPI_TYPE_UV          = "uv"
	KPI_TYPE_ACTION1     = "action1"            //独立号码
	KPI_TYPE_CONSUMPTION = "actual_consumption" //实际消费 [账户消费/1.38*(1-11%+2%)]
	KPI_TYPE_ORDERCOST   = "order_cost"         //订单成本
	KPI_TYPE_ORDERUV     = "order_uv"           //订单量/uv
	KPI_TYPE_ORDERCLICK  = "order_click"        //订单量/点击
	KPI_TYPE_UVCLICK     = "uv_click"           //uv/点击
)
const (
	CONSTANT_LEVEL_STRING_ADVERTISER = "advertiser"
	CONSTANT_LEVEL_STRING_MEDIA      = "media"
	CONSTANT_LEVEL_STRING_ACCOUNT    = "account"
	CONSTANT_LEVEL_STRING_CAMPAIGN   = "campaign"
	CONSTANT_LEVEL_STRING_ADGROUP    = "adgroup"
	CONSTANT_LEVEL_STRING_CREATIVE   = "creative"
	CONSTANT_LEVEL_STRING_KEYWORD    = "keyword"
)

// 抓取数据错误码
const (
	CONST_PWD_WRONG       int = 99914
	CONST_REQUST_FREQUENT int = 9011111
)

const (
	CONST_LEVEL_ACCOUNT int32 = iota
	CONST_LEVEL_CAMPAIGN
	CONST_LEVEL_ADGROUP
	CONST_LEVEL_CREATIVE
	CONST_LEVEL_KEYWORD
)
const ( // 统计时间单位
	CONST_UNIT_REPORT_DAILY     int32 = 5 // 分日
	CONST_UNIT_REPORT_WEEKLY          = 4 // 分周
	CONST_UNIT_REPORT_MONTHLY         = 3 // 分月
	CONST_UNIT_REPORT_YEARLY          = 1 // 分年
	CONST_UNIT_REPORT_HOURLY          = 7 // 分时
	CONST_UNIT_REPORT_TIMERANGE       = 8 // 请求时间段汇总
)
const ( // 统计范围
	CONST_STAT_RANGE_ACCOUNT   int32 = 2  // 账户范围
	CONST_STAT_RANGE_CAMPAIGN        = 3  // 计划范围
	CONST_STAT_RANGE_ADGROUP         = 5  // 单元范围
	CONST_STAT_RANGE_CREATIVE        = 7  // 创意范围
	CONST_STAT_RANGE_KEYWORDID       = 11 // 关键词范围
	CONST_STAT_RANGE_WORDIS          = 6  // 关键词范围
)

const ( // 指定返回的数据层级
	CONST_DETAILS_LEVEL_ACCOUNT                int32 = 2  // 账户粒度
	CONST_DETAILS_LEVEL_CAMPAIGN                     = 3  // 计划粒度
	CONST_DETAILS_LEVEL_ADGROUP                      = 5  // 单元粒度
	CONST_DETAILS_LEVEL_CREATIVE                     = 7  // 创意粒度
	CONST_DETAILS_LEVEL_KEYWORD                      = 11 // 关键词粒度--KEYWORDID
	CONST_DETAILS_LEVEL_KEYWORDID_AND_CREATIVE       = 12 // 关键词+创意粒度--KEYWORDID
	CONST_DETAILS_LEVEL_WORDID                       = 6  // 关键词粒度--WORDID
)
const (
	CONST_SUBJECT_ALL     int32 = 0
	CONST_SUBJECT_WEBSITE       = 1
	CONST_SUBJECT_IOS           = 2
	CONST_SUBJECT_ANDROID       = 3
)
const (
	CONST_PRODUCT_TYPE_ALL    int32 = 0 // 全部
	CONST_PRODUCT_TYPE_LIST         = 1 // 列表页
	CONST_PRODUCT_TYPE_DETAIL       = 2 // 详情页
)
const (
	CONST_MATERIAL_STYLE_ALL    int32 = 0 // 全部
	CONST_MATERIAL_STYLE_SINGLE       = 1 // 单图
	CONST_MATERIAL_STYLE_TRIPLE       = 2 // 三图
)
const ( // 用户状态
	CONST_USER_STAT_REFUSEED  int32 = 1  // 被拒绝
	CONST_USER_STAT_WORKING         = 2  // 生效
	CONST_USER_STAT_NOBALANCE       = 3  // 余额为零
	CONST_USER_STAT_INREVIEW        = 6  // 审核中
	CONST_USER_STAT_DISABLE         = 7  // 被禁用
	CONST_USER_STAT_NOBUDGET        = 11 // 预算不足
)
const ( // 资金包
	CONST_BALANCE_PACKAGE_NATIVE  int32 = 0 // 原生资金包
	CONST_BALANCE_PACKAGE_FENGCAO       = 1 // 凤巢资金包
	CONST_BALANCE_PACKAGE_AGENT         = 3 // 代理商原生资金包
)
const ( // 是否开通FEED产品线
	CONST_FEED_PRODUCT_OPEN     int32 = 1 // 已开通产品线
	CONST_FEED_PRODUCT_NEEDOPEN       = 2 // 待开通产品线
	CONST_FEED_PRODUCT_DISABLE        = 3 // 不允许开通产品线
)
const ( // 流量禁推属性
	CONST_VALID_FLOWS_SHOUBAI int32 = 1 // 手机百度信息流
	CONST_VALID_FLOWS_TIEBA         = 2 // 百度贴吧信息流
)

const (
	CONST_DEPTH_ACCOUNT  int32 = 1
	CONST_DEPTH_CAMPAIGN       = 2
	CONST_DEPTH_ADGROUP        = 3
	CONST_DEPTH_CREATIVE       = 4
	CONST_DEPTH_KEYWORD        = 5
)

const ( // 致维定义的媒介账户状态
	MEDIA_ACCOUNT_STATUS_NORMAL   int32 = 1 // 正常
	MEDIA_ACCOUNT_STATUS_ABNORMAL int32 = 2 // 异常
	MEDIA_ACCOUNT_STATUS_STOP     int32 = 3 // 停用
)

type CONSTANT_CODE int32 // 对前端响应的状态码

const (
	CONSTANT_CODE_PARAMETER_ERROR CONSTANT_CODE = 10000

	CONSTANT_CODE_NAMEPASS_ERROR        = 10001
	CONSTANT_CODE_SERVER_ERROR          = 10002
	CONSTANT_CODE_TOKEN_EXPIRS          = 10003
	CONSTANT_CODE_NEED_LOGIN            = 10004
	CONSTANT_CODE_USER_INVALIDE         = 10005
	CONSTANT_CODE_USER_ALREADY_EXIST    = 10006
	CONSTANT_CODE_ROLE_LIMIT            = 10007
	CONSTANT_CODE_ACCOUNT_ALREADY_EXIST = 10008

	CONSTANT_CODE_REQUEST_SUCCESS = 20000
	CONSTANT_CODE_LOGIN_SUCCESS   = 20001

	CONSTANT_CODE_OPERATE_DELETE_FAILED   = 30000
	CONSTANT_CODE_OPERATE_UPDATE_FAILED   = 30001
	CONSTANT_CODE_OPERATE_CREATE_FAILED   = 30002
	CONSTANT_CODE_OPERATE_DELETE_SUCCESS  = 30003
	CONSTANT_CODE_OPERATE_UPDATE_SUCCESS  = 30004
	CONSTANT_CODE_OPERATE_CREATE_SUCCESS  = 30005
	CONSTANT_CODE_OPERATE_INTEGRAL_FAILED = 30006
)

type CONSTANT_ACCOUNT_STATUS int32

const (
	CONSTANT_ACCOUNT_STATUS_UNKOWN    CONSTANT_ACCOUNT_STATUS = 0  // 未设置获取该属性
	CONSTANT_ACCOUNT_STATUS_REFUSE                            = 1  // 被拒绝
	CONSTANT_ACCOUNT_STATUS_ENABLE                            = 2  // 生效
	CONSTANT_ACCOUNT_STATUS_NOBALANCE                         = 3  // 余额为零
	CONSTANT_ACCOUNT_STATUS_REVIEW                            = 6  // 审核中
	CONSTANT_ACCOUNT_STATUS_DISABLE                           = 7  // 被禁用
	CONSTANT_ACCOUNT_STATUS_NOBUDGET                          = 11 // 预算不足
)

type CONSTANT_CAMPAIGN_STATUS int32

const (
	CONSTANT_CAMPAIGN_STATUS_ENABLE               CONSTANT_CAMPAIGN_STATUS = 0  // 有效
	CONSTANT_CAMPAIGN_STATUS_PAUSEZONE                                     = 1  // 处于暂停时段
	CONSTANT_CAMPAIGN_STATUS_PAUSE                                         = 2  // 计划暂停
	CONSTANT_CAMPAIGN_STATUS_CAMPAIGN_NO_BUDGET                            = 3  // 计划预算不足
	CONSTANT_CAMPAIGN_STATUS_ACCOUNT_NO_BUDGET                             = 11 // 账户预算不足
	CONSTANT_CAMPAIGN_STATUS_ACCOUNT_ZERO_BALANCE                          = 20 // 账户余额为 0
	CONSTANT_CAMPAIGN_STATUS_APP_OFFLINE                                   = 24 // APP已下线
)

type CONSTANT_ADGROUP_STATUS int32

const (
	CONSTANT_ADGROUP_STATUS_ENABLE         CONSTANT_ADGROUP_STATUS = 0 // 有效
	CONSTANT_ADGROUP_STATUS_PAUSE                                  = 1 // 暂停推广
	CONSTANT_ADGROUP_STATUS_CAMPAIGN_PAUSE                         = 2 // 推广计划暂停推广
)

type CONSTANT_CREATIVE_STATUS int32

const (
	CONSTANT_CREATIV_STATUS_ENABLE        CONSTANT_CREATIVE_STATUS = 0 // 有效
	CONSTANT_CREATIV_STATUS_PAUSE                                  = 1 // 暂停推广
	CONSTANT_CREATIV_STATUS_REVIEW                                 = 2 // 审核中
	CONSTANT_CREATIV_STATUS_REVIEW_FAILED                          = 3 // 审核未通过
	CONSTANT_CREATIV_STATUS_DISABLE                                = 4 // 无效
)

const (
	CONSTANT_MODE_PRODUCT = "product"
	CONSTANT_MODE_DEVELOP = "develop"
)
const (
	CONSTANT_OPERATE_DELETE = "delete"
	CONSTANT_OPERATE_UPDATE = "update"
	CONSTANT_OPERATE_CREATE = "add"
)

const ( // 百度SEM -- 实时数据类型
	CONSTANT_REPORT_TYPE_ACCOUNT            int32 = 2  // 账户范围
	CONSTANT_REPORT_TYPE_CAMPAIGN                 = 10 // 计划范围
	CONSTANT_REPORT_TYPE_ADGROUP                  = 11 // 单元范围
	CONSTANT_REPORT_TYPE_KEYWORDID                = 14 // 关键词--keywordid
	CONSTANT_REPORT_TYPE_CREATIVE                 = 12 // 创意范围
	CONSTANT_REPORT_TYPE_PAIR                     = 15 // 配对
	CONSTANT_REPORT_TYPE_QUERY                    = 6  // 搜索词报告
	CONSTANT_REPORT_TYPE_REGION1                  = 3  // 地域
	CONSTANT_REPORT_TYPE_WORDID                   = 9  // 关键词--wordid
	CONSTANT_REPORT_TYPE_REGION2                  = 5  // 二级地域报告
	CONSTANT_REPORT_TYPE_PRESENTSTION             = 21 // 蹊径报告
	CONSTANT_REPORT_TYPE_ORDER_PRESENTSTION       = 38 // 历史数据排名报告
)

const ( // 百度信息流 -- 实时数据类型
	CONST_REPORT_TYPE_ACCOUNT  int32 = 700 // 账户范围
	CONST_REPORT_TYPE_CAMPAIGN       = 701 // 计划范围
	CONST_REPORT_TYPE_ADGROUP        = 702 // 单元范围
	CONST_REPORT_TYPE_CREATIVE       = 703 // 创意范围
)

const (
	CONST_REDIS_EXPIREAT_TIME = 5 * 60
)

//三江媒介产品线
const (
	CONST_MEDIA_BAIDUHAO123     = 1
	CONST_MEDIA_BAIDUSEM        = 2
	CONST_MEDIA_BAIDUFENFA      = 3
	CONST_MEDIA_BAIDUFEED       = 4
	CONST_MEDIA_BAIDUZHANSHILEI = 5
	CONST_MEDIA_ALIFEED         = 6
	CONST_MEDIA_ALISEM          = 7
	CONST_MEDIA_ALIFENFA        = 8
	CONST_MEDIA_KUAISHOU        = 9
)
const (
	CONST_ALERT_DONE    = 1
	CONST_ALERT_WAITING = 2
	CONST_ALERT_WARNING = 3
)
