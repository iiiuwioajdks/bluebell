package models

import "bluebell/pkg/time"

// 内存对齐

type Post struct {
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type PostDetail struct {
	AuthName      string `json:"auth_name" db:"username"`
	VoteNum       int64  `json:"votes"`
	CommunityName string `json:"community_name" db:"community_name"`
	*Post         `json:"post"`
}
