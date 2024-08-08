package fsm

// Transition Something that associates state changes.
type Transition[S, E comparable, C any] interface {
	GetEvent() E
	SetEvent(event E)

	SetType(typ TransitionType)

	GetSource() State[S, E, C]
	SetSource(source State[S, E, C])

	GetTarget() State[S, E, C]
	SetTarget(target State[S, E, C])

	GetCondition() Condition[C]
	SetCondition(condition Condition[C])

	GetAction() Action[S, E, C]
	SetAction(action Action[S, E, C])

	Transit(ctx C, checkCondition bool) (State[S, E, C], error)

	Verify() error
}
