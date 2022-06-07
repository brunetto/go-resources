package div_test

import (
	"testing"

	"github.com/brunetto/go-resources/testing/div"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	t.Parallel() // this test can be ran in parallel with other tests

	// struct containing function arguments (optional)
	type args struct {
		a float64
		b float64
	}

	// test cases
	tests := []struct {
		name string
		args args
		want float64
		fail bool
	}{
		{name: "ok", args: args{a: 6, b: 3}, want: 2, fail: false},
		{name: "ok divide zero", args: args{a: 0, b: 3}, want: 0, fail: false},
		{name: "error dividing by zero", args: args{a: 6, b: 0}, want: 0, fail: true},
	}

	// run test cases
	for _, tt := range tests {
		// https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		// https://eli.thegreenplace.net/2019/go-internals-capturing-loop-variables-in-closures/
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // run test cases in parallel

			// call (exported) function
			got, err := div.Div(tt.args.a, tt.args.b)
			if tt.fail { // check failure, in case exit
				assert.NotNil(t, err)
				return
			}

			// test output
			assert.Equal(t, tt.want, got)
		})
	}
}
