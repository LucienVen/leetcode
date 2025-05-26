package linked_list

// https://leetcode.cn/leetbook/read/linked-list/jy291/
// 设计链表

// 输入
//["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
//[[], [1], [3], [1, 2], [1], [1], [1]]
//输出
//[null, null, null, null, 2, null, 3]
//
//解释
//MyLinkedList myLinkedList = new MyLinkedList();
//myLinkedList.addAtHead(1);
//myLinkedList.addAtTail(3);
//myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
//myLinkedList.get(1);              // 返回 2
//myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
//myLinkedList.get(1);              // 返回 3

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	// 取第一个节点
	res := this

	for i := 0; i <= index; i++ {

		if res.next == nil {
			return -1
		}

		res = res.next
	}

	return res.val
}

// void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &MyLinkedList{val: val}
	newNode.next = this
}

// void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &MyLinkedList{val: val}

	tail := this // 获取尾节点
	for tail.next != nil {
		tail = tail.next
	}

	tail.next = newNode
}

// void addAtIndex(int index, int val)
// 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}

	if index == 0 {
		this.AddAtHead(index)
		return
	}

	newNode := &MyLinkedList{val: val}

	prev := &MyLinkedList{}
	current := this
	linkIndex := 0

	for current.next != nil {

		if index == linkIndex {
			// 追加, 插入到链表中下标为 index 的节点之前
			prev.next = newNode
			newNode.next = current

			return

		} else {
			prev = current

			current = current.next

			linkIndex++
		}
	}

}

// void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	prev := &MyLinkedList{}
	current := this

	currentIndex := 0

	for current.next != nil {
		if currentIndex == index {
			if currentIndex == 0 { // 首节点
				this = current.next

			} else {
				// 非首节点
				prev.next = current.next
			}
		} else {
			prev = current
			current = current.next
			currentIndex++
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
