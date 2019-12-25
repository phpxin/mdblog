package controllers

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/core"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
	"math"
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

	// todo fetch articles from db

	for _,item := range obj.Children {
		if len(item.Children)>0 {
			subjects = append(subjects, item)
		}else{
			articles = append(articles, item)
		}
	}

	half := int(math.Ceil(float64(len(subjects))/2))

	return core.HtmlResponse("subject3", struct{
		Menu template.HTML
		Subjects1 []*core.TreeFolder
		Subjects2 []*core.TreeFolder
		Articles []*core.TreeFolder
	}{
		template.HTML(Menu) ,
		subjects[:half] ,
		subjects[half:] ,
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

	contents,_ := ioutil.ReadFile(conf.ConfigInst.Docroot+"/"+obj.Path)
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

	contents,_ := ioutil.ReadFile(conf.ConfigInst.Docroot+"/"+obj.Path)
	// blackfriday.
	output := blackfriday.Run(contents)

	title:=obj.Title
	title = strings.Replace(title, "-", " ", -1)
	title = strings.Replace(title, ".md", "", -1)

	return core.HtmlResponse("detail3", struct{
		Title string
		Intro string
		Desc string
		Contents template.HTML
		Subjects map[string]*core.TreeFolder
		Menu template.HTML
	}{
		title,
		obj.Intro,
		obj.Desc,
		template.HTML(string(output)) ,
		core.SubjectIndexer,
		template.HTML(Menu) ,
	})
}