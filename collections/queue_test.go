package collections

import (
	"testing"
)

func TestNew(t *testing.T){
	q := New[int]()
	if q == nil{
		t.Fatalf("expected : %v , got : %v\n",&Queue[any]{}, q)
	}
	if q.items == nil{
		t.Fatalf("expected : %v , got : %v\n",[]any{}, q.items)	
	}
	if q.head != -1{
		t.Fatalf("expected : -1, got : %v\n",q.head)
	}
}


func TestEnqueue(t *testing.T){
	q := New[int]()
	for i := range 10{
		q.Enqueue(i+1)
	}

	assert(t, len(q.items) != 10, 10, len(q.items))

	for range 5{
		q.Dequeue()
	}
	for i := range 7{
		q.Enqueue(i + 20)
	}
}


func assert(t *testing.T, c bool, e any, g any){
	if c{
		t.Fatalf("expected : %v, got : %v\n",e,g)
	}
}
