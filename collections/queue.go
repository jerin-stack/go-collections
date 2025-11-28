package collections

import (
	"errors"
	"fmt"
)

type Queue[T any] struct{
	head int
	tail int
	items []T
}

func New[T any]() *Queue[T]{
	return &Queue[T]{
		head: -1,
		tail: -1,
		items: make([]T, 8),
	}
}

func (q *Queue[T])Enqueue(d T){
	space := 0
	if q.head <= q.tail{
		space = len(q.items) - (q.tail - q.head)
	} else{
		space = q.head - q.tail
	}
	
	if space > 0{
		if q.tail == len(q.items) -1{
			q.tail = 0
		}else{
			q.tail++
		}
		q.items[q.tail] = d
	}else{
		q.items = append(q.items, d)
		q.tail++
	}
	fmt.Println("head : ",q.head," tail : ",q.tail, " que : ", q.items, " space : ", space)
}

func (q *Queue[T])Dequeue() (T, error){
	if q.head == q.tail{
		var z T
		return z, errors.New("no elements in queue")
	}

	d := q.items[q.tail]
	if q.head == len(q.items)-1{
		q.head = 0
	}else{
		q.head++
	}
	return d, nil
}

func (q *Queue[T])Peek()(T, error){
	if q.head == q.tail{
		var z T
		return z, errors.New("no elements in queue")
	}
	return q.items[q.tail], nil
}
