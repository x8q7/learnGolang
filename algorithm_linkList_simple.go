package main

import "fmt"

// 单链表
type LinkLine struct {
	data map[string]string
	next *LinkLine
}

func AddData(headNode *LinkLine, data *LinkLine) {
	temp := headNode
	for {
		if temp.next == nil {
			// temp.next = data  循环停止后的 temp就是 最后一个
			break
		}
		temp = temp.next
	}

	temp.next = data
}

func (head *LinkLine) PrintLink() {
	tmp := head
	for {
		fmt.Println(tmp.data)
		if tmp.next == nil {
			return
		}
		tmp = tmp.next
	}
}

func main() {
	head := LinkLine{
		data: map[string]string{"name": "aaa", "age": "18"},
	}

	AddData(&head, &LinkLine{
		data: map[string]string{"name": "bbb", "age": "18"},
	})
	AddData(&head, &LinkLine{
		data: map[string]string{"name": "cc", "age": "12"},
	})
	AddData(&head, &LinkLine{
		data: map[string]string{"name": "ff", "age": "11"},
	})

	head.PrintLink()
}
