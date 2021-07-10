package model

import "time"

// Users 后台用户
type Users struct {
	ID        int64     `gorm:"id" json:"id"`                 // ID（主键）
	Statue    int64     `gorm:"statue" json:"statue"`         // 用户状态（1:正常 2:未激活 3:暂停使用）
	Name      string    `gorm:"name" json:"name"`             // 用户名
	Phone     string    `gorm:"phone" json:"phone"`           // 用户手机号
	CreatedAt time.Time `gorm:"created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"` // 更新时间
}

// TableName 表名
func (Users) TableName() string {
	return "user"
}
