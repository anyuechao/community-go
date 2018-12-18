package follow

import (
	"community-go/utility"
	"fmt"
	"community-go/model"
	"github.com/kataras/iris"
	"community-go/response"
)

func GetFollow(ctx iris.Context) {
	var userID int64
	var followID int64
	var err error
	followStatus := 0
	if userID, err = ctx.URLParamInt64("userID"); err != nil {
		userID = 1
		err = nil
	}
	if followID, err = ctx.URLParamInt64("followID"); err != nil {
		followID = 1
		err = nil
	}
	follow := model.Follow{}
	err = utility.DB.QueryRow("SELECT follow.followStatus FROM follow where userId = ? AND followID = ?",userID, followID).Scan(&follow.FollowStatus)
	if err != nil {
		follow.FollowStatus = 0
		fmt.Println("查询错误")
		followStatus = 0
	} else {
		err = nil
		fmt.Println("查询正确")
		followStatus = follow.FollowStatus
	}
	response.ResponseFunc(ctx, followStatus, err)
}