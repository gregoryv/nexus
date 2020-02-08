package nexus

import (
	"fmt"
	"testing"
)

func Test_First(t *testing.T) {
	var r int
	ok, _k := assert(t)
	ok(
		First(
			double(&r, 2), // 4
			double(&r, r), // 8
		),
	)
	_k(
		First(
			double(&r, 1),
			double(nil, 0), // fails
			double(&r, 2),
		),
	)
}

func double(r *int, n int) error {
	if n < 0 {
		return fmt.Errorf("double: n < 0")
	}
	if r == nil {
		return fmt.Errorf("double: r is nil")
	}
	*r = 2 * n
	return nil
}
