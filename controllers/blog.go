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
	resp = &core.HttpResponse{
		Content:     []byte("this is index"),
	}

	return resp
}

func (ctrl *BlogController) Detail(r *http.Request) (resp *core.HttpResponse) {
	// @todo 接收文章名称，获取文章正文
	contents,_ := ioutil.ReadFile("/Users/leo/Documents/Sites/git/phpxin.github.io/_draft/redis.md")
	output := blackfriday.Run(contents)

	return core.HtmlResponse("detail", struct{
		Contents template.HTML
	}{
		template.HTML(string(output)) ,
	})
}