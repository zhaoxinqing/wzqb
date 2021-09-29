package utils

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
	client := &http.Client{}
	info, err := json.Marshal(param)
	if err != nil {
		return
	}
	request, err := http.NewRequest(method, requestURL, strings.NewReader(string(info)))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", token)
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = fmt.Errorf("转发请求失败")
		return
	}
	return
}
