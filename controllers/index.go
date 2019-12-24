package controllers

import (
	"github.com/phpxin/mdblog/core"
	model "github.com/phpxin/mdblog/models"
	"github.com/phpxin/mdblog/tools/log"
	"html/template"
	"math"
	"net/http"
	"strconv"
)
type IndexController struct {

}

//https://startbootstrap.com/templates/blog-home/
func (ctrl *IndexController) Index(r *http.Request) (resp *core.HttpResponse) {

	//return core.HtmlResponse("index2", struct{
	//	List map[string]*core.TreeFolder
	//	Subjects map[string]*core.TreeFolder
	//	Menu template.HTML
	//}{
	//	core.DocsIndexer ,
	//	core.SubjectIndexer,
	//	template.HTML(Menu) ,
	//})

	qStr := r.URL.Query()
	page := qStr.Get("p")
	pagen,err := strconv.Atoi(page)
	if err!=nil {
		log.Error("get page failed %s", err.Error())
		pagen = 1
	}

	var limit int32 = 5

	docs,amount := model.GetDocsByPage(int32(pagen), limit)
	prevPage := -1
	nextPage := -1

	if pagen>1 {
		prevPage = pagen-1
	}
	if int32(pagen)*limit < amount {
		nextPage = pagen+1
	}

	subjects := make([]*core.TreeFolder, 0)

	for _,v := range core.SubjectIndexer {

		subjects = append(subjects, v)
	}
	half := int(math.Ceil(float64(len(subjects))/2))

	return core.HtmlResponse("index3", struct{
		List []*model.Doc
		Subjects1 []*core.TreeFolder
		Subjects2 []*core.TreeFolder
		Menu template.HTML
		PrevPage int
		NextPage int
	}{
		docs ,
		subjects[:half],
		subjects[half:],
		template.HTML(Menu) ,
		prevPage,
		nextPage,
	})
}

func (ctrl *IndexController) Regenerate(r *http.Request) (resp *core.HttpResponse) {

	err := core.GenerateTreeFolder()
	if err!=nil {
		return core.ApiError(core.API_ERR_MSG, err.Error())
	}
	Menu = menuHtml(core.GetTreeFolder())

	// @todo sync to database

	return core.ApiSuccess(map[string]interface{}{
		"msg":"generate success",
	})
}