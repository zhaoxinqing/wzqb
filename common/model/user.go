package model

import "time"

// Users 后台用户
type User struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	UserName  string    `gorm:"column:user_name;type:varchar(100)" json:"userName"` // 用户名
	Phone     string    `gorm:"column:phone;type:varchar(100)" json:"phone"`        // 手机号
	PassWord  string    `gorm:"column:pass_word;type:varchar(100)" json:"passWord"` // 密码
	Remark    string    `gorm:"remark" json:"remark"`                               // 备注
	CreatedAt time.Time `gorm:"created_at" json:"-"`                                //
	UpdatedAt time.Time `gorm:"updated_at" json:"-"`                                //
}

// TableName 表名
func (User) TableName() string {
	return "user"
}
