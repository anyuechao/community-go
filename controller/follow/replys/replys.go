package replys

import (
	"fmt"
	"community-go/utility"
	"community-go/model"
	"community-go/response"
	"github.com/kataras/iris"
	)

func GetReplys(ctx iris.Context) {

	replyList := make([]model.Reply, 0)
	var commentID int64
	var err error

	if commentID, err = ctx.URLParamInt64("commentID"); err != nil {
		commentID = 1
		err = nil
	}
	rows, err := utility.DB.Query("SELECT * FROM reply where commentID = ?",commentID)
	defer rows.Close()
	if err != nil {
		fmt.Printf("rows === %p", rows)
		fmt.Println("查询错误")
	} else {

		for rows.Next() {
			reply := model.Reply{}
			rows.Scan(&reply.CommentID, &reply.ReplyID, &reply.ReplyName, &reply.ReplyText, &reply.ReplyTime,
				&reply.UserID, &reply.UserName, &reply.UserIcon, &reply.PriseCount)
			err = nil
			replyList = append(replyList, reply)
		}
		fmt.Println("查询正确")
	}
	response.ResponseFunc(ctx, replyList, err)
}