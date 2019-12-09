package controllers

import (
	"github.com/phpxin/mdblog/core"
	"net/http"
)

type IndexController struct {

}

func menuHtml(tf *core.TreeFolder) {
	// hasSubMenu := false
	subMenus := make([]*core.TreeFolder, 0)
	if len(tf.Children)>0 {
		for _,v := range tf.Children {
			if len(v.Children)>0 {
				subMenus = append(subMenus, v)
			}
		}
	}

	if len(subMenus)>0 {
		
	}
}

func (ctrl *IndexController) Index(r *http.Request) (resp *core.HttpResponse) {

	return core.HtmlResponse("index", struct{
		List map[string]*core.TreeFolder
		Menu *core.TreeFolder
	}{
		core.DocsIndexer ,
		core.GetTreeFolder() ,
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