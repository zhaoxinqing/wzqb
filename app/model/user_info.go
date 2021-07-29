package model

import "time"

// Users 后台用户
type UserInfo struct {
	ID        int64     `gorm:"id" json:"id"`           //
	UserID    int64     `gorm:"user_id" json:"user_id"` // 用户id
	Info      string    `gorm:"info" json:"info"`       // 信息
	CreatedAt time.Time `gorm:"created_at" json:"-"`    // 创建时间
	UpdatedAt time.Time `gorm:"updated_at" json:"-"`    // 更新时间
}

// TableName 表名
func (UserInfo) TableName() string {
	return "user_info"
}
