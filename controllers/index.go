package controllers

import (
	"github.com/phpxin/mdblog/core"
	"html/template"
	"net/http"
)
type IndexController struct {

}

func (ctrl *IndexController) Index(r *http.Request) (resp *core.HttpResponse) {

	return core.HtmlResponse("index2", struct{
		List map[string]*core.TreeFolder
		Subjects map[string]*core.TreeFolder
		Menu template.HTML
	}{
		core.DocsIndexer ,
		core.SubjectIndexer,
		template.HTML(Menu) ,
	})
}

func (ctrl *IndexController) Regenerate(r *http.Request) (resp *core.HttpResponse) {

	err := core.GenerateTreeFolder()
	if err!=nil {
		return core.ApiError(core.API_ERR_MSG, err.Error())
	}
	Menu = menuHtml(core.GetTreeFolder())
	return core.ApiSuccess(map[string]interface{}{
		"msg":"generate success",
	})
}