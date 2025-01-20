package main

type Expr struct {
	Raw string

	FirstVal  float64
	SecondVal float64
	Operation string

	Result float64
	Err    error
}

type Stack struct {
	Expressions []Expr
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(expr Expr) {
	s.Expressions = append(s.Expressions, expr)
}

func (s *Stack) Pop() Expr {
	if !s.isValid() {
		return Expr{}
	}

	last := s.Expressions[len(s.Expressions)-1]

	s.Expressions = s.Expressions[:len(s.Expressions)-1]

	return last
}

func (s *Stack) isValid() bool {
	if len(s.Expressions) > 0 {
		return true
	}
	return false
}
