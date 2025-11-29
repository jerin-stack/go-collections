package collections

import (
	"errors"
)

type Queue[T any] struct{
	head int
	tail int
	items []T
	size int
}

func New[T any]() *Queue[T]{
	size := 8
	return &Queue[T]{
		head: 0,
		tail: -1,
		size: size,
		items: make([]T, size),
	}
}

func (q *Queue[T])Enqueue(d T){
	q.tail++
	if q.tail >= q.size{
		if t := q.tail % q.size; t < q.head{
			for i := range (q.size-q.head){
				q.items[i] = q.items[q.head+i]
			}
			q.tail = q.tail- q.head
			q.head = 0
		} else{
			// adding more space in the end
			l := int(q.size/2)
			q.items = append(q.items, make([]T, l)...)
			q.size += l
			if q.head > 0{
				for i:=0; i<q.size; i++{
					q.items[q.tail+i] = q.items[i]
					q.items[i] = q.items[q.head+i]
				}
				q.head = 0
			}
		}
	}
	q.items[q.tail] = d
}

func (q *Queue[T])Dequeue() (T, error){
	if q.head <= q.tail{
		d := q.items[q.head]
		q.head++
		return d, nil
	}
	var z T
	return z, errors.New("no elements in queue")
}

func (q *Queue[T])Peek()(T, error){
	if q.head == q.tail{
		var z T
		return z, errors.New("no elements in queue")
	}
	return q.items[q.tail], nil
}
