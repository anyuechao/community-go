package model

type Users struct {
	//用户ID
	AuthorID int `json:"authorID"`
	//用户名称
	AuthorName string `json:"authorName"`
	//用户头像
	AuthorIcon string `json:"authorIcon"`
}