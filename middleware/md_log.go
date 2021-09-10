package middleware

import (
	"fmt"

	"net/http"
	"time"

	"github.com/google/uuid"
)

type LogAddMiddleware struct {
	BussLog *Business_log
}

func NewLogAddMiddleware() *LogAddMiddleware {
	return &LogAddMiddleware{}
}

func (m *LogAddMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var BussLog = Business_log{
			ID:          uuid.New().String(),
			Name:        "日志列表",
			RequestURL:  r.URL.String(),
			RequestType: r.Method,
			IP:          r.Host,
			CreateTime:  time.Now(),
			UpdateTime:  time.Now(),
		}
		m.BussLog = &BussLog
		fmt.Println(BussLog)
		// Passthrough to next handler if need
		next(w, r)
	}
}

type Business_log struct {
	ID           string    `gorm:"primary_key;column:id;type:VARCHAR;" json:"id"`
	Name         string    `gorm:"column:name;type:VARCHAR;" json:"name"`
	IP           string    `gorm:"column:ip;type:VARCHAR;" json:"ip"`
	RequestURL   string    `gorm:"column:request_url;type:VARCHAR;" json:"requestUrl"`
	RequestType  string    `gorm:"column:request_type;type:VARCHAR;" json:"requestType"`
	RequestParam string    `gorm:"column:request_param;type:VARCHAR;" json:"requestParam"`
	CreateTime   time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time;type:TIMESTAMP;" json:"updateTime"`
}
