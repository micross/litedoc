package models

type Relationship struct {
	Id     int `json:"id"`
	UserId int `json:"member_id"`
	BookId int `json:"book_id"`
	// RoleId 角色：0 创始人(创始人不能被移除) / 1 管理员/2 编辑者/3 观察者
	RoleId int `json:"role_id"`
}
