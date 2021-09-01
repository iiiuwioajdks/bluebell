package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"github.com/yitter/idgenerator-go/idgen"
	"go.uber.org/zap"
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

// GetPostList2 按分数或者时间获取帖子列表
// p.order 是 time 或 score ，代表按照时间查或者按照分数查
// 去redis查id，然后去mysql查帖子详细
func GetPostList2(p *models.ParamPostList) (data []*models.PostDetail, err error) {
	// 1.去redis查id
	postIDs, err := redis.GetPostIDInOrder(p)
	if err != nil {
		return nil, err
	}

	if len(postIDs) == 0 {
		zap.L().Warn("redis get id not data")
		return nil, nil
	}
	// 2. 查 mysql
	posts, err := mysql.GetPostListByIDs(postIDs)

	// 提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(postIDs)
	if err != nil {
		return
	}

	data = make([]*models.PostDetail, 0, len(posts))
	for idx, post := range posts {
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
			VoteNum:       voteData[idx],
			CommunityName: communityName,
			Post:          post,
		}
		data = append(data, postDetail)
	}
	return
}

func GetCommunityPostList(p *models.ParamPostList) (data []*models.PostDetail, err error) {
	// 1.去redis查id
	postIDs, err := redis.GetCommunityPostIDInOrder(p)
	if err != nil {
		return nil, err
	}

	if len(postIDs) == 0 {
		zap.L().Warn("redis get id not data")
		return nil, nil
	}
	// 2. 查 mysql
	posts, err := mysql.GetPostListByIDs(postIDs)

	// 提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(postIDs)
	if err != nil {
		return
	}

	data = make([]*models.PostDetail, 0, len(posts))
	for idx, post := range posts {
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
			VoteNum:       voteData[idx],
			CommunityName: communityName,
			Post:          post,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostListNew(paramPost *models.ParamPostList) (data []*models.PostDetail, err error) {
	if paramPost.CommunityId == 0 {
		data, err = GetPostList2(paramPost)
	} else {
		data, err = GetCommunityPostList(paramPost)
	}
	return data, err
}
