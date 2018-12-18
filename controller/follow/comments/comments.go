package comments

import (
	"fmt"
	"community-go/model"
	"github.com/kataras/iris"
	"community-go/utility"
		"community-go/response"
)

//func GetComments(contentID int) ([]interface{}, error) {
//
//	rows, err := utility.DB.Query("SELECT contentID, commentID, userID, userName, userIcon, commentText, commentTime, priseCount FROM comments where contentID = ?",contentID)
//	defer rows.Close()
//	//logger.PrintResult(rows)
//	if err != nil {
//		fmt.Printf("rows === %p", rows)
//		fmt.Println("查询错误")
//		return make([]interface{}, 0), err
//	} else {
//		comments := model.Comments{}
//		commentsList := make([]interface{}, 0)
//		for rows.Next() {
//			rows.Scan(&comments.ContentID, &comments.CommentID, &comments.UserID, &comments.UserName, &comments.UserIcon,
//				&comments.CommentText, &comments.CommentTime, &comments.PriseCount)
//			commentsList = append(commentsList, comments)
//		}
//		//if replyErr != nil {
//		//	fmt.Printf("rows === %p", rows)
//		//	comments.ReplyCount = 0
//		//	comments.Replys = make([]Reply, 0)
//		//	fmt.Println("查询错误")
//		//} else {
//		//	comments.ReplyCount = len(replyList)
//		//	comments.Replys = replyList
//		//	replyErr = nil
//		//}
//		fmt.Println("查询正确")
//		return commentsList, err
//	}
//}

func GetComments(ctx iris.Context) {

	commentsList := make([]interface{}, 0)
	var contentID int64
	var err error

	if contentID, err = ctx.URLParamInt64("contentID"); err != nil {
		contentID = 1
		err = nil
	}

	rows, err := utility.DB.Query("SELECT contentID, commentID, userID, userName, userIcon, commentText, commentTime, priseCount FROM comments where contentID = ?",contentID)
	defer rows.Close()
	//logger.PrintResult(rows)
	if err != nil {
		fmt.Printf("rows === %p", rows)
		fmt.Println("查询错误")
		//return make([]interface{}, 0), err
	} else {
		comments := model.Comments{}

		for rows.Next() {
			rows.Scan(&comments.ContentID, &comments.CommentID, &comments.UserID, &comments.UserName, &comments.UserIcon,
				&comments.CommentText, &comments.CommentTime, &comments.PriseCount)
			commentsList = append(commentsList, comments)
		}
		//if replyErr != nil {
		//	fmt.Printf("rows === %p", rows)
		//	comments.ReplyCount = 0
		//	comments.Replys = make([]Reply, 0)
		//	fmt.Println("查询错误")
		//} else {
		//	comments.ReplyCount = len(replyList)
		//	comments.Replys = replyList
		//	replyErr = nil
		//}
		fmt.Println("查询正确")
		//return commentsList, err
	}
	response.ResponseFunc(ctx, commentsList, err)
}
