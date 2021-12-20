package main

import (
	"fmt"
)

// 单链表
type LinkLine struct {
	data map[string]string
	next *LinkLine
}

// 追加 节点
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

// 插入 节点
func InsertData(target *LinkLine, data *LinkLine) {
	temp := target.next

	target.next = data
	target.next.next = temp
}

// 删除 节点
func DeleteNode(head *LinkLine, condition map[string]string) {
	temp := head
	isMatched := false
	for {
		isMatched = IsMatch(temp.next, condition)

		if isMatched {
			break
		}
		if temp.next.next != nil {
			temp = temp.next
		} else {
			break
		}
	}

	if isMatched {
		temp.next = temp.next.next
	}
}

// 判断 条件
func IsMatch(node *LinkLine, condition map[string]string) (res bool) {
	isMatched := true
	for i, v := range condition {
		rVal, ok := node.data[i] //判断 temp 的下一个节点
		if !ok {
			isMatched = false
			break
		}
		if rVal != v {
			isMatched = false
			break
		}
	}
	return isMatched
}

// 输出 所有节点
func (head *LinkLine) PrintLink() {
	tmp := head
	for {
		fmt.Println(tmp.data)
		if tmp.next == nil {
			fmt.Println("-------------------")
			return
		}
		tmp = tmp.next
	}
}

func main() {
	head := LinkLine{
		data: map[string]string{"name": "aaa", "age": "18"},
	}
	// 添加 节点
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

	// 插入节点
	InsertData(head.next, &LinkLine{
		data: map[string]string{"name": "hh", "age": "8"},
	})

	head.PrintLink()

	// 删除节点
	DeleteNode(&head, map[string]string{"name": "cc"})

	head.PrintLink()
}
