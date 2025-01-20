package main

import (
	"testing"
)

func TestRPNEval(t *testing.T) {
	cases := []struct {
		expr string
		want Expr
	}{
		{
			expr: "5 4 +",
			want: Expr{FirstVal: 5, SecondVal: 4, Operation: "+", Raw: "5 4 +", Result: 9},
		},
		{
			expr: "(7 9 -) (4 6 ^) -",
			want: Expr{FirstVal: -2, SecondVal: 4096, Operation: "-", Raw: "(7 9 -) (4 6 ^) -", Result: -4098},
		},
		{
			expr: "((4 2 -) (4 2 ^) -) (5 7 +) -",
			want: Expr{FirstVal: -2, SecondVal: 4096, Operation: "-", Raw: "(7 9 -) (4 6 ^) -", Result: -4098},
		},
	}
	for _, c := range cases {
		evald := Eval(c.expr)

		if c.want.Result != evald.Result {
			t.Errorf("Expected Result: %f, Got: %f", c.want.Result, evald.Result)
		}
		if c.want.Raw != evald.Raw {
			t.Errorf("Expected expression: [%s], Got: [%s]", c.want.Raw, evald.Raw)
		}
	}

}
