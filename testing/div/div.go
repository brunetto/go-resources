package div

import "github.com/pkg/errors"

// not a serious function, just an example
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.Errorf("can0t divide by zero")
	}

	return a / b, nil
}
