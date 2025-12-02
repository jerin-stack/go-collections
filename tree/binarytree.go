package tree

import (
	"github.com/jerin-stack/go-collections/queue"
)

type node struct{
	data int
	left *node
	right *node
}

func (n *node)insert(data int){
	if n.data < data{
		if n.right == nil{
			nw := &node{
				data: data,
			}
			n.right = nw
		}else{
			n.right.insert(data)
		}
	}else{
		if n.left == nil{
			nw := &node{
				data: data,
			}
			n.left = nw
		}else{
			n.left.insert(data)
		}
	}
}

func (n *node)inOrderTraversal(q *queue.Queue[int]){
	if n == nil{
		return
	}
	n.left.inOrderTraversal(q)
	q.Enqueue(n.data)
	n.right.inOrderTraversal(q)
}

func (n *node)preOrderTraversal(q *queue.Queue[int]){
	if n == nil{
		return
	}
	q.Enqueue(n.data)
	n.left.preOrderTraversal(q)
	n.right.preOrderTraversal(q)
}

func (n *node)postOrderTraversal(q *queue.Queue[int]){
	if n == nil{
		return
	}
	n.left.postOrderTraversal(q)
	n.right.postOrderTraversal(q)
	q.Enqueue(n.data)
}
type BinaryTree struct{
	head *node
}

func (b *BinaryTree)Insert(data int){
	if b.head == nil{
		b.head = &node{
			data: data,
		}
	}else{
		b.head.insert(data)
	}
}

func (b *BinaryTree)InOrderTraversal() *queue.Queue[int]{
	q := queue.New[int]()
	b.head.inOrderTraversal(q)
	return q
}

func (b *BinaryTree)PreOrderTraversal() *queue.Queue[int]{
	q := queue.New[int]()
	b.head.preOrderTraversal(q)
	return q
}

func (b *BinaryTree)PostOrderTraversal() *queue.Queue[int]{
	q := queue.New[int]()
	b.head.postOrderTraversal(q)
	return q
}
