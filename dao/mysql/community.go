package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

var (
	ErrorInvalidId = errors.New("无效id")
)

// GetCommunityList 获取全部社区接口
func GetCommunityList() (communities []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	err = db.Select(&communities, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			zap.L().Warn(" GetCommunityList not data")
			return
		}
	}
	return
}

// GetCommunityDetail 根据id获取社区详情
func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	sqlStr := "select community_id,community_name,introduction,create_time from community where community_id=?"
	data = new(models.CommunityDetail)
	err = db.Get(data, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidId
			zap.L().Warn(" this community not data")
			return
		}
	}
	return
}
