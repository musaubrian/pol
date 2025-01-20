package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	ADD  = "+"
	SUB  = "-"
	DIV  = "/"
	MULT = "*"
	POW  = "^"
)

func Eval(expression string) *Expr {
	expression = strings.TrimSpace(expression)
	expr := &Expr{}

	if strings.HasPrefix(expression, "(") {
		return evalGroup(expression)
	}

	exprVals := []float64{}
	splitExpr := strings.Split(expression, " ")
	if len(splitExpr) < 3 {
		return &Expr{
			Err: errors.New("Invalid expression"),
		}
	}

	for _, char := range splitExpr {

		val, err := strconv.ParseFloat(char, 64)
		if err != nil {
			if isOperator(char) {
				expr.Raw = expression
				expr.FirstVal = exprVals[0]
				expr.SecondVal = exprVals[1]
				expr.Operation = char

				expr.calc()

				exprVals = []float64{}
				continue
			}

			return &Expr{
				Raw: expression,
				Err: errors.New("Invalid Expression"),
			}
		}
		exprVals = append(exprVals, val)
	}

	return expr
}

func evalGroup(expression string) *Expr {
	// group: (7 9 -) (4 6 ^) -
	// nested: ((4 2 -) (4 2 ^) -) (5 7 +) -
	exprVals := []float64{}
	currentVals := []float64{}
	e := &Expr{}

	for i, char := range expression {
		switch {
		case char == '(' || unicode.IsSpace(char):
			continue
		case char == ')':
			currentVals = append(currentVals, e.Result)
			expr := expression[i+2] // skip the space
			if isOperator(string(expr)) {
				g := &Expr{
					Raw:       expression,
					FirstVal:  currentVals[len(currentVals)-2],
					SecondVal: currentVals[len(currentVals)-1],
					Operation: string(expr),
				}
				g.calc()
				return g

			}

		case unicode.IsNumber(char):
			num, err := strconv.ParseFloat(string(char), 64)
			if err != nil {
				e.Err = fmt.Errorf("Failed to convert: %s", string(char))
			}
			exprVals = append(exprVals, num)
		case isOperator(string(char)) && len(exprVals) >= 2:
			e.FirstVal = exprVals[len(exprVals)-2]
			e.SecondVal = exprVals[len(exprVals)-1]
			e.Operation = string(char)
			e.calc()
		}
	}

	return e
}

func EvalFile(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	st := NewStack()
	for sc.Scan() {
		expr := sc.Text()
		evaluated := Eval(expr)
		st.Push(*evaluated)
	}
	if sc.Err() != nil {
		log.Error(sc.Err().Error())
	}

	totalExpresion := 0
	errordExpression := 0
	for _, ex := range st.Expressions {
		totalExpresion += 1
		if ex.Err != nil {
			printVals(ex.Raw, ex.Err.Error(), true)
			errordExpression += 1
			continue
		}
		printVals(ex.Raw, fmt.Sprint(ex.Result), false)
	}
	fmt.Printf(colorize(DARK, bold("\nEvaluated: (%d)\nFailed: (%d)\n")), totalExpresion, errordExpression)
}

func printVals(pre, post string, err bool) {
	sep := ":="
	postColor := DARK
	if err {
		postColor = REDISH
	}
	fmt.Printf("'%s' %s %s\n", colorize(DIM, bold(pre)), colorize(BLUEISH, bold(sep)), colorize(postColor, bold(post)))
}

func bold(v string) string {
	return fmt.Sprintf("\033[1m%s\033[21m", v)
}

func (e *Expr) calc() {
	var result float64

	switch e.Operation {
	case ADD:
		result = e.FirstVal + e.SecondVal
	case SUB:
		result = e.FirstVal - e.SecondVal
	case MULT:
		result = e.FirstVal * e.SecondVal
	case DIV:
		result = e.FirstVal / e.SecondVal
	case POW:
		result = math.Pow(e.FirstVal, e.SecondVal)
	default:
	}

	e.Result = result
}

func isOperator(v string) bool {
	switch v {
	case ADD, SUB, DIV, MULT, POW:
		return true
	default:
		return false
	}
}
