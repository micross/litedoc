package models

import (
	"time"
)

type Blog struct {
	Id    string `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	//文章标识
	BlogIdentify string `json:"blog_identify"`
	//排序序号
	Order int `json:"order"`
	//所属用户
	UserId int `json:"user_id"`
	//用户头像
	UserAvatar string `gorm:"-" json:"user_avatar"`
	//文章类型:0 普通文章/1 链接文章
	Type int `json:"type"`
	//链接到的项目中的文档ID
	DocumentId int `json:"document_id"`
	//文章的标识
	DocumentIdentify string `gorm:"-" json:"document_identify"`
	//关联文档的项目标识
	BookIdentify string `gorm:"-" json:"book_identify"`
	//关联文档的项目ID
	BookId int `gorm:"-" json:"book_id"`
	//文章摘要
	BlogExcerpt string `json:"blog_excerpt"`
	//文章内容
	BlogContent string `json:"blog_content"`
	//发布后的文章内容
	BlogRelease string `json:"blog_release"`
	//文章当前的状态，枚举enum(’publish’,’draft’,’password’)值，publish为已 发表，draft为草稿，password 为私人内容(不会被公开) 。默认为publish。
	BlogStatus string `json:"blog_status"`
	//文章密码，varchar(100)值。文章编辑才可为文章设定一个密码，凭这个密码才能对文章进行重新强加或修改。
	Password string `json:"-"`
	//最后修改时间
	UpdatedAt time.Time `json:"updated_at"`
	//创建时间
	CreatedAt  time.Time `json:"created_at"`
	CreateName string    `gorm:"-" json:"create_name"`
	//版本号
	Version int64 `json:"version"`
	//附件列表
	Attachments []*Attachment `gorm:"-" json:"attachments"`
}
