package models

import (
	"time"
)

type Attachment struct {
	Id         int       `json:"id"`
	BookId     int       `json:"book_id"`
	DocumentId int       `json:"doc_id"`
	FileName   string    `json:"file_name"`
	FilePath   string    `json:"file_path"`
	FileSize   float64   `json:"file_size"`
	HttpPath   string    `json:"http_path"`
	FileExt    string    `json:"file_ext"`
	CreatedAt  time.Time `json:"created_at"`
}
