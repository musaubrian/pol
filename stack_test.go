package main

import "testing"

func TestStackNew(t *testing.T) {
	st := NewStack()

	if st == nil {
		t.Errorf("Expected an empty stack got <nil>")
	}

	if len(st.Expressions) != 0 {
		t.Errorf("Expected an empty stack got %d", len(st.Expressions))
	}

}

func TestStackPush(t *testing.T) {
	st := NewStack()

	exprs := []Expr{
		{FirstVal: 2, SecondVal: 2, Operation: "/"},
		{FirstVal: 5, SecondVal: 2, Operation: "*"},
	}

	for _, v := range exprs {
		st.Push(v)
	}

	if len(st.Expressions) != len(exprs) {
		t.Errorf("Expected One value in stack, Got %d", len(st.Expressions))
	}
}

func TestStackPop(t *testing.T) {
	st := NewStack()

	exprs := []Expr{
		{FirstVal: 2, SecondVal: 2, Operation: "/"},
		{FirstVal: 5, SecondVal: 2, Operation: "*"},
	}

	for _, v := range exprs {
		st.Push(v)
	}

	prevCount := len(st.Expressions)

	st.Pop()

	if len(st.Expressions) != 1 {
		t.Errorf("Expected one item in stack: got %d", len(st.Expressions))
	}
	if len(st.Expressions) >= prevCount {
		t.Errorf("Expected: %d: Got %d", prevCount, len(st.Expressions))
	}
}
