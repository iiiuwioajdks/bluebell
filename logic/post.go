package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"github.com/yitter/idgenerator-go/idgen"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) error {
	// 生成 post_id
	post.ID = int64(idgen.NextId())
	return mysql.CreatePost(post)
}

// ShowPost 展示帖子信息
func ShowPost(id int64) (data *models.PostDetail, err error) {
	post, err := mysql.ShowPost(id)
	if err != nil {
		return nil, err
	}
	name, err := mysql.GetAuthName(post.AuthorID)
	if err != nil {
		return nil, err
	}
	communityName, err := mysql.GetCommunityName(post.CommunityID)
	if err != nil {
		return nil, err
	}
	data = &models.PostDetail{
		AuthName:      name,
		CommunityName: communityName,
		Post:          post,
	}
	return data, err
}

// GetPostList 获取帖子列表
func GetPostList(pageNum, pageSize int) (data []*models.PostDetail, err error) {
	posts, err := mysql.GetPostList(pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	data = make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		name, err := mysql.GetAuthName(post.AuthorID)
		if err != nil {
			continue
		}
		communityName, err := mysql.GetCommunityName(post.CommunityID)
		if err != nil {
			continue
		}
		postDetail := &models.PostDetail{
			AuthName:      name,
			CommunityName: communityName,
			Post:          post,
		}
		data = append(data, postDetail)
	}
	return
}
