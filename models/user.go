package models

import (
	"time"
)

type User struct {
	Id          int    `json:"id"`
	LoginName   string `json:"login_name"`
	DisplayName string `json:"display_name"`
	Password    string `json:"-"`
	//认证方式: local 本地数据库 /ldap LDAP
	AuthMethod  string `json:"auth_method"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Avatar      string `json:"avatar"`
	//用户角色：0 超级管理员 /1 管理员/ 2 普通用户 .
	Role      int       `json:"role"`
	RoleName  string    `gorm:"-" json:"role_name"`
	Status    int       `json:"status"` //用户状态：0 正常/1 禁用
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
}
