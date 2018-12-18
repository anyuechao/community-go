package model

type Follow struct {
	//被关注人的ID
	FollowID int `json:"followID"`
	//当前用户ID
	UserID int `json:"userId"`
	//关注状态 是否关注 关注为 1 未关注为0
	FollowStatus int `json:"followStatus""`
}
