package fsm

type DefaultState[S, E comparable, C any] struct {
	stateId       S
	transitionMap map[E][]Transition[S, E, C]
}

func NewDefaultState[S, E comparable, C any](stateId S) *DefaultState[S, E, C] {
	return &DefaultState[S, E, C]{
		stateId:       stateId,
		transitionMap: make(map[E][]Transition[S, E, C]),
	}
}

func (d *DefaultState[S, E, C]) GetStateId() S {
	return d.stateId
}

func (d *DefaultState[S, E, C]) AddTransition(event E, target State[S, E, C], transitionType TransitionType) Transition[S, E, C] {
	transition := NewDefaultTransition[S, E, C]()
	transition.SetSource(d)
	transition.SetTarget(target)
	transition.SetEvent(event)
	transition.SetType(transitionType)
	d.transitionMap[event] = append(d.transitionMap[event], transition)
	return transition
}

func (d *DefaultState[S, E, C]) GetEventTransitions(event E) []Transition[S, E, C] {
	return d.transitionMap[event]
}

func (d *DefaultState[S, E, C]) GetAllTransitions() []Transition[S, E, C] {
	var result []Transition[S, E, C]
	for _, t := range d.transitionMap {
		result = append(result, t...)
	}
	return result
}
