package models

type Option struct {
	Id     int    `json:"option_id"`
	Title  string `json:"title"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Remark string `json:"remark"`
}
