package models

import "time"

// Users 后台用户
type Users struct {
	ID        int64     `gorm:"column:id;primary_key;not null;" json:"id"`                     // 账户系统userID
	Name      string    `gorm:"column:name;" json:"name" `                                     //备注
	Status    bool      `gorm:"column:status;type:bool;not null;" json:"status" form:"status"` // 状态(1:正常 2:未激活 3:暂停使用)
	GroupID   int64     `gorm:"column:group_id;index:group_id;not null;" json:"groupID" form:"groupID"`
	IsManager bool      `gorm:"column:is_manager;type:bool;not null;" json:"isManager" form:"isManager"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"` // 更新时间
}

// TableName 表名
func (Users) TableName() string {
	return "sys_users"
}
