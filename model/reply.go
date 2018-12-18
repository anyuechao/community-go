package model

type Reply struct {
	//评论id
	CommentID int `json:"commentID"`
	//回复id
	ReplyID int `json:"replyID"`
	//评论昵称
	ReplyName string `json:"replyName"`
	//评论内容
	ReplyText string `json:"replyText"`
	//评论时间
	ReplyTime string `json:"replyTime"`
	//用户ID
	UserID int `json:"userID"`
	//评论昵称
	UserName string `json:"userName"`
	//评论头像
	UserIcon string `json:"userIcon"`
	//点赞数
	PriseCount int `json:"priseCount"`
}
