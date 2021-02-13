package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`     //社区ID 01 -04
	Name string `json:"name" db:"community_name"` //社区名字
}

type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`     //社区ID
	Name         string    `json:"name" db:"community_name"` //社区名字
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
type Reply struct {
	ID         int64     `json:"id,string" db:"post_id"`
	AuthorID   int64     `json:"author_id" db:"author_id"`
	ReplyID    int64     `json:"reply_id" db:"reply_id"`
	Content    string    `json:"content" db:"content" binding:"required"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"`
	VoteNum          int64              `json:"vote_num"`
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息

}

type ReplyList struct {
	*ApiPostDetail
	Reply []*Reply
}
