package models

import (
	"time"
)

type Comment struct {
	Id     int `json:"id"`
	Floor  int `json:"floor"`
	BookId int `json:"book_id"`
	// DocumentId 评论所属的文档.
	DocumentId int `json:"document_id"`
	// Author 评论作者.
	Author string `json:"author"`
	//MemberId 评论用户ID.
	UserId int `json:"user_id"`
	// IP 评论者的IP地址
	IP string `json:"ip"`
	// 评论日期.
	CreatedAt time.Time `json:"created_at"`
	//Content 评论内容.
	Content string `json:"content"`
	// Approved 评论状态：0 待审核/1 已审核/2 垃圾评论/ 3 已删除
	Approved int `json:"approved"`
	// UserAgent 评论者浏览器内容
	UserAgent string `json:"user_agent"`
	// Parent 评论所属父级
	ParentId  int `json:"parent_id"`
	UpVotes   int `json:"up_votes"`
	DownVotes int `json:"down_votes"`
}
