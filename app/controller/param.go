package controller

type OrmTest struct{}

type PostParam struct {
	Name      string // 名称
	Age       int64  // 年龄
	Operation string // 操作
}

type GetByTimeParam struct {
	Time1 string `json:"time1"`
	Time2 string `json:"time2"`
}
