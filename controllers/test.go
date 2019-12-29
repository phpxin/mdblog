package controllers

import (
	"github.com/phpxin/mdblog/core"
)

type TestController struct {

}

func (ctrl *TestController) Index(ctx *core.HttpContext) (resp *core.HttpResponse) {

	return core.ApiSuccess(map[string]interface{}{
		"msg": "It works." ,
	})
}