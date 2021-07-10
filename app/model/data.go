package model

import "time"

// Users 后台用户
type UserInfo struct {
	ID        int64     `gorm:"id" json:"id"`                 // ID（主键）
	UserID    int64     `gorm:"user_id" json:"UserID"`        // 用户id
	Avatar    string    `gorm:"avatar" json:"avatar"`         // 用户头像
	Data      string    `gorm:"data" json:"data"`             // 数据
	CreatedAt time.Time `gorm:"created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"` // 更新时间
}

// TableName 表名
func (UserInfo) TableName() string {
	return "data"
}
