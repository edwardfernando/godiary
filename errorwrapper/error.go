package errorwrapper

// ErrWrapper is a wrapper for error with err type custom field
type ErrWrapper struct {
	ErrType string
	Err     error
}
