package fsm

import (
	"errors"
	"fmt"
)

type DefaultTransition[S, E comparable, C any] struct {
	transitionType TransitionType
	source         State[S, E, C]
	target         State[S, E, C]
	event          E
	condition      Condition[C]
	action         Action[S, E, C]
}

func NewDefaultTransition[S, E comparable, C any]() *DefaultTransition[S, E, C] {
	return &DefaultTransition[S, E, C]{}
}

func (d *DefaultTransition[S, E, C]) GetEvent() E {
	return d.event
}

func (d *DefaultTransition[S, E, C]) SetEvent(event E) {
	d.event = event
}

func (d *DefaultTransition[S, E, C]) SetType(typ TransitionType) {
	d.transitionType = typ
}

func (d *DefaultTransition[S, E, C]) GetSource() State[S, E, C] {
	return d.source
}

func (d *DefaultTransition[S, E, C]) SetSource(source State[S, E, C]) {
	d.source = source
}

func (d *DefaultTransition[S, E, C]) GetTarget() State[S, E, C] {
	return d.target
}

func (d *DefaultTransition[S, E, C]) SetTarget(target State[S, E, C]) {
	d.target = target
}

func (d *DefaultTransition[S, E, C]) GetCondition() Condition[C] {
	return d.condition
}

func (d *DefaultTransition[S, E, C]) SetCondition(condition Condition[C]) {
	d.condition = condition
}

func (d *DefaultTransition[S, E, C]) GetAction() Action[S, E, C] {
	return d.action
}

func (d *DefaultTransition[S, E, C]) SetAction(action Action[S, E, C]) {
	d.action = action
}

func (d *DefaultTransition[S, E, C]) Transit(ctx C, checkCondition bool) (State[S, E, C], error) {
	if err := d.Verify(); err != nil {
		return d.source, err
	}

	if checkCondition && d.condition != nil && !d.condition.Satisfied(ctx) {
		return d.source, errors.New(fmt.Sprintf("condition is not satisfied"))
	}

	if d.action != nil {
		err := d.action.Execute(d.source.GetStateId(), d.target.GetStateId(), d.event, ctx)
		if err != nil {
			return d.source, err
		}
	}
	return d.target, nil
}

func (d *DefaultTransition[S, E, C]) Verify() error {
	if d.transitionType == TransitionTypes.Internal && d.source != d.target {
		return errors.New("internal transition source must equal to target")
	}
	return nil
}
