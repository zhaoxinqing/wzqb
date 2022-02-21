package model

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

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

// 使用参数化查询，禁止拼接SQL语句，另外对于传入参数用于order by或表名的需要通过校验
// bad
func Handler(db *sql.DB, req *http.Request) {
	q := fmt.Sprintf("SELECT ITEM,PRICE FROM PRODUCT WHERE ITEM_CATEGORY='%s' ORDER BY PRICE", req.URL.Query()["category"])
	db.Query(q)
}

// good
func HandlerGood(db *sql.DB, req *http.Request) {
	// 使用?占位符
	q := "SELECT ITEM,PRICE FROM PRODUCT WHERE ITEM_CATEGORY='?' ORDER BY PRICE"
	db.Query(q, req.URL.Query()["category"])
}
