package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("../hello/abc.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString("hi, world \n")
	}

	writer.Flush()
}
