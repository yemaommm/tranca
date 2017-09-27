package mygin

import (
	"IginServer/lib/Imartini"
	"IginServer/lib/mysqldb"
	"IginServer/lib/mysqldb/obj"
	"IginServer/lib/session"
	"github.com/flosch/pongo2"
	// "fmt"
	"github.com/gin-gonic/gin"
)

type IContext struct {
	DB              *obj.DB
	session         map[string]interface{}
	isChangeSession bool
	*gin.Context
}

func (this *IContext) SetSession(key string, value interface{}) {
	if this.session == nil {
		this.session = session.GetSession(this.Writer, this.Request)
	}
	this.session[key] = value
	this.isChangeSession = true
}

func (this *IContext) GetSession(key string) interface{} {
	if this.session == nil {
		this.session = session.GetSession(this.Writer, this.Request)
	}
	return this.session[key]
}

func (this *IContext) GetAllSession() map[string]interface{} {
	if this.session == nil {
		this.session = session.GetSession(this.Writer, this.Request)
	}
	return this.session
}

func (this *IContext) HTML(str string, context pongo2.Context) {
	Imartini.HTML(this.Writer, this.Request, str, context)
}

func Context(c *gin.Context) *IContext {
	return &IContext{mysqldb.GetConnect(), nil, false, c}
}

func Handler(fn func(*IContext)) func(*gin.Context) {
	return func(c *gin.Context) {
		nc := Context(c)
		defer func() {
			if nc.isChangeSession {
				session.SaveSession(c.Writer, c.Request, nc.session)
			}
		}()
		fn(nc)
	}
}
