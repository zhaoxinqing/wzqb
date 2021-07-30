package doc

import "fmt"

func CommitMsg(msg string) {
	switch msg {
	case "新功能":
		fmt.Println("feat")
	case "修复bug":
		fmt.Println("fix")
	case "优化":
		fmt.Println("optimize")
	case "升级改造":
		fmt.Println("upgrade")
	case "格式":
		fmt.Println("style")
	case "小改动":
		fmt.Println("chore")
	case "重构":
		fmt.Println("refactor")
	case "提升性能":
		fmt.Println("perf")
	default:
		fmt.Println("")
	}
}

// : 格式（不影响代码运行的变动）
// perf: Performance的缩写，提升代码性能
// refactor：（既不是新增功能，也不是修改bug的代码变动）
// br: 此项特别针对bug号，用于向测试反馈bug列表的bug修改情况
// feat：新功能（feature）
// fix：修补
// docs：文档（documentation）
// style： 格式（不影响代码运行的变动）
// refactor：重构（即不是新增功能，也不是修改bug的代码变动）
// test：增加测试
// chore：其他的小改动. 一般为仅仅一两行的改动, 或者连续几次提交的小改动属于这种
// revert：feat(pencil): add 'graphiteWidth' option (撤销之前的commit)
// upgrade：升级改造
// bugfix：修补bug
// optimize：优化
// perf: Performance的缩写,
