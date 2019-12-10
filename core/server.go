package core

import (
	"bytes"
	"encoding/json"
	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/tools/log"
	"github.com/phpxin/mdblog/tools/strutils"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	if len(r.URL.Path)>9 && r.URL.Path[:10]=="/resources" {
		contentType := "text/plain"
		fpath := conf.ConfigInst.Resourcepath+r.URL.Path[10:]
		switch filepath.Ext(fpath) {
		case ".css":
			contentType="text/css"
			break
		case ".js":
			contentType="text/javascript"
			break
		case ".jpg":
		case ".jpeg":
			contentType="image/jpeg"
			break
		case ".png":
			contentType="image/png"
			break
		}

		fp,err := os.Open(fpath)
		if err!=nil {
			w.WriteHeader(404)
			return
		}

		defer fp.Close()

		w.Header().Set("Content-Type", contentType)
		io.Copy(w,fp) // io copy 会调用 sendfile 提升传输效率
		return
	}

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

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Write(hRes.Content)

}

func ApiError(code int32, msg string) *HttpResponse {
	ret := &ApiRet{
		Code:code,
		Msg:msg,
	}

	retj,_ := json.Marshal(ret)
	return &HttpResponse{
		ContentType: "application/json",
		Content:     retj,
	}
}

func ApiSuccess(data map[string]interface{}) *HttpResponse {
	ret := &ApiRet{
		Code:API_SUCCESS,
		Msg:"",
		Data:data,
	}

	retj,_ := json.Marshal(ret)
	return &HttpResponse{
		ContentType: "application/json",
		Content:     retj,
	}
}

func HtmlResponse(templateFile string, vars interface{}) *HttpResponse {
	buf := make([]byte, 0)
	wbf := bytes.NewBuffer(buf)

	t, _ := template.ParseFiles(conf.ConfigInst.Resourcepath+"/htmls/"+templateFile+".html")
	//执行模板
	err := t.Execute(wbf, vars)
	if err!=nil {
		log.Error("", "read template wrong, %s", err.Error())
		return &HttpResponse{
			ContentType: "text/html",
			Content:     []byte("read template wrong"),
		}
	}

	return &HttpResponse{
		ContentType: "text/html",
		Content:     wbf.Bytes(),
	}
}
