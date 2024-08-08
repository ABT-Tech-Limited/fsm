package fsm

func NewFsm[S, E comparable, C any](id string) *DefaultStateMachine[S, E, C] {
	return NewDefaultStateMachine[S, E, C](id)
}
