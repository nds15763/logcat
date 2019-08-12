package common

import (
	"fmt"
	"reflect"
	"strings"
	"net/http"
	"encoding/json"

	"adwetec.com/tools/utils"
	"adwetec.com/tools/constant"

	"github.com/kataras/iris"
	"github.com/golang/protobuf/proto"
)

var AccountStatusMap = make(map[constant.CONSTANT_ACCOUNT_STATUS]string)
var CampaignStatusMap = make(map[constant.CONSTANT_CAMPAIGN_STATUS]string)
var AdgroupStatusMap = make(map[constant.CONSTANT_ADGROUP_STATUS]string)
var CreativeStatusMap = make(map[constant.CONSTANT_CREATIVE_STATUS]string)

var CodeMsgMap = make(map[constant.CONSTANT_CODE]string)

func init() {

	// 前端响应状态码
	// 10000-19999
	CodeMsgMap[constant.CONSTANT_CODE_PARAMETER_ERROR] = "请求参数格式错误！"
	CodeMsgMap[constant.CONSTANT_CODE_NAMEPASS_ERROR] = "用户名密码不匹配！"
	CodeMsgMap[constant.CONSTANT_CODE_SERVER_ERROR] = "系统内部错误！"
	CodeMsgMap[constant.CONSTANT_CODE_TOKEN_EXPIRS] = "超时重新登录！"
	CodeMsgMap[constant.CONSTANT_CODE_NEED_LOGIN] = "请先登录！"
	CodeMsgMap[constant.CONSTANT_CODE_USER_INVALIDE] = "此账户无效!"
	CodeMsgMap[constant.CONSTANT_CODE_USER_ALREADY_EXIST] = "此用户已经存在!"
	CodeMsgMap[constant.CONSTANT_CODE_ROLE_LIMIT] = "角色权限限制!"
	CodeMsgMap[constant.CONSTANT_CODE_ACCOUNT_ALREADY_EXIST] = "此账户已经存在!"

	// 20000-29999
	CodeMsgMap[constant.CONSTANT_CODE_REQUEST_SUCCESS] = "请求成功！"
	CodeMsgMap[constant.CONSTANT_CODE_LOGIN_SUCCESS] = "登录成功！"

	// 30000-29999
	CodeMsgMap[constant.CONSTANT_CODE_OPERATE_DELETE_FAILED] = "删除失败！"
	CodeMsgMap[constant.CONSTANT_CODE_OPERATE_UPDATE_FAILED] = "更新失败！"
	CodeMsgMap[constant.CONSTANT_CODE_OPERATE_CREATE_FAILED] = "添加失败！"
	CodeMsgMap[constant.CONSTANT_CODE_OPERATE_DELETE_SUCCESS] = "删除成功！"
	CodeMsgMap[constant.CONSTANT_CODE_OPERATE_UPDATE_SUCCESS] = "更新成功！"
	CodeMsgMap[constant.CONSTANT_CODE_OPERATE_CREATE_SUCCESS] = "添加成功！"
	CodeMsgMap[constant.CONSTANT_CODE_OPERATE_INTEGRAL_FAILED] = "用户操作次数不足"

	// 账号当前状态
	AccountStatusMap[constant.CONSTANT_ACCOUNT_STATUS_UNKOWN] = "未设置获取该属性"
	AccountStatusMap[constant.CONSTANT_ACCOUNT_STATUS_REFUSE] = "被拒绝"
	AccountStatusMap[constant.CONSTANT_ACCOUNT_STATUS_ENABLE] = "生效"
	AccountStatusMap[constant.CONSTANT_ACCOUNT_STATUS_NOBALANCE] = "余额为零"
	AccountStatusMap[constant.CONSTANT_ACCOUNT_STATUS_REVIEW] = "审核中"
	AccountStatusMap[constant.CONSTANT_ACCOUNT_STATUS_DISABLE] = "被禁用"
	AccountStatusMap[constant.CONSTANT_ACCOUNT_STATUS_NOBUDGET] = "预算不足"

	CampaignStatusMap[constant.CONSTANT_CAMPAIGN_STATUS_ENABLE] = "有效"
	CampaignStatusMap[constant.CONSTANT_CAMPAIGN_STATUS_PAUSEZONE] = "处于暂停时段"
	CampaignStatusMap[constant.CONSTANT_CAMPAIGN_STATUS_PAUSE] = "计划暂停"
	CampaignStatusMap[constant.CONSTANT_CAMPAIGN_STATUS_CAMPAIGN_NO_BUDGET] = "计划预算不足"
	CampaignStatusMap[constant.CONSTANT_CAMPAIGN_STATUS_ACCOUNT_NO_BUDGET] = "账户预算不足"
	CampaignStatusMap[constant.CONSTANT_CAMPAIGN_STATUS_ACCOUNT_ZERO_BALANCE] = "账户余额为0"
	CampaignStatusMap[constant.CONSTANT_CAMPAIGN_STATUS_APP_OFFLINE] = "APP已下线"

	AdgroupStatusMap[constant.CONSTANT_ADGROUP_STATUS_ENABLE] = "有效"
	AdgroupStatusMap[constant.CONSTANT_ADGROUP_STATUS_PAUSE] = "暂停推广"
	AdgroupStatusMap[constant.CONSTANT_ADGROUP_STATUS_CAMPAIGN_PAUSE] = "推广计划暂停推广"

	CreativeStatusMap[constant.CONSTANT_CREATIV_STATUS_ENABLE] = "有效"
	CreativeStatusMap[constant.CONSTANT_CREATIV_STATUS_PAUSE] = "暂停推广"
	CreativeStatusMap[constant.CONSTANT_CREATIV_STATUS_REVIEW] = "审核中"
	CreativeStatusMap[constant.CONSTANT_CREATIV_STATUS_REVIEW_FAILED] = "审核未通过"
	CreativeStatusMap[constant.CONSTANT_CREATIV_STATUS_DISABLE] = "无效"

	CodeMsgMap[constant.CONST_PDB_ADVERTISER_NAME_REPE] = "广告主名称已存在"
	CodeMsgMap[constant.CONST_PDB_ADVERTISER_URL_REPE] = "网址URL已存在"
	CodeMsgMap[constant.CONST_PDB_ACCOUNT_NAME_REPE] = "推广账户名已存在"
	CodeMsgMap[constant.CONST_PDB_CAMPAIGN_NAME_REPE] = "推广计划名称已存在"
	CodeMsgMap[constant.CONST_PDB_CAMPAIGN_DEALID_REPE] = "交易ID已存在"
	CodeMsgMap[constant.CONST_PDB_ADGROUP_NAME_REPE] = "推广单元名称已存在"
	CodeMsgMap[constant.CONST_PDB_CREATIVE_NAME_REPE] = "创意名称已存在"
	CodeMsgMap[constant.CONST_PDB_ADGROUP_RATIO_ERROR] = "采购比例错误"
	CodeMsgMap[constant.CONST_ZYB_CREATIVE_TITLE_ERROR] = "标题错误"
	CodeMsgMap[constant.CONST_ZYB_CREATIVE_DESC_ERROR] = "描述长度错误"
	CodeMsgMap[constant.CONST_ZYB_MATERIEL_VIDEO_CODE_ERROR] = "视频编码错误"
	CodeMsgMap[constant.CONST_ZYB_MATERIEL_VIDEO_FORMAT_ERROR] = "视频尺寸错误"
	CodeMsgMap[constant.CONST_ZYB_MATERIEL_VIDEO_SIZE_ERROR] = "视频大小错误"
	CodeMsgMap[constant.CONST_ZYB_MATERIEL_VIDEO_TIME_ERROR] = "视频时长错误"
	CodeMsgMap[constant.CONST_ZYB_CREATIVE_NUMS_ERROR] = "素材数量错误"
	CodeMsgMap[constant.CONST_ZYB_CREATIVE_URL_ERROR] = "创意URL包含中文"
	CodeMsgMap[constant.CONST_ZYB_CREATIVE_TITLENUMS_ERROR] = "创意标题或描述字数超出"
}

var HourShortName = [...]string{
	"00:00",
	"01:00",
	"02:00",
	"03:00",
	"04:00",
	"05:00",
	"06:00",
	"07:00",
	"08:00",
	"09:00",
	"10:00",
	"11:00",
	"12:00",
	"13:00",
	"14:00",
	"15:00",
	"16:00",
	"17:00",
	"18:00",
	"19:00",
	"20:00",
	"21:00",
	"22:00",
	"23:00",
}
var WeekShortName = [...]string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}
var MonthShortName = [...]string{
	"Dec",
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
}

// ******************************************************************************
// 对前端的统一响应格式
type ServerResponseData struct {
	Body *ServerResponseBody `json:"body"`
}
type ServerResponseBody struct {
	Header *ServerResponseHeader `json:"header"`
	Entity *ServerResponseEntity `json:"entity"`
}
type ServerResponseHeader struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
type ServerResponseEntity struct {
	Data interface{} `json:"data"`
}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(json)
}
func ProtoResponse(response proto.Message, w http.ResponseWriter) {

	proto, err := proto.Marshal(response)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Connection", "keep-alive")

	w.Write(proto)
}
func ResponseErrorCode(ctx iris.Context, code constant.CONSTANT_CODE) {

	response := &ServerResponseData{
		Body: &ServerResponseBody{
			Header: &ServerResponseHeader{
				Code: int32(code),
				Msg:  CodeMsgMap[code],
			},
			Entity: &ServerResponseEntity{
				Data: "",
			},
		},
	}

	JsonResponse(response, ctx.ResponseWriter())
}
func ResponseErrorMsg(ctx iris.Context, code constant.CONSTANT_CODE, msg string) {

	response := &ServerResponseData{
		Body: &ServerResponseBody{
			Header: &ServerResponseHeader{
				Code: int32(code),
				Msg:  msg,
			},
			Entity: &ServerResponseEntity{
				Data: "",
			},
		},
	}

	JsonResponse(response, ctx.ResponseWriter())
}

func ResponseSuccess(ctx iris.Context, code constant.CONSTANT_CODE, data interface{}) {

	response := &ServerResponseData{
		Body: &ServerResponseBody{
			Header: &ServerResponseHeader{
				Code: int32(code),
				Msg:  CodeMsgMap[code],
			},
			Entity: &ServerResponseEntity{
				Data: data,
			},
		},
	}

	JsonResponse(response, ctx.ResponseWriter())
}

// ******************************************************************************
//type ServerHttpContext struct {
//	// 客户端请求的上下文数据
//	Mode string
//
//	Logger    *logrus.Logger
//	Daomgr    *dao.DaoManager
//	PdbDaomgr *dao.DaoManager
//	//BidurDaomgr    *dao.DaoManager
//	CrawlersDaomgr *dao.DaoManager
//	//Redismgr       *redis.RedisManager
//	Token      *jwt.Token
//	Tokenmutex *sync.Mutex
//
//	AdwetecUserMaps  map[int64]*model.AdwetecUserMap //待删除
//	AdwetecUsers     map[int64]*model.AdwetecUser
//	AdwetecResources map[int64]*model.AdwetecResource
//
//	AdwetecRoleRuleMap      map[int64][]int64                        //所有角色的规则
//	AdwetecUserDataMap      map[int64]*model.AdwetecUserDataMap      //用户拥有账户数据
//	AdwetecUserSalesDataMap map[int64]*model.AdwetecUserSalesDataMap //销售类用户分配广告主
//
//	AdwetecRoles map[int64]*model.AdwetecRole
//	//AdwetecIpcode      map[int]*model.AdwetecIpcode
//	AdwetecMedias      map[int64]*model.AdwetecMedia
//	AdwetecAdvertisers map[int64]*model.AdwetecAdvertiser
//	AdwetecAccounts    map[int64]*model.AdwetecAccount
//	AdwetecAccountIds  map[int64]*model.AdwetecAccount // 通过ID索引
//	AdwetecEntities    map[int64]*model.AdwetecEntity
//
//	PdbAdwetecAdvertiserMaps  map[int64]*model.AdwetecPdbAdvertiser
//	PdbAdwetecAccountMaps     map[int64]*model.AdwetecPdbAccount
//	PdbAdwetecCampaignMaps    map[int64]*model.AdwetecPdbCampaign
//	PdbAdwetecAdvgroupMaps    map[int64]*model.AdwetecPdbAdvgroup
//	PdbAdwetecCreativeMaps    map[int64]*model.AdwetecPdbCreative
//	PdbAdwetecIndustryMaps    map[int64]string
//	PdbAdwetecRestrictionKind map[int64]string
//	PdbAdwetecMedia           map[int64]string
//	PdbAdwetecClientType      map[int64]*model.AdwetecPdbClientType
//	PdbAdwetecClient          map[int64]*model.AdwetecPdbClient
//	PdbAdwetecAdstyle         map[int64]*model.AdwetecPdbAdstyle
//
//	Record               map[string]map[int32][]ResponseTime //响应请求记录 record[api地址][接口响应区间][接口响应时间结构体...]
//	ImageToExcelUrlMap   map[string]*ExcelUrlMapResponse
//	MonitorDownloadExcel map[string]*DownloadExcel
//}
//
//type ExcelUrlMapResponse struct {
//	Code constant.CONSTANT_CODE
//	Data map[string]string
//}
//
//type DownloadExcel struct {
//	Code     constant.CONSTANT_CODE
//	ExcelUrl string
//}
//
//type ResponseTime struct {
//	ReUseTime   int64 //使用时间 毫秒
//	ReStartTime int64 //接口请求开始时间戳 毫秒
//	ReEndTime   int64 //接口请求结束时间戳 毫秒
//}
//
//func (this *ServerHttpContext) InitToken() {
//
//	this.Tokenmutex = new(sync.Mutex)
//
//	this.Token = jwt.New(jwt.SigningMethodHS256)
//
//	claims := make(jwt.MapClaims)
//
//	this.Token.Claims = claims
//
//}

// ******************************************************************************
type ServerContext struct {
	iris.Context
	Userid int64
	Roleid int64
	Limit  int64
	Reqid  string
}

// ******************************************************************************
// 对接FEED的API
type BridgeRequestHeader struct {
	// 请求头--代理权限
	Sign       string `json:"sign"`
	UserNum    int    `json:"usernum"`
	TgUsername string `json:"tgUsername"`
	TgPassword string `json:"tgPassword"` // 加密密码
}

type CommonFeedRequestHeader struct {
	// 请求头--普通权限
	OpUsername string `json:"opUsername"`
	OpPassword string `json:"opPassword"`
	TgUsername string `json:"tgUsername"`
	TgPassword string `json:"tgPassword"`
	BceUser    string `json:"bceUser"`
}
type ManagerFeedRequestHeader struct {
	// 请求头--管家权限
	OpUsername string `json:"opUsername"`
	OpPassword string `json:"opPassword"`
	TgUsername string `json:"tgUsername"`
	TgPassword string `json:"tgPassword"`
	BceUser    string `json:"bceUser"`
}
type AgentFeedRequestHeader struct {
	// 请求头--代理权限
	OpUsername string `json:"opUsername"`
	OpPassword string `json:"opPassword"`
	TgUsername string `json:"tgUsername"`
	TgPassword string `json:"tgPassword"`
	BceUser    string `json:"bceUser"`
}

func NewAgentFeedRequestHeader(tgName, tgPassword string) *AgentFeedRequestHeader {
	return &AgentFeedRequestHeader{
		OpUsername: constant.FEED_API_AGENT_ACCOUNT,
		OpPassword: constant.FEED_API_AGENT_PASSWORD,
		TgUsername: tgName,
		TgPassword: tgPassword,
		BceUser:    constant.BCE_USER,
	}
}

// ******************************************************************************
type FeedResponseHeader struct {
	// FEED接口返回中的消息头
	Desc     string      `json:"desc"`
	Failures interface{} `json:"failures"`
	Oprs     int32       `json:"oprs"`
	Succ     int32       `json:"succ"`
	Status   int32       `json:"status"`
	Oprtime  int32       `json:"oprtime"`
}

func (this *FeedResponseHeader) CheckHeader() bool { // 验证请求是否正常
	if this.Desc == "success" {
		return true
	} else {
		return false
	}
}
func (this *FeedResponseHeader) Is31DaysEarlierError() bool {

	codestr := ""

	code := this.Failures.([]interface{})[0].(map[string]interface{})["code"]

	switch code.(type) {
	case float64:
		codestr = fmt.Sprintf("%d", int(code.(float64)))
	case string:
		codestr = code.(string)
	default:
		fmt.Println("未识别的类型: " + reflect.TypeOf(code).String())
		return false
	}
	if strings.Contains(this.Desc, "failure") && (codestr == "9011036" || codestr == "9121036") {
		return true
	}

	return false

}
func (this *FeedResponseHeader) IsPasswdError() bool {

	codestr := ""

	code := this.Failures.([]interface{})[0].(map[string]interface{})["code"]

	switch code.(type) {
	case float64:
		codestr = fmt.Sprintf("%d", int(code.(float64)))
	case string:
		codestr = code.(string)
	default:
		fmt.Println("未识别的类型: " + reflect.TypeOf(code).String())
		return false
	}

	if strings.Contains(this.Desc, "failure") && codestr == "99914" {
		return true
	}

	return false
}
func (this *FeedResponseHeader) IsNeedRetry() bool {

	codestr := ""

	code := this.Failures.([]interface{})[0].(map[string]interface{})["code"]

	switch code.(type) {
	case float64:
		codestr = fmt.Sprintf("%d", int(code.(float64)))
	case string:
		codestr = code.(string)
	default:
		fmt.Println("未识别的类型: " + reflect.TypeOf(code).String())
		return false
	}

	if strings.Contains(this.Desc, "failure") && (codestr == "9011519" ||
		codestr == "91288801" ||
		codestr == "9011380" ||
		codestr == "9011111" ||
		codestr == "9013001" ||
		codestr == "9013002" ||
		codestr == "901919" ||
		codestr == "912919" ||
		codestr == "901177" ||
		codestr == "901761" ||
		codestr == "901897" ||
		codestr == "99909" ||
		codestr == "99915" ||
		codestr == "99920" ||
		codestr == "9014" ||
		codestr == "9124") {
		return true
	}

	return false
}

// ******************************************************************************
type AgentSemRequestHeader struct {
	// 请求头--代理权限
	OpUsername string `json:"opUsername"`
	OpPassword string `json:"opPassword"`
	TgUsername string `json:"tgUsername"`
	TgPassword string `json:"tgPassword"`
	BceUser    string `json:"bceUser"`
}

func NewAgentSemRequestHeader(tgName, tgPassword string) *AgentSemRequestHeader {
	return &AgentSemRequestHeader{
		OpUsername: constant.SEM_API_AGENT_ACCOUNT,
		OpPassword: constant.SEM_API_AGENT_PASSWORD,
		TgUsername: tgName,
		TgPassword: tgPassword,
		BceUser:    constant.BCE_USER,
	}
}

// ******************************************************************************
type SemResponseHeader struct {
	// SEM接口返回中的消息头
	Desc     string      `json:"desc"`
	Failures interface{} `json:"failures"`
	Oprs     int32       `json:"oprs"`
	Succ     int32       `json:"succ"`
	Status   int32       `json:"status"`
	Oprtime  int32       `json:"oprtime"`
}

func (this *SemResponseHeader) CheckHeader() bool { // 验证请求是否正常

	if this.Desc == "success" {
		return true
	} else {
		return false
	}

}
func (this *SemResponseHeader) Is31DaysEarlierError() bool {

	codestr := ""

	code := this.Failures.([]interface{})[0].(map[string]interface{})["code"]

	switch code.(type) {
	case float64:
		codestr = fmt.Sprintf("%d", int(code.(float64)))
	case string:
		codestr = code.(string)
	default:
		fmt.Println("未识别的类型: " + reflect.TypeOf(code).String())
		return false
	}

	if strings.Contains(this.Desc, "failure") && codestr == "9011036" {
		return true
	}

	return false

}
func (this *SemResponseHeader) IsPasswdError() bool {

	codestr := ""

	code := this.Failures.([]interface{})[0].(map[string]interface{})["code"]

	switch code.(type) {
	case float64:
		codestr = fmt.Sprintf("%d", int(code.(float64)))
	case string:
		codestr = code.(string)
	default:
		fmt.Println("未识别的类型: " + reflect.TypeOf(code).String())
		return false
	}

	if strings.Contains(this.Desc, "failure") && codestr == "99914" {
		return true
	}

	return false
}

func (this *SemResponseHeader) IsNeedRetry() bool {

	codestr := ""

	code := this.Failures.([]interface{})[0].(map[string]interface{})["code"]

	switch code.(type) {
	case float64:
		codestr = fmt.Sprintf("%d", int(code.(float64)))
	case string:
		codestr = code.(string)
	default:
		fmt.Println("未识别的类型: " + reflect.TypeOf(code).String())
		return false
	}

	if strings.Contains(this.Desc, "failure") && (codestr == "9011519" ||
		codestr == "91288801" ||
		codestr == "9011111" ||
		codestr == "9013001" ||
		codestr == "9011380" ||
		codestr == "9013002" ||
		codestr == "912919" ||
		codestr == "901919" ||
		codestr == "901761" ||
		codestr == "901897" ||
		codestr == "901164" ||
		codestr == "901177" ||
		codestr == "99909" ||
		codestr == "99915" ||
		codestr == "99920" ||
		codestr == "9014" ||
		codestr == "9013") {
		return true
	}

	return false
}
func (this *SemResponseHeader) KeywordNumbersLimit() bool {

	codestr := ""

	code := this.Failures.([]interface{})[0].(map[string]interface{})["code"]

	switch code.(type) {
	case float64:
		codestr = fmt.Sprintf("%d", int(code.(float64)))
	case string:
		codestr = code.(string)
	default:
		fmt.Println("未识别的类型: " + reflect.TypeOf(code).String())
		return false
	}

	if strings.Contains(this.Desc, "failure") && (codestr == "901125") {
		return true
	}

	return false
}

// ******************************************************************************
// 对接神马API
type ShenmaRequestHeader struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	// Target   string `json:"target"` // 如果USERNAME标识的账户是账户管家账户则TARGET字段表示被该账户管家管辖和操作的账户名
}
type ShenmaResponseHeader struct {
	Status    int32       `json:"status"`    // 0:执行成功 1:部分动作执行成功、部分失败 2:执行失败 3:系统内部错误
	Desc      string      `json:"desc"`      //
	Quota     int32       `json:"quota"`     // 本次操作消耗的配额值
	LeftQuota int32       `json:"leftQuota"` //
	Failures  interface{} `json:"failures"`  //
}

func NewShenmaRequestHeader(tgName, tgPassword string) *ShenmaRequestHeader {
	return &ShenmaRequestHeader{
		Username: tgName,
		Password: tgPassword,
		Token:    constant.ALI_SHENMA_TOKEN,
	}
}
func (this *ShenmaResponseHeader) CheckHeader() bool { // 验证请求是否正常
	if this.Desc == "执行成功" || this.Desc == "success" {
		return true
	} else {
		return false
	}
}

// ******************************************************************************
// 对接汇川API
type HuichuanRequestHeader struct {
	Username string `json:"username"` // 账户名称
	Password string `json:"password"` // 账户密码
	Token    string `json:"token"`    // 申请汇川API时得到的TOKEN值
	// 注: 代理商若想请求名下子账户的数据使用子账户的账户名称和密码加上代理商TOKEN即可
}
type HuichuanResponseHeader struct {
	Status   int32       `json:"status"` // 0:执行成功 1:部分动作执行成功、部分失败 2:执行失败 3:系统内部错误
	Desc     string      `json:"desc"`
	Failures interface{} `json:"failures"` // (是个数组)错误消息--错误码、错误消息描述
}

func NewHuichuanRequestHeader(tgName, tgPassword string) *HuichuanRequestHeader {
	return &HuichuanRequestHeader{
		Username: tgName,
		Password: tgPassword,
		Token:    constant.ALI_HUICHUAN_TOKEN,
	}
}
func (this *HuichuanResponseHeader) CheckHeader() bool { // 验证请求是否正常
	if this.Desc == "执行成功" {
		return true
	} else {
		return false
	}
}

// ******************************************************************************

func GetAuth(now, path string) string { // 生成AUTH字符串
	return utils.GenAuth(
		constant.BCE_ACCESS_KEY,
		constant.BCE_SECRET_KEY,
		now,
		path,
		"POST",
	)
}

func (this *HuichuanResponseHeader) IsPasswdError() bool {

	codestr := ""

	code := this.Failures.([]interface{})[0].(map[string]interface{})["code"]

	switch code.(type) {
	case float64:
		codestr = fmt.Sprintf("%d", int(code.(float64)))
	case string:
		codestr = code.(string)
	default:
		fmt.Println("未识别的类型: " + reflect.TypeOf(code).String())
		return false
	}

	if strings.Contains(this.Desc, "执行失败") && codestr == "8603" {
		return true
	}

	return false
}

func (this *ShenmaResponseHeader) IsPasswdError() bool {

	codestr := ""

	code := this.Failures.([]interface{})[0].(map[string]interface{})["code"]

	switch code.(type) {
	case float64:
		codestr = fmt.Sprintf("%d", int(code.(float64)))
	case string:
		codestr = code.(string)
	default:
		fmt.Println("未识别的类型: " + reflect.TypeOf(code).String())
		return false
	}

	if strings.Contains(this.Desc, "执行失败") && codestr == "8102" {
		return true
	}

	return false
}
