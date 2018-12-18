package main

import (
	"github.com/kataras/iris"
	"community-go/controller/follow/content"
	"community-go/controller/follow/replys"
	"community-go/utility"
	"community-go/controller/follow/follow"
	"community-go/controller/follow/comments"
	"fmt"
	"community-go/controller/waterpool"
	"database/sql"
	"community-go/utility/redisTool"
	"community-go/config"
	"github.com/go-redis/redis"
	"community-go/community-goTests"
	)

func connectionDB() {
	//链接本地数据库
	db, _ := sql.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库链接失败", err.Error())
	} else {
		fmt.Println("数据库链接成功")
	}
	utility.DB = db
}

func connectionredis()  {
	var client = redis.NewClient(&redis.Options{
		Addr: config.RedisConfig.Address,
		Password: config.RedisConfig.Password,
		DB:config.RedisConfig.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis链接失败")
	} else {
		fmt.Println("redis链接成功")
		redisTool.RedisDB = client
	}
}

func init() {
	//connectionDB()//初始化数据库
	connectionredis()
}

func main() {
	app := iris.Default()

	connectionInterface(app)
	//server.Connectgrpc()
	connectionInterfaceTest(app)

	app.Run(iris.Addr(config.ServerConfig.PortStr))
	defer utility.DB.Close()

}

func connectionInterface(app *iris.Application)  {

	//crs := cors.New(cors.Options{
	//	AllowedOrigins: []string("*"),
	//	AllowCredentials: true,
	//})

	v1 := app.Party("/api/v1").AllowMethods(iris.MethodOptions)
	{
		//887
		v1.Get("/comments", func(ctx iris.Context) {
			comments.GetComments(ctx)
		})
		v1.Get("/contents", func(ctx iris.Context) {
			content.GetContentList(false, ctx)
		})
		//1001
		v1.Get("/replys",  func(ctx iris.Context) {
			replys.GetReplys(ctx)
		})
		//?userID=119&followID=23
		v1.Get("/follow", func(ctx iris.Context) {
			follow.GetFollow(ctx)
		})
		//?page=0&size=20
		v1.Get("/waterpool", func(ctx iris.Context) {
			waterpool.Getwaterpool(ctx)
		})
	}
}

func connectionInterfaceTest(app *iris.Application) {
	community_goTests.PutPartyTest(app)
	community_goTests.PostPartTest(app)
	community_goTests.GetPartTest(app)
}