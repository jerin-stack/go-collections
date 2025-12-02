package tree 

import (
	"testing"
)

func TestInsert(t *testing.T){
	bi := &BinaryTree{}
	insertDataToQueue(bi)
	q := bi.InOrderTraversal()

	iotr := []int{1,2,3,4,4,5,5,6,6}

	for _, i := range iotr{
		v, er := q.Dequeue()
		if er != nil{
			t.Fatalf("error while reading queue\n")
		}
		if v != i{
			t.Fatalf("excepted: %d,got: %d\n",i,v)
		}
	}
}

func TestInOrder(t *testing.T){
	bi := &BinaryTree{}
	insertDataToQueue(bi)
	q := bi.InOrderTraversal()

	iotr := []int{1,2,3,4,4,5,5,6,6}

	for _, i := range iotr{
		v, er := q.Dequeue()
		if er != nil{
			t.Fatalf("error while reading queue\n")
		}
		if v != i{
			t.Fatalf("excepted: %d,got: %d\n",i,v)
		}
	}
}

func TestPreOrder(t *testing.T){
	bi := &BinaryTree{}
	insertDataToQueue(bi)
	q := bi.PreOrderTraversal()

	iotr := []int{1, 2, 5, 3, 4, 4, 5, 6, 6}

	for _, i := range iotr{
		v, er := q.Dequeue()
		if er != nil{
			t.Fatalf("error while reading queue\n")
		}
		if v != i{
			t.Fatalf("excepted: %d,got: %d\n",i,v)
		}
	}
}

func TestPostOrder(t *testing.T){
	bi := &BinaryTree{}
	insertDataToQueue(bi)
	q := bi.PostOrderTraversal()

	iotr := []int{4, 5, 4, 3, 6, 6, 5, 2, 1}

	for _, i := range iotr{
		v, er := q.Dequeue()
		if er != nil{
			t.Fatalf("error while reading queue\n")
		}
		if v != i{
			t.Fatalf("excepted: %d,got: %d\n",i,v)
		}
	}
}

func insertDataToQueue(bi *BinaryTree){
	bi.Insert(1)
	bi.Insert(2)
	bi.Insert(5)
	bi.Insert(6)
	bi.Insert(3)
	bi.Insert(4)
	bi.Insert(4)
	bi.Insert(5)
	bi.Insert(6)
}
