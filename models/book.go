package models

import (
	"time"
)

type Book struct {
	Id int `json:"id"`
	// BookName 项目名称.
	Name string `json:"name"`
	// Identify 项目唯一标识.
	Identify string `json:"identify"`
	//是否是自动发布 0 否/1 是
	AutoRelease int `json:"auto_release"`
	//是否开启下载功能 0 是/1 否
	Downloadable int `json:"downloadable"`
	Order        int `json:"order"`
	// Description 项目描述.
	Description string `json:"description"`
	//发行公司
	Publisher string `json:"publisher"`
	Label     string `json:"label"`
	// PrivatelyOwned 项目私有： 0 公开/ 1 私有
	Private int `json:"private"`
	// 当项目是私有时的访问Token.
	ViewToken string `json:"view_token"`
	//状态：0 正常/1 已删除
	Status int `json:"status"`
	//默认的编辑器.
	Editor string `json:"editor"`
	// DocCount 包含文档数量.
	DocCount int `json:"doc_count"`
	// CommentStatus 评论设置的状态:open 为允许所有人评论，closed 为不允许评论, group_only 仅允许参与者评论 ,registered_only 仅允许注册者评论.
	CommentStatus string `json:"comment_status"`
	CommentCount  int    `json:"comment_count"`
	//封面地址
	Cover string `json:"cover"`
	//主题风格
	Theme string `json:"theme"`
	// CreateTime 创建时间 .
	CreateAt time.Time `json:"created_at"`
	//每个文档保存的历史记录数量，0 为不限制
	HistoryCount int `json:"history_count"`
	//是否启用分享，0启用/1不启用
	EnableShare int       `json:"enable_share"`
	UserId      int       `json:"user_id"`
	UpdatedAt   time.Time `json:"updated_at"`
	Version     int64     `json:"version"`
	//是否使用第一篇文章项目为默认首页,0 否/1 是
	UseFirstDocument int `json:"use_first_document"`
}
