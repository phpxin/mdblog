package controllers

import (
	"github.com/phpxin/mdblog/core"
	"net/http"
)

type IndexController struct {

}

func (ctrl *IndexController) Index(r *http.Request) (resp *core.HttpResponse) {
	return core.HtmlResponse("index", struct{
		List map[string]*core.TreeFolder
	}{
		core.DocsIndexer ,
	})
}

func (ctrl *IndexController) Regenerate(r *http.Request) (resp *core.HttpResponse) {
	err := core.GenerateTreeFolder()
	if err!=nil {
		return core.ApiError(core.API_ERR_MSG, err.Error())
	}
	return core.ApiSuccess(map[string]interface{}{
		"msg":"generate success",
	})
}