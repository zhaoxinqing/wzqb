package model

import "time"

// Users 后台用户
type User struct {
	ID        int64     `gorm:"id" json:"id"`         //
	Name      string    `gorm:"name" json:"name"`     // 姓名
	Age       int64     `gorm:"age" json:"age"`       // 年龄
	Sex       string    `gorm:"sex" json:"sex"`       // 性别：0女、1男、2未知
	Remark    string    `gorm:"remark" json:"remark"` // 备注
	CreatedAt time.Time `gorm:"created_at" json:"-"`  //
	UpdatedAt time.Time `gorm:"updated_at" json:"-"`  //
}

// TableName 表名
func (User) TableName() string {
	return "user"
}
