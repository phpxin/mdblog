package controllers

import (
	"github.com/phpxin/mdblog/core"
)

var (
	Menu string
)

func InitController() {
	Menu = menuHtml(core.GetTreeFolder())
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
