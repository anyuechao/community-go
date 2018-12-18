package put

import (
	"github.com/kataras/iris"
	"community-go/response"
)

func RequestFuncTest(ctx iris.Context) {
	var	err error
	v1 := ctx.FormValue("v1")
	v2 := ctx.FormValue("v2")

	data := &struct {
		V1 string
		V2 string
	}{
		v1,
		v2,
	}
	response.ResponseFunc(ctx,data,err)
}