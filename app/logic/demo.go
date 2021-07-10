package logic

import (
	"Kilroy/app/auth"
	"fmt"
)

func Demo(a auth.A) (result interface{}, err error) {
	err = fmt.Errorf("hello error")
	return a, err
}
