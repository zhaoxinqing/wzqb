package tool

import (
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"
)

// 子进程回收
func ReapChildren() {
	c := make(chan os.Signal, 100)
	signal.Notify(c, unix.SIGCHLD)

	for {
		<-c
		var status unix.WaitStatus
		for {
			pid, err := unix.Wait4(-1, &status, unix.WNOHANG, nil)
			switch err {
			case nil:
				if pid > 0 {
					fmt.Println("Reap pid", pid)
				}
			case unix.ECHILD:
				// No more children, we are done.
				// break
			case unix.EINTR:
				continue
			default:
				fmt.Println(err)
			}
		}
	}
}
