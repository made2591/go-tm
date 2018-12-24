package transaction

type Transaction interface {
	Validate() bool
	Simulate() *State
	Execute() *State
}

type transaction struct{}
