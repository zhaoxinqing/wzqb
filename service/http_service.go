package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// AddUserResult ...
type ResultParam struct {
	Data interface{}
	Code int64
	Msg  string
}

func GetHttpServe(method string, requestURL string, token string, param interface{}) (result ResultParam, err error) {
	var (
		client = &http.Client{}
	)

	info, err := json.Marshal(param)
	req, err := http.NewRequest(method, requestURL, strings.NewReader(string(info)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyContent, &result)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf("转发请求失败")
		return
	}
	return
}
