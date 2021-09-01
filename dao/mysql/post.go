package mysql

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
)

var ErrorNotData = "没有该帖子"

func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post(post_id,title,content,author_id,community_id) values (?,?,?,?,?)"

	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		return err
	}

	err = redis.CreatePost(p)
	return err
}

func ShowPost(id int64) (post *models.Post, err error) {
	sqlStr := "select post_id,title,content,author_id,community_id,create_time,status from post where post_id = ?"
	post = new(models.Post)
	err = db.Get(post, sqlStr, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return post, err
}

func GetPostList(pageNum, pageSize int) (posts []*models.Post, err error) {
	// 按时间降序，从新到旧
	sqlStr := `select 
        post_id,title,content,author_id,community_id,create_time,status 
		from post 
        Order BY create_time
        desc 
        limit ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (pageNum-1)*pageSize, pageSize)
	return posts, err
}

func GetPostListByIDs(ids []string) (posts []*models.Post, err error) {
	sqlStr := `select
		post_id,title,content,author_id,community_id,create_time,status
		from post
		where post_id in (?)
		order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	// 记得加 ...
	db.Select(&posts, query, args...)
	return
}
