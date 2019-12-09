package controllers

import (
	"github.com/phpxin/mdblog/core"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
	"net/http"
)

type BlogController struct {

}

func (ctrl *BlogController) Index(r *http.Request) (resp *core.HttpResponse) {
	return core.HtmlResponse("index", struct{
		List map[string]*core.TreeFolder
	}{
		core.DocsIndexer ,
	})
}

func (ctrl *BlogController) Detail(r *http.Request) (resp *core.HttpResponse) {
	// @todo 接收文章名称，获取文章正文
	qStr := r.URL.Query()
	mdname := qStr.Get("md")

	obj,ok := core.DocsIndexer[mdname]
	if !ok {
		return core.HtmlResponse("errors/404", nil)
	}

	contents,_ := ioutil.ReadFile(obj.Path)
	output := blackfriday.Run(contents)

	return core.HtmlResponse("detail", struct{
		Contents template.HTML
	}{
		template.HTML(string(output)) ,
	})
}