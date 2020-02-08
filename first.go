package nexus

// First returns the first error of the given ones or nil
func First(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}
