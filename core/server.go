package core

import (
	"encoding/json"
	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/tools/log"
	"github.com/phpxin/mdblog/tools/strutils"
	"net/http"
	"reflect"
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

	// @TODO BeforeRouter 这里可以做参数校验、权限验证等中间件操作
	startTime := time.Now().Nanosecond()
	moduleStr := "index"
	action := "index" // 默认方法是 IndexController.Index 方法

	if r.URL.Path != "" && r.URL.Path!="/" {
		routerGroup := strings.Split(strings.ToLower(strings.Trim(r.URL.Path, "/")), "/")
		moduleStr = routerGroup[0]
		if len(routerGroup) > 1 {
			action = routerGroup[1]
		}
	}

	module,ok := routerTable[moduleStr]
	if !ok {
		w.WriteHeader(404)
		return
	}

	action = strutils.UcFirst(action)
	obj := reflect.ValueOf(module)
	method := obj.MethodByName(action)
	if method.Kind() == reflect.Invalid {
		w.WriteHeader(404)
		return
	}

	rets := method.Call([]reflect.Value{reflect.ValueOf(r)})
	hRes := rets[0].Interface().(*HttpResponse)

	// @TODO AfterRouter 这里可以记录程序请求日志、统计程序时长等中间件操作
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
