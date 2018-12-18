package waterpool

import (
	"community-go/utility/redisTool"
	"fmt"
	"encoding/json"
	wp "community-go/grpc/waterpool/proto"
	"github.com/kataras/iris"
	"community-go/response"
	)

func Getwaterpool(ctx iris.Context){
	articlelist := make([]*wp.Article,0)
	page, err := ctx.URLParamInt64("page")
	fmt.Println("page:",page)
	if err != nil {
		fmt.Println("page错误")
		response.ResponseFunc(ctx, articlelist, err)
		return
	}
	size, err := ctx.URLParamInt64("size")
	if err != nil {
		fmt.Println("size错误")
		response.ResponseFunc(ctx, articlelist, err)
		return
	}
	data, err := redisTool.RedisDB.ZRange("data:articles2",(page - 1) * size,(page * size) -1 ).Result()
	if err != nil {
		fmt.Println("redis查询错误")
		response.ResponseFunc(ctx, articlelist, err)
		return
	}
	for _,val := range data{
		str := []byte(val)
		a := &wp.Article{}
		err := json.Unmarshal(str,&a)
		if err != nil {
			fmt.Println(err)
			response.ResponseFunc(ctx, articlelist, err)
			return
		}
		fmt.Print("a=========",a)
		articlelist = append(articlelist, a)
	}
	response.ResponseFunc(ctx, articlelist, err)
}