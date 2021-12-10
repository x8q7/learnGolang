package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	from := ""
	to := ""
	fmt.Println("请输入 要拷贝的文件!")
	fmt.Scanln(&from)
	fmt.Println("请输入 新文件的路径!")
	fmt.Scanln(&to)
	CopyFile(to, from)
}

func CopyFile(dstPath string, sourPath string) {
	fmt.Println("copy start!", dstPath, sourPath)
	sourse, err := os.Open(sourPath)

	if err != nil {
		fmt.Println(err)
	}

	defer sourse.Close()
	reader := bufio.NewReader(sourse)

	dest, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)

	if err != nil {
		fmt.Println(err)
	}

	defer dest.Close()
	writer := bufio.NewWriter(dest)

	writen, err := io.Copy(writer, reader)

	fmt.Println("ok", writen, err)
}
