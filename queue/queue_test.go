package queue

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T){
	q := New[int]()
	assert(t, q == nil, "queue", q)
	
}


func TestEnqueue(t *testing.T){
	q := New[int]()
	q.Enqueue(1)
	v, er := q.Dequeue()
	assert(t, er != nil, nil, er)
	assert(t, v != 1, 1, v)
	q.Enqueue(1)
	v, er = q.Dequeue()
	assert(t, er != nil, nil, er)
	assert(t, v != 1, 1, v)
	_, er = q.Dequeue()
	assert(t, er == nil, nil, errors.New(""))

	for i := range 12{
		q.Enqueue(i+1)
	}
	for range 8{
		q.Dequeue()
	}
	for i := range 12{
		q.Enqueue(i+20)
	}
	for range 15{
		q.Dequeue()
	}
	v, _ = q.Dequeue()
	assert(t, v != 31, 31, v)
	_, err := q.Dequeue()
	assert(t, err == nil, errors.New(""), err)
}


func assert(t *testing.T, c bool, e any, g any){
	if c{
		t.Fatalf("expected : %v, got : %v\n",e,g)
	}
}
