package list

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered]struct{
	data T
	next *node[T]
}
type LinkedList[T constraints.Ordered] struct{
	head *node[T]
}

func (l *LinkedList[T])Add(d T){
	n := &node[T]{
			data: d,
		}
	if l.head == nil{
		l.head = n
		return
	}
	lastNode := l.head
	l.head = n
	n.next = lastNode
}

func (l *LinkedList[T])AddAt(i int, d T) error{
	n := l.head
	for j := range i+1{
		if n == nil{
			return fmt.Errorf("can't insert at %d, list size is %d",i, j)
		}
		n = n.next
	}
	nx := n.next
	n.next = &node[T]{
		data: d,
		next: nx,
	}
	return nil
}

func (l *LinkedList[T])RemovedFrom(i int) error{
	if l.head == nil{
			return fmt.Errorf("can't insert at %d, list is empty",i)
	}
	n := l.head
	for j := range i{
		if n.next == nil{
			return fmt.Errorf("can't insert at %d, list size is %d",i, j)
		}
		n = n.next
	}

	x := n.next.next
	n.next = x
	return nil
}

func (l *LinkedList[T])GetIndex(d T) (int, error){
	for i, n:= 1, l.head; n.next != nil; n = n.next {
		i++
		if n.data == d{
			return i, nil
		}
	}
	return -1, errors.New("can't find item in the list")
}
