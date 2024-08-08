package fsm

// Condition
// pre-checking from source state to target state, can be null.
type Condition[C any] interface {
	Name() string         // Name of this condition
	Satisfied(ctx C) bool // Condition satisfied or not
}
