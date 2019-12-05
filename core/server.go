package core

import (
	"encoding/json"
	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/tools/log"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	API_SUCCESS = 0
	API_ERR_MSG = 10001 // 显示错误信息
)

type HttpResponse struct {
	ContentType string
	Content []byte
}

type ApiRet struct {
	Code int32                  `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

func InitServer() {
	http.HandleFunc("/", RpcHandle)
	err := http.ListenAndServe(conf.ConfigInst.RpcHost,nil)
	if err!=nil {
		panic(err)
	}
}

// http 请求代理函数，路由函数
func RpcHandle(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	// BeforeRouter 这里可以做参数校验、权限验证等中间件操作
	startTime := time.Now().Nanosecond()

	var act func (*http.Request) *HttpResponse

	switch r.URL.Path {
	case "/detail":
		act = detail
		break
	case "/":
		act = index
		break
	default:
		w.WriteHeader(404)
		break
	}

	// AfterRouter 这里可以记录程序请求日志、统计程序时长等中间件操作

	hRes := act(r)
	endTime := time.Now().Nanosecond()

	log.Info("performance", "use time %d nano", endTime-startTime)

	w.Header().Set("Content-Type", hRes.ContentType)
	w.Write(hRes.Content)
}

func apiError(w http.ResponseWriter, code int32, msg string) {
	ret := &ApiRet{
		Code:code,
		Msg:msg,
	}

	retj,err := json.Marshal(ret)
	if err!=nil {
		log.Error("system", "marshal json error,%s", err.Error())
		return
	}

	_,err = w.Write(retj)
	if err!=nil {
		log.Error("system", "response error ,%s", err.Error())
		return
	}

	return
}

func apiSuccess(w http.ResponseWriter, data map[string]interface{})  {
	ret := &ApiRet{
		Code:API_SUCCESS,
		Msg:"",
		Data:data,
	}

	retj,err := json.Marshal(ret)
	if err!=nil {
		log.Error("system", "marshal json error,%s", err.Error())
		return
	}

	_,err = w.Write(retj)
	if err!=nil {
		log.Error("system", "response error ,%s", err.Error())
		return
	}

	return
}

func HtmlResponse(content []byte) *HttpResponse {
	return &HttpResponse{
		ContentType: "text/html",
		Content:     content,
	}
}

func detail(r *http.Request) (*HttpResponse) {

	htmlContents,_ := ioutil.ReadFile("./resources/htmls/detail.html")
	// @todo 接收文章名称，获取文章正文
	contents,_ := ioutil.ReadFile("/Users/leo/Documents/Sites/git/phpxin.github.io/_draft/redis.md")
	output := blackfriday.Run(contents)
	htmlResult := strings.Replace(string(htmlContents), "#contents#", string(output), 1)

	return HtmlResponse([]byte(htmlResult))
}

// http : /addtask
// 执行爬虫任务
func index(r *http.Request) (*HttpResponse) {

	return HtmlResponse([]byte("this is index"))
}