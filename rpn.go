package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	ADD  = "+"
	SUB  = "-"
	DIV  = "/"
	MULT = "*"
	POW  = "^"
)

func Parse(expression string) (*Stack, error) {
	exp := strings.Split(expression, " ")
	if len(exp) < 3 || !isOperand(exp[len(exp)-1]) {
		return nil, fmt.Errorf("Incomplete expression")
	}

	st := NewStack()

	vals := []float64{}
	for _, v := range exp {
		if strings.HasPrefix(v, "(") {
			res := parsenEvalGroup(v)
			vals = append(vals, res)
			continue
		}

		num, err := strconv.ParseFloat(v, 4)
		if err != nil {
			if isOperand(v) && len(vals) >= 2 {
				expr := &Expr{
					FirstVal:  vals[len(vals)-2],
					SecondVal: vals[len(vals)-1],
					Operation: v,
				}
				st.Push(*expr)
				vals = []float64{}
			}
			continue
		}
		vals = append(vals, float64(num))
	}

	return st, nil
}

// Workaround to allow working with parenthesis
func parsenEvalGroup(groupExpr string) float64 {
	expr := strings.Split(groupExpr, "")
	vals := []float64{}

	exp := &Expr{}
	for _, e := range expr {
		num, err := strconv.ParseFloat(e, 64)
		if err != nil {
			if isOperand(e) && len(vals) == 2 {
				exp.FirstVal = vals[0]
				exp.SecondVal = vals[1]
				exp.Operation = e
				vals = []float64{}
			}
			continue
		}
		vals = append(vals, float64(num))
	}
	return Calc(*exp)
}

func Calc(expr Expr) float64 {
	var result float64

	switch expr.Operation {
	case ADD:
		result = expr.FirstVal + expr.SecondVal
	case SUB:
		result = expr.FirstVal - expr.SecondVal
	case MULT:
		result = expr.FirstVal * expr.SecondVal
	case DIV:
		result = expr.FirstVal / expr.SecondVal
	case POW:
		result = math.Pow(expr.FirstVal, expr.SecondVal)
	default:
	}

	return result
}

func isOperand(v string) bool {
	switch v {
	case ADD, SUB, DIV, MULT, POW:
		return true
	default:
		return false
	}
}
