package controllers

import (
	"github.com/phpxin/mdblog/core"
	model "github.com/phpxin/mdblog/models"
	"html/template"
)

type IndexController struct {
}

//https://startbootstrap.com/templates/blog-home/
func (ctrl *IndexController) Index(ctx *core.HttpContext) (resp *core.HttpResponse) {
	keywords, _ := ctx.GetString("keywords", "")
	pagen, err := ctx.GetInt32("p", 1)
	if err != nil {
		pagen = 1
	}

	var limit int32 = 5

	var docs []*model.Doc
	var amount int32

	if keywords != "" {
		docs, amount = model.SearchDocs(keywords, pagen, limit)
	} else {
		docs, amount = model.GetDocsByPage(pagen, limit)
	}

	var prevPage int32 = -1
	var nextPage int32 = -1

	if pagen > 1 {
		prevPage = pagen - 1
	}
	if pagen*limit < amount {
		nextPage = pagen + 1
	}

	hot := model.GetHotRanging()

	sidebar := sidebar(core.SubjectIndexer, hot)
	nav := nav()
	footer := footer()
	analytics := analytics()

	return core.HtmlResponse("index", struct {
		List      []*model.Doc
		Sidebar   template.HTML
		Nav       template.HTML
		Footer    template.HTML
		Analytics template.HTML
		PrevPage  int32
		NextPage  int32
		Keywords  string
	}{
		docs,
		template.HTML(sidebar),
		template.HTML(nav),
		template.HTML(footer),
		template.HTML(analytics),
		prevPage,
		nextPage,
		keywords,
	})
}
