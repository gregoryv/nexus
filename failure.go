package nexus

type Failure struct {
	err error
}

// Fail sets the given error unless already failed. An error can only be set once.
func (me *Failure) Fail(err error) error {
	if me.err != nil {
		return me.err
	}
	me.err = err
	return err
}

// Ok return true if no error is set.
func (me *Failure) Ok() bool { return me.err == nil }

// Error returns the failure error or nil if not set
func (me *Failure) Error() error { return me.err }
