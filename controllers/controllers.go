package controllers

import (
	"bytes"
	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/core"
	model "github.com/phpxin/mdblog/models"
	"github.com/phpxin/mdblog/tools/log"
	"html/template"
	"math"
)

var (
	Menu string
)

func InitController() {
	//Menu = menuHtml(core.GetTreeFolder())
}

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

func sidebar(subjects []*core.TreeFolder, hotArticles []*model.HotDoc) string {

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

func menuHtml(tf *core.TreeFolder) string {
	subMenus := make([]*core.TreeFolder, 0)
	if len(tf.Children)>0 {
		for _,v := range tf.Children {
			if len(v.Children)>0 {
				subMenus = append(subMenus, v)
			}
		}
	}

	if len(subMenus)>0 {
		html := "<ul class=\"nav navbar-nav\">"
		for _,secMenuItem := range subMenus {
			html += secMenuHtml(secMenuItem)
		}
		html += "</ul>"
		return html
	}

	return ""
}

func secMenuHtml(tf *core.TreeFolder) string {
	subMenus := make([]*core.TreeFolder, 0)
	if len(tf.Children)>0 {
		for _,v := range tf.Children {
			if len(v.Children)>0 {
				subMenus = append(subMenus, v)
			}
		}
	}

	if len(subMenus)>0 {
		html := "<li class=\"dropdown\"><a href=\"#\" class=\"dropdown-toggle\" data-toggle=\"dropdown\" data-target=\"dropdownMenu1\" role=\"button\" aria-haspopup=\"true\" aria-expanded=\"false\">"+tf.Title+" <span class=\"caret\"></span></a><ul class=\"dropdown-menu multi-level\" role=\"menu\" aria-labelledby=\"dropdownMenu\">"
		html += "<li><a href=\"/blog?subject="+tf.Name+"\"> -&gt; "+tf.Title+"</a></li>"
		for _,thirdMenuItem := range subMenus {
			html += thirdMenuHtml(thirdMenuItem)
		}
		html += "</ul></li>"
		return html
	}else{
		return "<li><a href=\"/blog?subject="+tf.Name+"\">"+tf.Title+"</a></li>"
	}
}

func thirdMenuHtml(tf *core.TreeFolder) string {
	subMenus := make([]*core.TreeFolder, 0)
	if len(tf.Children)>0 {
		for _,v := range tf.Children {
			if len(v.Children)>0 {
				subMenus = append(subMenus, v)
			}
		}
	}

	if len(subMenus)>0 {
		html := "<li class=\"dropdown-submenu\"><a tabindex=\"-1\" href=\"javascript:void();\">"+tf.Title+"</a><ul class=\"dropdown-menu\">"
		html += "<li><a href=\"/blog?subject="+tf.Name+"\"> -&gt; "+tf.Title+"</a></li>"
		for _,thirdMenuItem := range subMenus {
			html += thirdMenuHtml(thirdMenuItem)
		}
		html += "</ul></li>"
		return html
	}else{
		return "<li><a href=\"/blog?subject="+tf.Name+"\">"+tf.Title+"</a></li>"
	}
}
