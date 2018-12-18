package model

type Content struct {
	//文章ID
	ContentId int `json:"contentId"`
	//文章作者id
	AuthorID int `json:"authorID"`
	//文章标题
	ContentTitle string `json:"contentTitle"`
	//文章内容
	ContentText string `json:"contentText"`
	//文章简要
	ContentBrief string `json:"contentBrief"`
	//文章视频
	ContentMp4 string `json:"contentMp4"`
	//文章图片切片
	ContentImageArr string `json:"contentImageArr"`
	//文章发布时间
	ReleaseTime string `json:"releaseTime"`
	//文章状态
	ContentStatus int `json:"contentStatus"`
	//分享数目
	ShareCount int `json:"shareCount"`
	//阅读量
	ReadCount int `json:"readCount"`
	//点赞数
	PriseCount int `json:"priseCount"`
	//评论数
	CommentsCount int `json:"commentsCount"`
	//文章查看数
	WatchCount int `json:"watchCount"`
	//文章标签
	ContentTag string `json:"contentTag"`
	//文章类别 0 普通文章 1美丽日记
	ContentType int `json:"contentType"`
	//评论列表
	CommentList []Comments `json:"commentList"`
}