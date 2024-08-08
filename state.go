package fsm

// State definition
type State[S, E comparable, C any] interface {
	GetStateId() S                                                                                   // get current state
	AddTransition(event E, target State[S, E, C], transitionType TransitionType) Transition[S, E, C] // add transition(from target to this state)
	GetEventTransitions(event E) []Transition[S, E, C]                                               // get transitions by event
	GetAllTransitions() []Transition[S, E, C]                                                        // get all transitions
}
