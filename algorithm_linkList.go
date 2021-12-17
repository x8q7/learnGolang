package main

import "fmt"

// 队列
type Queue struct {
	head  int
	tail  int
	queue [4]int
	max   int
}

// 添加
func (q *Queue) Push(val int) {
	// 是否 满了
	if q.IsFull() {
		fmt.Println("Queue is full")
		return
	}
	defer q.AfterPush()
	q.queue[q.tail] = val
}
func (q *Queue) AfterPush() {
	q.tail = (q.tail + 1) % q.max
}

// 取出
func (q *Queue) Unshift() (res int) {
	// 是否为空
	if q.IsEmpty() {
		fmt.Println("Queue is Empty")
		return
	}
	defer q.AfterUnshift()
	return q.queue[q.head]
}
func (q *Queue) AfterUnshift() {
	q.queue[q.head] = 0
	q.head = (q.head + 1) % q.max
}

// 是否满 [0,0,1,2,0]
func (q *Queue) IsFull() (res bool) {
	return (q.tail+1)%q.max == q.head
}

// 是否为空
func (q *Queue) IsEmpty() (res bool) {
	return q.head == q.tail
}

// 打印
func (q *Queue) PrintQ(i string) {
	fmt.Println(i, q.queue)
}

func main() {
	que := Queue{
		queue: [4]int{},
		head:  0,
		tail:  0,
		max:   4,
	}

	que.Push(1)
	que.Push(2)
	que.Push(3)
	que.PrintQ("11111111")
	que.Unshift()
	que.PrintQ("22222222")
	que.Push(4)
	que.PrintQ("33333333")
	que.Unshift()
	que.PrintQ("44444444")
	que.Push(5)
	que.PrintQ("55555555")
	que.Unshift()
	que.PrintQ("66666666")
	que.Unshift()
}
