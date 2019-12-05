package controllers

import (
	"github.com/phpxin/mdblog/core"
	"net/http"
)

type IndexController struct {

}

func (ctrl *IndexController) Index(r *http.Request) (resp *core.HttpResponse) {
	resp = &core.HttpResponse{
		Content:     []byte("It works."),
	}

	return resp
}