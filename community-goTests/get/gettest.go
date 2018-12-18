package get

import (
	"github.com/kataras/iris"
	"community-go/response"
	"os"
	"io/ioutil"
	"community-go/logger"
		"fmt"
)

func RequestFuncTest(ctx iris.Context) {
	var	err error
	v1 := ctx.URLParam("v1")
	v2 := ctx.URLParam("v2")

	data := &struct {
		V1 string
		V2 string
	}{
		v1,
		v2,
	}
	response.ResponseFunc(ctx,data,err)
}

func RequestFuncHeaderTest(ctx iris.Context) {
	var	err error
	v1 := ctx.URLParam("v1")
	v2 := ctx.URLParam("v2")
	access_token := ctx.GetHeader("access_token")
	refresh_token := ctx.GetHeader("refresh_token")

	data := &struct {
		V1 string
		V2 string
		V3 string
		Accesstoken string
		Refreshtoken string
	}{
		v1,
		v2,
		"10",
		access_token,
		refresh_token,
	}
	//"\(access_token)",
	response.ResponseFunc(ctx,data,err)
}

func ReqeustHtmlTest(ctx iris.Context) {
	f, err := os.OpenFile("./community-goTests/uploadresources/index.html",os.O_RDONLY,0666)

	if err != nil {
		logger.Error(err)
		return
	}
	byte, err := ioutil.ReadAll(f)
	if err != nil {
		logger.Error(err)
		return
	}
	ctx.HTML(string(byte))
}

func Requestimage(imgid string, ctx iris.Context) {
	res := GetAvatar(imgid)
	ctx.Write(res)
}

func Requestvideo(videoid string, ctx iris.Context) {
	res := GetAvatarVideo(videoid)
	ctx.Write(res)
}


func GetAvatarVideo(avatarName string) []byte {
	filePath := "community-goTests/video"
	avatarPath := fmt.Sprintf("%s/%s", filePath, avatarName)
	file, err := os.Open(avatarPath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return buff
}


func GetAvatar(avatarName string) []byte {
	filePath := "community-goTests/uploadresources"
	avatarPath := fmt.Sprintf("%s/%s", filePath, avatarName)
	file, err := os.Open(avatarPath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return buff
}

//func followAction(app *iris.Application) {
//
//	//followList.FollowAction(app)
//
//	app.Handle("GET", "followList", func(ctx iris.Context) {
//		//users, userErr := users.GetUser(ctx)
//		//follow, followErr := follow.GetFollow(ctx)
//		//content, contentErr := content.GetContent(ctx)
//		//comment, commentErr := comments.GetComments(content.ContentId)
//		comment, commentErr := comments.GetComments(887)
//		contentList := make([]interface{}, 0)
//		//commentsList := make([]comments.Comments, 0)
//		//commentsList = append(commentsList, comment)
//		//content.CommentList = commentsList
//		contentMap := map[string]interface{}{
//			//"follow":  follow,
//			//"content": content,
//			//"users":   users,
//			"comments" : comment,
//		}
//		contentList = append(contentList, contentMap)
//		//contentDic := map[string]interface{}{
//		//	"contentList": contentList,
//		//}
//		//fmt.Println(contentErr)
//		//fmt.Println(commentErr)
//		//fmt.Println(userErr)
//		responseFunc(ctx, contentList, commentErr)
//
//	})
//}

//func followAction(app *iris.Application) {
//	app.Handle("GET", "follow", func(ctx iris.Context) {
//		//ctx.WriteString("string")
//		res, err := follow.GetFollow(ctx)
//		res1, _ := banner.GetBanner(ctx)
//		group := map[string]interface{} {
//			"follow" : res["followList"],
//			"bannner" : res1["bannerlist"],
//		}
//		responseFunc(ctx, group, err)
//	})
//}