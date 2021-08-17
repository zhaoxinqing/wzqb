package model

import "Kilroy/util"

// Users 后台用户
type User struct {
	ID        int64     `gorm:"id" json:"id"`         //
	Name      string    `gorm:"name" json:"name"`     // 姓名
	Age       int64     `gorm:"age" json:"age"`       // 年龄
	Sex       int64     `gorm:"sex" json:"sex"`       // 性别：0女、1男、2未知
	Remark    string    `gorm:"remark" json:"remark"` // 备注
	CreatedAt util.Time `gorm:"created_at" json:"-"`  //
	UpdatedAt util.Time `gorm:"updated_at" json:"-"`  //
}

// TableName 表名
func (User) TableName() string {
	return "user"
}
