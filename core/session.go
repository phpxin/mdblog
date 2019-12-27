package core

import (
	"github.com/hashicorp/go-uuid"
	"github.com/phpxin/mdblog/conf"
	"github.com/phpxin/mdblog/tools/log"
	"net/http"
	"time"
)

const (
	SessionKeyName = "MDBLOG_SESSID"
	SessionKeyExpired = 3600 // s
)

func SessionId() string {
	u,err := uuid.GenerateUUID()
	if err!=nil {
		log.Error("", "get session failed, generate uuid failed")
		return ""
	}

	return u
}

func SessionInit(ctx *HttpContext, respWriter http.ResponseWriter) {
	c,err := ctx.RawReq.Cookie(SessionKeyName)

	if err!=nil {
		sid := SessionId()
		c = new(http.Cookie)
		c.Name = SessionKeyName
		c.Value = sid
		c.Path = "/"
		if conf.ConfigInst.Webhost!="" {
			c.Domain = conf.ConfigInst.Webhost
		}
		c.Expires = time.Unix(time.Now().Unix()+SessionKeyExpired, 0)
		c.HttpOnly = true

		//req.AddCookie(c)
		http.SetCookie(respWriter, c)
	}

	ctx.SessionId = c.Value
}