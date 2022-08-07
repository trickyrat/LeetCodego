package main

import (
	"strconv"
	"strings"
)

const mask1, mask2 = 1 << 7, 1<<7 | 1<<6

func createListNode(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	dummyHead := head
	for i := 1; i < len(nums); i++ {
		dummyHead.Next = &ListNode{Val: nums[i]}
		dummyHead = dummyHead.Next
	}
	return head
}

func createTreeNodeWithBFS(data string) *TreeNode {
	sp := strings.Split(data, ",")
	if sp[0] == "null" {
		return nil
	}
	val, _ := strconv.Atoi(sp[0])
	root := &TreeNode{val, nil, nil}
	var queue []*TreeNode
	queue = append(queue, root)
	index := 1
	for index < len(sp) {
		var node = queue[0]
		queue = queue[1:]
		leftStr := sp[index]
		rightStr := sp[index+1]
		if leftStr != "null" {
			leftVal, _ := strconv.Atoi(leftStr)
			leftNode := &TreeNode{leftVal, nil, nil}
			if node != nil {
				node.Left = leftNode
			}
			queue = append(queue, leftNode)
		}
		if rightStr != "null" {
			rightVal, _ := strconv.Atoi(rightStr)
			rightNode := &TreeNode{rightVal, nil, nil}
			if node != nil {
				node.Right = rightNode
			}
			queue = append(queue, rightNode)
		}
		index += 2
	}
	return root
}

func createTreeNodeWithDFS(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}

func (h *ListNode) toString() string {
	var res = "["
	res += strconv.Itoa(h.Val)
	for h.Next != nil {
		res += strconv.Itoa(h.Val)
		h = h.Next
	}
	res += "]"
	return res
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func isSelfDividing(num int) bool {
	for x := num; x > 0; x /= 10 {
		if d := x % 10; d == 0 || num%d != 0 {
			return false
		}
	}
	return true
}

func parseComplexNumber(num string) (real, imag int) {
	i := strings.IndexByte(num, '+')
	real, _ = strconv.Atoi(num[:i])
	imag, _ = strconv.Atoi(num[i+1 : len(num)-1])
	return
}

func getBytes(num int) int {
	if num&mask1 == 0 {
		return 1
	}
	n := 0
	for mask := mask1; num&mask != 0; mask >>= 1 {
		n++
		if n > 4 {
			return -1
		}
	}
	if n >= 2 {
		return n
	}
	return -1
}
