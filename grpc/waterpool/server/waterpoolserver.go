package server

import (
	"golang.org/x/net/context"
	wp "community-go/grpc/waterpool/proto"
	"community-go/utility/redisTool"
	"fmt"
	"encoding/json"
	"net"
	"google.golang.org/grpc"
	"community-go/config"
)

type WaterflowServer struct {

}
func (s *WaterflowServer) GetWaterFlow(ctx context.Context, in *wp.WaterFlowRequest) (*wp.WaterFlowReply, error) {
	data, err := redisTool.RedisDB.ZRange("data:articles2",int64((in.PageIndex-1)*in.PageSize),int64((in.PageIndex * in.PageSize)-1)).Result()
	//data, err := redisTool.RedisDB.ZRange("data:articles2",1,5).Result()
	articlelist := make([]*wp.Article,0)
	if err != nil {
		fmt.Println("redis查询错误")
		//data := map[string]interface{}{}
		return &wp.WaterFlowReply{Message: "faild", Success:false,Data:articlelist}, err
	}
	//articlelist := make([]interface{}, 0)
	for _,val := range data{
		str := []byte(val)
		a := &wp.Article{}
		err := json.Unmarshal(str,&a)
		if err != nil {
			fmt.Println(err)
			//data := map[string]interface{}{}
			return &wp.WaterFlowReply{Message: "faild", Success:false,Data:articlelist}, err
		}
		fmt.Print("a=========",a)
		articlelist = append(articlelist, a)
	}
	return &wp.WaterFlowReply{Message: "Success", Success:true,Data:articlelist}, nil
}

func Connectgrpc() {
	port := config.ServerConfig.PortStr
	lis, err := net.Listen("tcp",port)
	if err != nil {
		fmt.Println("grpc链接失败")
	} else {
		fmt.Println("grpc链接成功")
		s := grpc.NewServer()
		server := &WaterflowServer{}
		wp.RegisterWaterflowServer(s,server)
		s.Serve(lis)
	}
}
