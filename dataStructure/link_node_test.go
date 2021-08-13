// @Author : a1234
// @Time : 2021/8/12 5:16 下午 
// @Describe :  

package dataStructure

import (
	"fmt"
	"testing"
)

func TestGetNewLinkInt(t *testing.T) {
	head := GetNewLinkInt()

	if head == nil {
		return
	}

	// 遍历链表
	headNode := head
	for {
		fmt.Println("data: ", headNode.Data)

		if headNode.NextNode != nil {
			headNode = headNode.NextNode
		} else {
			break
		}
	}

	fmt.Println("链表长度：", GetLinkLength(head))


	fmt.Println("打印链表")
	PrintLink(head)
}

