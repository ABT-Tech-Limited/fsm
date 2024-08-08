package fsm

// Action
// The actual action that triggered by given event.
// And here is the extension point for implementing the business logic.
type Action[S, E comparable, C any] interface {
	Execute(fromState, toState S, triggeredEvent E, ctx C) error
}
