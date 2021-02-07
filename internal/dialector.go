package gondm

// Dialector GORDM database dialector
type Dialector interface {
	Name() string
	Init(*DB) error
}
