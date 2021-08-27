package models

import (
	"bluebell/pkg/time"
)

type Community struct {
	CommunityID   int    `db:"community_id"`
	CommunityName string `db:"community_name"`
}

type CommunityDetail struct {
	CommunityID   int64     `json:"id"db:"community_id"`
	CommunityName string    `json:"name"db:"community_name"`
	Introduction  string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime    time.Time `json:"create_time" db:"create_time"`
}
