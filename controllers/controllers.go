package controllers

import (
	"bytes"
	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/core"
	model "github.com/phpxin/mdblog/models"
	"github.com/phpxin/mdblog/tools/log"
	"html/template"
	"math"
	"sort"
)

func footer() string {
	t, _ := template.ParseFiles(conf.ConfigInst.Resourcepath+"/htmls/footer.html")
	//执行模板
	buf := make([]byte, 0)
	wbf := bytes.NewBuffer(buf)
	err := t.Execute(wbf, nil)
	if err!=nil {
		log.Error("", "read template wrong, %s", err.Error())
		return ""
	}

	return wbf.String()
}

func nav() string {
	t, _ := template.ParseFiles(conf.ConfigInst.Resourcepath+"/htmls/nav.html")
	//执行模板
	buf := make([]byte, 0)
	wbf := bytes.NewBuffer(buf)
	err := t.Execute(wbf, nil)
	if err!=nil {
		log.Error("", "read template wrong, %s", err.Error())
		return ""
	}

	return wbf.String()
}

func sidebar(submap map[string]*core.TreeFolder, hotArticles []*model.HotDoc) string {

	subjects := make([]*core.TreeFolder, 0)
	sublen:=len(submap)
	subKeys := make([]string, sublen)
	i:=0
	for k,_ := range submap {
		subKeys[i] = k
		i++
	}

	sort.Strings(subKeys)

	for _,k := range subKeys {
		subjects = append(subjects, submap[k])
	}

	half := int(math.Ceil(float64(len(subjects))/2))

	t, _ := template.ParseFiles(conf.ConfigInst.Resourcepath+"/htmls/sidebar.html")
	//执行模板
	buf := make([]byte, 0)
	wbf := bytes.NewBuffer(buf)
	err := t.Execute(wbf, struct{
		HotArticles []*model.HotDoc
		Subjects1 []*core.TreeFolder
		Subjects2 []*core.TreeFolder
	}{
		hotArticles ,
		subjects[:half],
		subjects[half:],
	})
	if err!=nil {
		log.Error("", "read template wrong, %s", err.Error())
		return ""
	}

	return wbf.String()
}
