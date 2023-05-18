package nexus

import (
	"fmt"
	"strings"
)

func NewStepper(err *error) *Stepper {
	return &Stepper{err}
}

type Stepper struct {
	err *error
}

// Stepf checks internal error before calling the next step. If error
// is set after step it's wrapped using the given format string, which
// must contain one optional %w. If format does not contain %w the
// format will be extended with ': %w' at the end.
func (me *Stepper) Stepf(format string, step func()) {
	if *me.err != nil {
		return
	}
	step()
	if *me.err != nil {
		if !strings.Contains(format, "%w") {
			format += ": %w"
		}
		*me.err = fmt.Errorf(format, *me.err)
	}
}

// Step checks internal error before calling the next step.
func (me *Stepper) Step(step func()) {
	if *me.err != nil {
		return
	}
	step()
}
