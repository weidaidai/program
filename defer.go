//一般用于资源清理，
//解锁和关闭文件
//在函数中多次调用，先入后出，类似于栈的机制

package main

import (
	"fmt"
	"os"
)

func main() {

}

func readFile(filename string) {

	f1.err :=os.Open("float.go")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	btf := make([]byte, 1024)
	f1.Read(btf)

}
