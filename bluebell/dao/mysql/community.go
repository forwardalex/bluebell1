package mysql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golearn/bluebell/models"
	"strings"
)

//GetCommunityList 查询社区清单
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id,community_name from community`
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select 
			community_id, community_name, introduction, create_time
			from community 
			where community_id = ?
	`
	if err = db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}

//CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id, title, content, author_id, community_id)
	values (?, ?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

//GetPostById 根据帖子ID查询单个帖子的详细数据
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?
	`
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 查询帖子列表函数  第一个问号为查询起始点，第二个位查询个数
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select 
	post_id, title, content, author_id, community_id, create_time
	from post
	ORDER BY create_time
	DESC
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2) // 不要写成make([]*models.Post, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// GetPostListByIDs 根据给定的ids列表查询帖子数据  ids由redis的zset有序排列给出
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)
	`

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...) // !!!!!!
	return
}

func GetPostReply(pid int64) (reply []*models.Reply, err error) {
	sqlStr := `select post_id, author_id,reply_id,content, create_time
	from reply
	where post_id in (?)
	order by create_time
	`
	reply = make([]*models.Reply, 0, 2)
	err = db.Get(&reply, sqlStr, pid)
	return
}
func CreateReply(p *models.Reply) (err error) {
	sqlStr := `insert into reply(
	post_id, author_id, reply_id, content)
	values (?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.AuthorID, p.ReplyID, p.Content)
	return
}
