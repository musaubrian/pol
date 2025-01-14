package main

import (
	"testing"
)

func TestRPNParse(t *testing.T) {
	cases := []struct {
		expr string
		want Stack
	}{
		{
			expr: "5 4 +",
			want: Stack{Expressions: []Expr{
				{FirstVal: 5, SecondVal: 4, Operation: "+"},
			},
			},
		},
		{
			expr: "(54+) (23-) +",
			want: Stack{Expressions: []Expr{
				{FirstVal: 9, SecondVal: -1, Operation: "+"},
			},
			},
		},
		{
			expr: "(54+) (23-) + (16*) -",
			want: Stack{Expressions: []Expr{
				{FirstVal: 10, SecondVal: 6, Operation: "-"},
			},
			},
		},
	}
	for _, c := range cases {
		st, err := Parse(c.expr)
		if err != nil {
			t.Fatalf("Expected no error: Found: %v", err)
		}
		if st == nil {
			t.Fatalf("Expected none empty stack, got %v", st)
		}
		if len(st.Expressions) != len(c.want.Expressions) {
			t.Errorf("Expected %d expressions, got: %d", len(c.want.Expressions), len(st.Expressions))
		}
		if st.Expressions[0].FirstVal != c.want.Expressions[0].FirstVal {
			t.Errorf("Expected similar first values got: %f %f", st.Expressions[0].FirstVal, c.want.Expressions[0].FirstVal)
		}
		if st.Expressions[0].SecondVal != c.want.Expressions[0].SecondVal {
			t.Errorf("Expected similar second values got: %f %f", st.Expressions[0].SecondVal, c.want.Expressions[0].SecondVal)
		}
	}

}

func TestRPNCalc(t *testing.T) {

	cases := []struct {
		expr string
		want float64
	}{
		{expr: "5 9 +", want: 14},
		{expr: "15 3 /", want: 5},
		{expr: "2 2 *", want: 4},
		{expr: "2 2 ^", want: 4},
		{expr: "4 2 ^", want: 16},
		{expr: "(54+) (23-) +", want: 8},
		{expr: "(42/) (23-) *", want: -2},
		{expr: "(22^) (50*) * (51+) -", want: -6},
	}

	for _, c := range cases {
		st, _ := Parse(c.expr)
		for _, expr := range st.Expressions {
			res := Calc(expr)
			if res != c.want {
				t.Errorf("Expected %.3f, Got %.3f", res, c.want)
			}
		}

	}
}
