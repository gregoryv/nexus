package nexus

import "fmt"

func NewStepper(err *error) *Stepper {
	return &Stepper{err}
}

type Stepper struct {
	err *error
}

// Stepf checks internal error before calling the next step. If error is
// set after step it's wrapped using the given format string, which
// must contain only one %w.
func (me *Stepper) Stepf(format string, step func()) {
	if *me.err != nil {
		return
	}
	step()
	if *me.err != nil {
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
