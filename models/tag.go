package models

type Tag struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	BookId int    `json:"book_id"`
}
