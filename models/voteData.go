package models

type VoteData struct {
	PostID    int64 `json:"post_id,string" binding:"required"`
	Direction int8  `json:"direction" binding:"oneof=1 0 -1"` // 赞成票为1，反对票为-1,取消投票为0
}
