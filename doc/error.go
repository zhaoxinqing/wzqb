package doc

import (
	"errors"
	"fmt"
)

// panic
// 我们给出异常处理的作用域（场景）：

// 空指针引用
// 下标越界
// 除数为0
// 不应该出现的分支，比如default
// 输入不应该引起函数错误

func UserErrorsNew() {
	err := errors.New("检查：新建错误提示")
	fmt.Println(err)
}
