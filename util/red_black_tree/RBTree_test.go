package RBTree

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	a := []int{10, 40, 30, 60, 90, 70, 20, 50, 80}
	fmt.Println("原始数据：", a)
	rbtree := NewRBTree()
	for _, i := range a {
		fmt.Println("\n添加的数据：", i)
		rbtree.Insert(i)
		fmt.Println("\n前序遍历:")
		rbtree.root.PreOrder()
		fmt.Println("\n中序遍历:")
		rbtree.root.inOrder()
		fmt.Println("\n后序遍历:")
		rbtree.root.postOrder()
	}
	for i, _ := range a {
		i = a[8-i]
		fmt.Println("\n删除的数据：", i)
		rbtree.Remove(i)
		fmt.Println("\n前序遍历:")
		rbtree.root.PreOrder()
		fmt.Println("\n中序遍历:")
		rbtree.root.inOrder()
		fmt.Println("\n后序遍历:")
		rbtree.root.postOrder()
	}
}
