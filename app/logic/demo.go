package logic

import (
	"fmt"
)

func Demo() (result interface{}, err error) {
	err = fmt.Errorf("hello error")
	return
}
