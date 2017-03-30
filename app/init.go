package app

import (
	"fmt"
	"time"

	iris "gopkg.in/kataras/iris.v6"
)

func Init() iris.HandlerFunc {
	return func(ctx *iris.Context) {
		now := time.Now()

		// rc.Response = &access.LogResponseWriter{rc.Response, http.StatusOK, 0}

		ac := newRequestScope(now, ctx.Request)
		// ac := newRequestScope(now, logger, rc.Request)
		ctx.Set("Context", ac)

		fmt.Printf("=> %+v\n", "here we go")

		// TODO: error validation here
		// fault.Recovery(ac.Errorf, convertError)(rc)

		// logAccess(rc, ac.Infof, ac.Now())

		// return nil
		ctx.Next()
	}
}

func GetRequestScope(c *iris.Context) RequestScope {
	return c.Get("Context").(RequestScope)
}
