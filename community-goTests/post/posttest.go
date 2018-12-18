package post

import (
	"github.com/kataras/iris"
	"community-go/response"
	"fmt"
		"time"
	"strconv"
	"os"
	"io"
	"path"
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

//上传图片
func UploadFuncTest(ctx iris.Context) {

	file, handler, err := ctx.FormFile("uploadfile")//接收客户端传来的文件 uploadfile 与客户端保持一致
	if err != nil {
		fmt.Println("ERROR:%v",err)
		return
	}
	defer file.Close()
	//上传的文件保存路径
	ext := path.Ext(handler.Filename)
	fileNewName := string(time.Now().Format("2006-0102-150405-")) + strconv.Itoa(time.Now().Nanosecond()) + ext
	f, err := os.OpenFile("./community-goTests/uploadresources/" + fileNewName, os.O_WRONLY|os.O_CREATE,0666)
	if err != nil {
		fmt.Println("ERROR:%v",err)
		return
	}
	defer f.Close()
	io.Copy(f,file)
	//http://10.2.0.55:8082/image/2018-1204-104056-959929486.jpg
	imageUrl :=  "image/" + fileNewName
	data := &struct {
		ImageURL string
	}{
		imageUrl,
	}
	fmt.Println("upload ok!%v",data)
	response.ResponseFunc(ctx, data, err)
}


//上传视频
func UploadVideoFuncTest(ctx iris.Context) {

	file, handler, err := ctx.FormFile("uploadfile")//接收客户端传来的文件 uploadfile 与客户端保持一致
	if err != nil {
		fmt.Println("ERROR:%v",err)
		return
	}
	defer file.Close()
	//上传的文件保存路径
	ext := path.Ext(handler.Filename)
	fileNewName := string(time.Now().Format("2006-0102-150405-")) + strconv.Itoa(time.Now().Nanosecond()) + ext
	f, err := os.OpenFile("./community-goTests/video/" + fileNewName, os.O_WRONLY|os.O_CREATE,0666)
	if err != nil {
		fmt.Println("ERROR:%v",err)
		return
	}
	defer f.Close()
	io.Copy(f,file)
	//http://10.2.0.55:8082/image/2018-1204-104056-959929486.jpg
	videoUrl :=  "video/" + fileNewName
	data := &struct {
		VideoURL string
	}{
		videoUrl,
	}
	fmt.Println("upload ok!%v",data)
	response.ResponseFunc(ctx, data, err)
}

//func downloadpath() string {
//	home,_ := utility.Home(true)
//	println(home)
//	f, err := os.OpenFile( home + "/" + fileNewName, os.O_WRONLY|os.O_CREATE,0666)
//	return ""
//}

//func upload(writer http.ResponseWriter, request *http.Request) {
//	request.ParseMultipartForm(32<<20)
//	//接收客户端传来的文件 uploadfile 与客户端保持一致
//	file, handler, err := request.FormFile("uploadfile")
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//
//	defer file.Close()
//	//上传的文件保存在ppp路径下
//	ext := path.Ext(handler.Filename)       //获取文件后缀
//	fileNewName := string(time.Now().Format("20060102150405"))+strconv.Itoa(time.Now().Nanosecond())+ext
//
//	f, err := os.OpenFile("./ppp/"+fileNewName, os.O_WRONLY|os.O_CREATE, 0666)
//	//f, err := os.OpenFile("/Users/anyuechao/go/src/go学习/ppp/"+fileNewName, os.O_WRONLY|os.O_CREATE, 0666)
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	defer f.Close()
//
//	io.Copy(f, file)
//
//	fmt.Fprintln(writer, "upload ok!"+fileNewName)
//}