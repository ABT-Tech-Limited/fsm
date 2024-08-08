package fsm

import (
	"errors"
	"fmt"
)

type DefaultStateMachine[S, E comparable, C any] struct {
	machineId string
	ready     bool
	states    map[S]State[S, E, C]
}

func NewDefaultStateMachine[S, E comparable, C any](machineId string) *DefaultStateMachine[S, E, C] {
	return &DefaultStateMachine[S, E, C]{
		machineId: machineId,
		states:    make(map[S]State[S, E, C]),
		ready:     false,
	}
}

func (dsm *DefaultStateMachine[S, E, C]) AddExternalTransition(fromStateId, toStateId S, event E, when Condition[C], trigger Action[S, E, C]) {
	fromState := dsm.getState(fromStateId)
	toState := dsm.getState(toStateId)

	transition := fromState.AddTransition(event, toState, TransitionTypes.External)
	if when != nil {
		transition.SetCondition(when)
	}
	transition.SetAction(trigger)
}

func (dsm *DefaultStateMachine[S, E, C]) AddInternalTransition(stateId S, event E, when Condition[C], trigger Action[S, E, C]) {
	state := dsm.getState(stateId)

	transition := state.AddTransition(event, state, TransitionTypes.Internal)
	if when != nil {
		transition.SetCondition(when)
	}
	transition.SetAction(trigger)
}

func (dsm *DefaultStateMachine[S, E, C]) Ready() {
	dsm.ready = true
}

func (dsm *DefaultStateMachine[S, E, C]) GetMachineId() string {
	return dsm.machineId
}

func (dsm *DefaultStateMachine[S, E, C]) FireEvent(sourceStateId S, event E, ctx C) (S, error) {
	if err := dsm.ensureReady(); err != nil {
		return sourceStateId, err
	}

	transition := dsm.stateTransitionRoute(sourceStateId, event, ctx)
	if transition == nil {
		return sourceStateId, errors.New(fmt.Sprintf("there is no transition for this state: %s and event: %s or not satisfied", sourceStateId, event))
	}

	state, err := transition.Transit(ctx, true)
	if err != nil {
		return sourceStateId, err
	}

	return state.GetStateId(), err
}

func (dsm *DefaultStateMachine[S, E, C]) ensureReady() error {
	if !dsm.ready {
		return errors.New("state machine not ready, can't work, please set ready first")
	}
	return nil
}

func (dsm *DefaultStateMachine[S, E, C]) getState(stateId S) State[S, E, C] {
	state, ok := dsm.states[stateId]
	if !ok {
		state = NewDefaultState[S, E, C](stateId)
		dsm.states[stateId] = state
	}
	return state
}

func (dsm *DefaultStateMachine[S, E, C]) stateTransitionRoute(sourceStateId S, event E, ctx C) Transition[S, E, C] {
	state := dsm.getState(sourceStateId)
	transitions := state.GetEventTransitions(event)
	if len(transitions) == 0 {
		return nil
	}

	// Find the target transition according to conditions
	var transition Transition[S, E, C]
	for _, v := range transitions {
		if v.GetCondition() == nil {
			transition = v
		} else if v.GetCondition().Satisfied(ctx) {
			transition = v
			break
		}
	}
	return transition
}
