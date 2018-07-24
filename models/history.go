package models

import (
	"time"
)

type History struct {
	Id           int       `json:"id"`
	Action       string    `json:"action"`
	ActionName   string    `json:"action_name"`
	DocumentId   int       `json:"document_id"`
	DocumentName string    `json:"document_name"`
	ParentId     int       `json:"parent_id"`
	Markdown     string    `json:"markdown"`
	Content      string    `json:"content"`
	UserId       int       `json:"user_id"`
	UpdatedAt    time.Time `json:"updated_at"`
	Version      int64     `json:"version"`
}
