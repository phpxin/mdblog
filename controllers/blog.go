package controllers

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/phpxin/mdblog/core"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
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

func (ctrl *BlogController) Detail2(r *http.Request) (resp *core.HttpResponse) {
	// @todo 全局参数获取、过滤、格式化、校验插件
	qStr := r.URL.Query()
	mdname := qStr.Get("md")

	obj,ok := core.DocsIndexer[mdname]
	if !ok {
		return core.HtmlResponse("errors/404", nil)
	}

	contents,_ := ioutil.ReadFile(obj.Path)
	//render := blackfriday.NewHTMLRenderer(MarkdownToHtmlCommonHtmlFlags)
	output := blackfriday.Run(contents)
	//md := markdown.New(markdown.XHTMLOutput(true))
	//output:=md.RenderToString(contents)

	doc,_:=goquery.NewDocumentFromReader(bytes.NewReader(output))
	sel :=doc.Find("code")
	sel.Each(func(i int, selection *goquery.Selection) {
		classText,_:=selection.Attr("class")
		//fmt.Println(c)
		codeBlock,_:=selection.Html()
		codeHtml:="<code class=\""+classText+"\"><ol>\n"
		splited := strings.Split(codeBlock, "\n")

		for _,v := range splited {
			codeHtml+="<li>"+v+"</li>\n"
		}
		codeHtml+="</ol></code>\n"

		selection.ReplaceWithHtml(codeHtml)

	})

	docHtml,_ := doc.Html()

	return core.HtmlResponse("detail2", struct{
		Contents template.HTML
		Subjects map[string]*core.TreeFolder
		Menu template.HTML
	}{
		template.HTML(docHtml) ,
		//template.HTML(string(output)) ,
		core.SubjectIndexer,
		template.HTML(Menu) ,
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
	// blackfriday.
	output := blackfriday.Run(contents)

	return core.HtmlResponse("detail2", struct{
		Title string
		Intro string
		Desc string
		Contents template.HTML
		Subjects map[string]*core.TreeFolder
		Menu template.HTML
	}{
		obj.Title,
		obj.Intro,
		obj.Desc,
		template.HTML(string(output)) ,
		core.SubjectIndexer,
		template.HTML(Menu) ,
	})
}