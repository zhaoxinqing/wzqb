package model

import (
	"time"
)

// AccountTable 账户表
type AccountTable struct {
	//
	ID       int64  `gorm:"primarykey" json:"id"`
	Name     string `gorm:"column:name;type:varchar(120);common:用户名" json:"name"`
	Phone    string `gorm:"column:phone;type:varchar(120);comment:手机号" json:"phone"`
	PassWord string `gorm:"column:pass_word;type:varchar(120);comment:密码" json:"passWord"`
	//
	RoleID int64 `gorm:"column:role_id;type:int8;comment:角色id" json:"roleID"`
	//
	CreatedAt time.Time `gorm:"created_at" json:"-"`
	UpdatedAt time.Time `gorm:"updated_at" json:"-"`
}

// TableName
func (t *AccountTable) TableName() string {
	return "t_account"
}
