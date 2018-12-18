package model

type Comments struct {
	//文章ID
	ContentID int `json:"contentID"`
	//评论id
	CommentID int `json:"commentID"`
	//用户ID
	UserID int `json:"userID"`
	//评论昵称
	UserName string `json:"userName"`
	//评论头像
	UserIcon string `json:"userIcon"`
	//评论内容
	CommentText string `json:"commentText"`
	//评论时间
	CommentTime string `json:"commentTime"`
	//点赞数
	PriseCount int `json:"priseCount"`
	//回复条数
	ReplyCount int `json:"replyCount"`
	//评论列表
	Replys []Reply `json:"replys"`
}
