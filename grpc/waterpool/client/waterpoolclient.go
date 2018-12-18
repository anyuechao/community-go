package client

import (
	"google.golang.org/grpc"
	"fmt"
	pb "community-go/grpc/waterpool/proto"
		"golang.org/x/net/context"
	)

const (
	defaultName = "world"
	address = "localhost:50051"
)

func ConnectWPserver() {
	 conn, err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
	defer conn.Close()
	 c := pb.NewWaterflowClient(conn)
	 //namee := defaultName
	 //if len(os.Args) >1 {
	 //	namee = os.Args[1]
	 //}
	 r, err := c.GetWaterFlow(context.Background(), &pb.WaterFlowRequest{PageIndex:0,PageSize:5})
	 if err != nil {
	 	fmt.Println("请求失败: %v", err)
	 }
	 transtion(r)
}

func transtion(reply *pb.WaterFlowReply)  {
	//str, err := json.Marshal(reply)
	//
	//if err != nil {
	//	fmt.Println("JSON转化失败: %v", err)
	//	return
	//}
	//fmt.Println(str)
	fmt.Println(reply)
}
