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
	qStr := r.URL.Query()
	subject := qStr.Get("subject")

	obj,ok := core.SubjectIndexer[subject]
	if !ok {
		return core.HtmlResponse("errors/404", nil)
	}

	subjects := make([]*core.TreeFolder, 0)
	articles := make([]*core.TreeFolder, 0)

	for _,item := range obj.Children {
		if len(item.Children)>0 {
			subjects = append(subjects, item)
		}else{
			articles = append(articles, item)
		}
	}

	return core.HtmlResponse("subject", struct{
		Menu template.HTML
		Subjects []*core.TreeFolder
		Articles []*core.TreeFolder
	}{
		template.HTML(Menu) ,
		subjects ,
		articles,
	})
}

func (ctrl *BlogController) Detail(r *http.Request) (resp *core.HttpResponse) {
	// @todo 全局参数获取、过滤、格式化、校验插件
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