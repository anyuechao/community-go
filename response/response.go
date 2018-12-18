package response

import (
	"github.com/kataras/iris"
	"community-go/model"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SendErrorJson(message string, ctx iris.Context) {
	ctx.JSON(iris.Map{
		"errorCode": model.ERROR,
		message: message,
		"data": iris.Map{},
	})
}

func ResponseFunc(ctx iris.Context, res interface{}, err error) {
	response := Response{}
	//data := Follows{}
	if err != nil {
		//注册失败，返回空用户
		response.Code = 1
		response.Msg = err.Error()
	} else {
		response.Msg = "success"
		response.Code = 0
	}
	response.Data = res
	ctx.JSON(response)
}