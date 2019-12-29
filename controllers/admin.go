package controllers

import (
	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/core"
)

type AdminController struct {
}

func (ctrl *AdminController) Regenerate(ctx *core.HttpContext) (resp *core.HttpResponse) {
	auth, _ := ctx.GetString("a", "")
	if auth == "" || auth != conf.ConfigInst.Adminkey {
		return core.ApiError(core.API_ERR_MSG, "please confirm your admin-key")
	}

	err := core.GenerateTreeFolder()
	if err != nil {
		return core.ApiError(core.API_ERR_MSG, err.Error())
	}

	return core.ApiSuccess(map[string]interface{}{
		"msg": "generate success",
	})
}
