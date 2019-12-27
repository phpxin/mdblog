package controllers

import (
	"github.com/phpxin/mdblog/core"
	"github.com/phpxin/mdblog/tools/strutils"
)

type TestController struct {

}

func (ctrl *TestController) Index(ctx *core.HttpContext) (resp *core.HttpResponse) {

	ip := strutils.ClientIP(ctx.RawReq)

	return core.ApiSuccess(map[string]interface{}{
		"ssid":ctx.SessionId ,
		"ip":ip,
	})
}

func (ctrl *TestController) Test(ctx *core.HttpContext) (resp *core.HttpResponse) {

	return core.ApiSuccess(map[string]interface{}{
		"ssid":ctx.SessionId ,
	})
}