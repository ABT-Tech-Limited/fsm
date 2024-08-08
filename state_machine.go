package fsm

type StateMachine[S, E comparable, C any] interface {
	GetMachineId() string
	FireEvent(sourceId S, event E, ctx C) (S, error) // Send an event to the state machine, driving state flow.
}
