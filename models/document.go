package models

import (
	"time"
)

type Document struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	// Identify 文档唯一标识
	Identify string `json:"identify"`
	BookId   int    `json:"book_id"`
	ParentId int    `json:"parent_id"`
	Order    int    `json:"order"`
	// Markdown markdown格式文档.
	Markdown string `json:"markdown"`
	// Release 发布后的Html格式内容.
	Release string `json:"release"`
	// Content 未发布的 Html 格式内容.
	Content     string        `json:"content"`
	CreatedAt   time.Time     `json:"created_at"`
	UserId      int           `json:"user_id"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Version     int64         `json:"version"`
	Attachments []*Attachment `gorm:"-" json:"attachments"`
}
