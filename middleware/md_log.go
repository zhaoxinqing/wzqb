package middleware

import (
	"Kilroy/util"
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
			ID:            uuid.New().String(),
			Name:          "日志列表",
			RequestURL:    r.URL.String(),
			RequestType:   r.Method,
			IP:            r.Host,
			OperatingTime: util.Time(time.Now()),
			CreateTime:    util.Time(time.Now()),
			UpdateTime:    util.Time(time.Now()),
		}
		m.BussLog = &BussLog
		fmt.Println(BussLog)
		// Passthrough to next handler if need
		next(w, r)
	}
}

// Business_log struct is a row record of the t_business_log table in the gemdale database
type Business_log struct {
	ID            string    `gorm:"primary_key;column:id;type:VARCHAR;" json:"id"`
	Name          string    `gorm:"column:name;type:VARCHAR;" json:"name"`
	RequestURL    string    `gorm:"column:request_url;type:VARCHAR;" json:"requestUrl"`
	RequestType   string    `gorm:"column:request_type;type:VARCHAR;" json:"requestType"`
	RequestParam  string    `gorm:"column:request_param;type:VARCHAR;" json:"requestParam"` // logic
	Code          string    `gorm:"column:code;type:VARCHAR;" json:"code"`                  // 返回时
	CodeDesc      string    `gorm:"column:code_desc;type:VARCHAR;" json:"codeDesc"`         // 返回时
	Username      string    `gorm:"column:username;type:VARCHAR;" json:"username"`          // logic
	UID           string    `gorm:"column:uid;type:VARCHAR;" json:"uid"`                    // logic
	IP            string    `gorm:"column:ip;type:VARCHAR;" json:"ip"`
	CostTime      int64     `gorm:"column:cost_time;type:INT4;" json:"costTime"`
	Message       string    `gorm:"column:message;type:VARCHAR;" json:"message"`
	Del           bool      `gorm:"column:del;type:BOOL;" json:"del"`
	Content       string    `gorm:"column:content;type:VARCHAR;size:255;" json:"content"`
	IPInfo        string    `gorm:"column:ip_info;type:VARCHAR;size:255;" json:"ipInfo"`
	Module        string    `gorm:"column:module;type:VARCHAR;size:255;" json:"module"`
	Operating     string    `gorm:"column:operating;type:VARCHAR;size:255;" json:"operating"`
	OperatingTime util.Time `gorm:"column:operating_time;type:TIMESTAMP;" json:"operatingTime"`
	CreateTime    util.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"createTime"`
	UpdateTime    util.Time `gorm:"column:update_time;type:TIMESTAMP;" json:"updateTime"`
}
