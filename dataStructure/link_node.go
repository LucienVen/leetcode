// @Author : liangliangtoo
// @Time : 2021/8/12 5:10 下午 
// @Describe :  链表实现

package dataStructure

import "fmt"

type LinkNode struct {
	Data     interface{}
	NextNode *LinkNode
}

type LinkNodeInt struct {
	Data     int
	NextNode *LinkNodeInt
}

// test 生成链表结构
func GetNewLinkInt() *LinkNodeInt {
	//4,5,1,9
	head := new(LinkNodeInt)
	head.Data = 4

	node1 := new(LinkNodeInt)
	node1.Data = 5
	head.NextNode = node1

	node2 := new(LinkNodeInt)
	node2.Data = 1
	node1.NextNode = node2

	node3 := new(LinkNodeInt)
	node3.Data = 9
	node2.NextNode = node3

	return head
}

// 顺序遍历链表
func PrintLink(head *LinkNodeInt)  {
	per := head
	if per != nil {
		fmt.Println("node: ", per.Data)
		per = per.NextNode
	}
}


// 链表长度
func GetLinkLength(head *LinkNodeInt) int {
	length := 0
	per := head
	for per != nil {
		length += 1
		per = per.NextNode
	}

	return length
}

// 删除链表的倒数第 n 个结点，并且返回链表的头结点。
// head = [1,2,3,4,5], n = 2
// [1,2,3,5]
func RemoveNthFromEnd(head *LinkNodeInt, n int) *LinkNodeInt {
	//获取链表长度
	linkLen := GetLinkLength(head)
	if linkLen < n {
		return head
	}

	last := linkLen - n

	perNode := head
	// 找出删除节点的前一个节点
	for i := 0; i < last - 1; i++ {
		perNode = perNode.NextNode
	}

	fmt.Println("perNode: ", perNode)




	return nil
}