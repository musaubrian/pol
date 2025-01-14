package main

type Expr struct {
	FirstVal  float64
	SecondVal float64
	Operation string
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

func (s *Stack) Pop() {
	if !s.isValid() {
		return
	}

	if len(s.Expressions) == 2 {
		s.Expressions = s.Expressions[:len(s.Expressions)-1]
		return
	}

	s.Expressions = s.Expressions[:len(s.Expressions)-2]
}

func (s *Stack) Eval() []float64 {
	expResults := []float64{}
	for _, exp := range s.Expressions {
		val := Calc(exp)
		expResults = append(expResults, val)
	}

	return expResults
}

func (s *Stack) isValid() bool {
	if len(s.Expressions) > 0 {
		return true
	}
	return false
}
