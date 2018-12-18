package content

import (
			"fmt"
			"community-go/utility"
	"github.com/kataras/iris"
			"community-go/response"
	"community-go/model"
)

func GetContent() (interface{}, error) {
	rows, err := utility.DB.Query("SELECT content.contentId, content.authorID, content.contentTitle, content.contentText, content.contentBrief, content.contentMp4, content.ContentImageArr, content.releaseTime, content.contentStatus, content.shareCount, content.readCount, content.priseCount, content.commentsCount, content.watchCount, content.contentTag, content.contentType FROM content")
	//rows, err := utility.DB.Query("SELECT content.contentId, content.authorID, content.contentTitle, content.contentText, content.contentBrief, content.contentMp4, content.ContentImageArr, content.releaseTime, content.contentStatus, content.shareCount, content.readCount, content.priseCount, content.commentsCount, content.watchCount, content.contentTag, content.contentType, comments.contentId, comments.commentID, comments.userID, comments.commentText, comments.commentTime, comments.priseCount  FROM content, comments")
	//rows, err := utility.DB.Query("SELECT * FROM content")
	//logger.PrintResult(rows)
	defer rows.Close()
	if err != nil {
		fmt.Printf("rows === %p", rows)
		fmt.Println("查询错误")
		return make([]model.Content, 0), err
	} else {
		contentList := make([]interface{}, 0)
		//commentsList := make([]comments.Comments, 0)

		for rows.Next() {
			content := model.Content{}
			//comments := comments.Comments{}

			rows.Scan(&content.ContentId, &content.AuthorID, &content.ContentTitle, &content.ContentText, &content.ContentBrief, &content.ContentMp4, &content.ContentImageArr, &content.ReleaseTime, &content.ContentStatus, &content.ShareCount, &content.ReadCount, &content.PriseCount, &content.CommentsCount, &content.WatchCount, &content.ContentTag, &content.ContentType)
			//rows.Scan(&content.ContentId, &content.AuthorID, &content.ContentTitle, &content.ContentText, &content.ContentBrief, &content.ContentMp4, &content.ContentImageArr, &content.ReleaseTime, &content.ContentStatus, &content.shareCount, &content.ReadCount, &content.PriseCount, &content.CommentsCount, &content.WatchCount, &content.ContentTag, &content.ContentType)
			//for rows.Next() {
			//	rows.Scan(&comments.ContentID, &comments.CommentID, &comments.UserID, &comments.CommentText, &comments.CommentTime, &comments.PriseCount)
			//	commentsList = append(commentsList, comments)
			//}
			//user, err := users.GetUser(content.UserID)
			//follow, follorErr := follow.GetFollow(userID, content.UserID)
			//comment, commentErr := comments.GetComments(content.ContentId)
			//if err != nil {
			//	fmt.Println("查询错误")
			//}
			//if follorErr != nil {
			//	fmt.Println("查询错误")
			//}
			//if commentErr != nil {
			//	fmt.Println("查询错误")
			//}
			//commentsList = append(commentsList, comment)
			//content.CommentList = commentsList
			//contentMap := map[string]interface{}{
			//	//"follow":  follow,
			//	"content": content,
			//	//"user":   user,
			//}
			contentList = append(contentList, content)
		}
		fmt.Println("查询正确")
		return contentList, err
	}
}

func GetContentList(getAll bool, ctx iris.Context) {
	contentList := make([]interface{}, 0)
	var pageNum int64
	var err error
	//if pageNum, err = strconv.Atoi(ctx.FormValue("pageNum")); err != nil {
	if pageNum, err = ctx.URLParamInt64("pageNum"); err != nil {
		pageNum = 1
		err = nil
	}

	if pageNum < 1 {
		pageNum = 1
	}

	//pageSize := config.ServerConfig.PageSize
	//
	//offset := (pageNum - 1) * pageSize

	//var startTime string
	//var endTime string
	//
	//// 时间查询
	//if startAt, err := strconv.Atoi(ctx.FormValue("startAt")); err != nil {
	//	// 1970 年的时间
	//	startTime = time.Unix(0, 0).Format("2006-01-02 15:04:05")
	//} else {
	//	startTime = time.Unix(int64(startAt/1000), 0).Format("2006-01-02 15:04:05")
	//}
	//
	//if endAt, err := strconv.Atoi(ctx.FormValue("endAt")); err != nil {
	//	// 当前时间
	//	endTime = time.Now().Format("2006-01-02 15:04:05")
	//} else {
	//	endTime = time.Unix(int64(endAt/1000), 0).Format("2006-01-02 15:04:05")
	//}

	// 默认按创建时间 降序排列

	//order := "created_at"
	//var orderASC string
	//if ctx.Params().Get("asc") == "1" {
	//	orderASC = "ASC"
	//} else {
	//	orderASC = "DASC"
	//}

	//cateID := ctx.FormValue("cateID")
	//
	//cateID := ctx.Params().Get("cateID")
	var categoryID int
	//// Atoi 转换空字符串为 0
	//if categoryID, err = strconv.Atoi(cateID); err != nil {
	//	response.SendErrorJson("错误的分类 ID 类型",ctx)
	//	return
	//}

	if categoryID > 0 {
		//根据分类ID查询

	} else {
		// 没有 categoryId 则获取所有数据
		var sql = "SELECT content.contentId, content.authorID, content.contentTitle, content.contentText, content.contentBrief, content.contentMp4, content.ContentImageArr, content.releaseTime, content.contentStatus, content.shareCount, content.readCount, content.priseCount, content.commentsCount, content.watchCount, content.contentTag, content.contentType FROM content"
		rows, err := utility.DB.Query(sql)
		defer rows.Close()
		if err != nil {
			fmt.Printf("rows === %p", rows)
			fmt.Println("查询错误")
		} else {

			for rows.Next() {
				content := model.Content{}
				rows.Scan(&content.ContentId, &content.AuthorID, &content.ContentTitle, &content.ContentText, &content.ContentBrief, &content.ContentMp4, &content.ContentImageArr, &content.ReleaseTime, &content.ContentStatus, &content.ShareCount, &content.ReadCount, &content.PriseCount, &content.CommentsCount, &content.WatchCount, &content.ContentTag, &content.ContentType)
				contentList = append(contentList, content)
			}
			fmt.Println("查询正确")
			//return contentList, err

		}
		//if getAll {
		//
		//} else {
		//
		//}

	}
	response.ResponseFunc(ctx, contentList, err)
}