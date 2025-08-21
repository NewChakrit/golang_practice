package main

import "fmt"

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.

กำหนด head ซึ่งเป็นส่วนหัวของ linked list ให้พิจารณาว่า linked list มีี circle หรือไม่
linked list จะมี circle หากมี node ใด node หนึ่งใน list ที่สามารถเข้าถึงได้อีกครั้งโดยการติดตาม pointer ถัดไปอย่างต่อเนื่อง ฟังก์ชัน pos ภายในใช้ระบุ index ของ node ที่ pointer ถัดไปของส่วนหางเชื่อมต่ออยู่ด้วย โปรดทราบว่า pos ไม่ได้ถูกส่งเป็นพารามิเตอร์
คืนค่า true หากมี circle ใน linked list มิฉะนั้น ให้คืนค่า false

Concept: Linked List Cycle Detection
Fast & Slow Pointer (Floyd’s Cycle Detection)
มี pointer 2 ตัว:
slow เดินทีละ 1 ก้าว
fast เดินทีละ 2 ก้าว
ถ้ามี cycle → สุดท้าย fast กับ slow จะเจอกัน
ถ้าไม่มี cycle → fast จะเจอ nil
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	// สร้าง cycle ตาม pos = 1
	node4.Next = node2

	fmt.Println(hasCycle(node1))
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
