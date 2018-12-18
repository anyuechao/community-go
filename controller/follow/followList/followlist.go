package followList

import (
	"github.com/kataras/iris"
	"community-go/response"
	content2 "community-go/controller/follow/content"
)

func FollowAction(app *iris.Application) {
	app.Handle("GET", "followList", func(ctx iris.Context) {
		contentList, contentErr := content2.GetContent()
		contentDic := map[string]interface{}{
			"contentList": contentList,
		}
		responseFunc(ctx, contentDic, contentErr)
	})
}

func responseFunc(ctx iris.Context, res interface{}, err error) {
	response := response.Response{}
	//data := Follows{}
	if err != nil {
		//注册失败，返回空用户
		response.Code = 1
		response.Msg = "Faild"
	} else {
		response.Code = 0
		response.Msg = "Success"
	}
	response.Data = res
	ctx.JSON(response)
}