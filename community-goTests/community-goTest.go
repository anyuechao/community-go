package community_goTests

import (
	"github.com/kataras/iris"
	"community-go/community-goTests/put"
	"community-go/community-goTests/post"
	"community-go/community-goTests/get"
	"log"
)

func PutPartyTest(app *iris.Application) {
	p1 := app.Party("/api/testput").AllowMethods(iris.MethodOptions)
	{
		p1.Put("/p1", func(ctx iris.Context) {
			put.RequestFuncTest(ctx)
		})
	}
}

func PostPartTest(app *iris.Application) {
	p2 := app.Party("/api/testpost").AllowMethods(iris.MethodOptions)
	{
		p2.Post("/p2", func(ctx iris.Context) {
			post.RequestFuncTest(ctx)
		})
		p2.Post("/upload", func(ctx iris.Context) {
			post.UploadFuncTest(ctx)
		})
		p2.Post("/uploadVideo", func(ctx iris.Context) {
			post.UploadVideoFuncTest(ctx)
		})
	}
}

func GetPartTest(app *iris.Application) {
	p3 := app.Party("/api/testget").AllowMethods(iris.MethodOptions)
	{
		p3.Get("/p3", func(ctx iris.Context) {
			get.RequestFuncTest(ctx)
		})
		p3.Get("/p3/header", func(ctx iris.Context) {
			get.RequestFuncHeaderTest(ctx)
		})
		p3.Get("/html", func(ctx iris.Context) {
			get.ReqeustHtmlTest(ctx)
		})
		//app.Handle("GET", "/users/avatar/{avatarname}", func(ctx iris.Context) {
		//	avatarName := ctx.Params().Get("avatarname")
		//	res := user.GetAvatar(avatarName)
		//	ctx.Write(res)
		//})
	}

	//获取图片
	app.Handle("GET","/image/{imageid}", func(ctx iris.Context){
		log.Println("image request")
		imagid :=  ctx.Params().Get("imageid")
		get.Requestimage(imagid,ctx)
	})

	//获取视频
	//http://10.2.0.55:8082/video/2018-1204-110118-27701111111111.mp4
	app.Handle("GET","/video/{videoid}", func(ctx iris.Context){
		videoid :=  ctx.Params().Get("videoid")
		get.Requestvideo(videoid,ctx)
	})
}

//func index(writer http.ResponseWriter, request *http.Request) {
//	writer.Write([]byte(tpl))
//}
//const tpl = `<html>
//<head>
//<title>上传文件</title>
//</head>
//<body>
//<form enctype="multipart/form-data" action="/upload2" method="post">
//<input type="file" name="uploadfile">
//<input type="hidden" name="token" value="{...{.}...}">
//<input type="submit" value="upload">
//</form>
//</body>
//</html>