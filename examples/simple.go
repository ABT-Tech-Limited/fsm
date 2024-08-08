package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/0xFredZhang/fsm"
)

type State string
type Event string
type Context struct {
	ctx string
}

const (
	StateInit       State = "Init"
	StatePending    State = "Pending"
	StateProcessing State = "Processing"
	StateSuccess    State = "Success"
	StateFailed     State = "Failed"

	EventPlaceOrder     Event = "PlaceOrder"
	EventDoDeposit      Event = "DoDeposit"
	EventDepositSuccess Event = "DepositSuccess"
	EventDepositTimeout Event = "DepositTimeout"
)

type orderHandler struct{}

func (p orderHandler) Execute(fromState, toState State, triggeredEvent Event, ctx *Context) error {
	log.Println(fmt.Sprintf("from %s to %s by %s", fromState, toState, triggeredEvent))
	return nil
}

func (p orderHandler) Name() string {
	return "placeOrderHandler"
}

func (p orderHandler) Satisfied(ctx *Context) bool {
	return false
}

func main() {
	sm := fsm.NewFsm[State, Event, *Context]("state-machine-id")
	sm.AddExternalTransition(StateInit, StatePending, EventPlaceOrder, orderHandler{}, orderHandler{})
	sm.AddExternalTransition(StatePending, StateProcessing, EventDoDeposit, orderHandler{}, orderHandler{})
	sm.AddExternalTransition(StateProcessing, StateSuccess, EventDepositSuccess, orderHandler{}, orderHandler{})
	sm.AddExternalTransition(StateProcessing, StateFailed, EventDepositTimeout, orderHandler{}, orderHandler{})
	sm.Ready()

	log.Println(sm.GetMachineId())

	ctx := &Context{ctx: "ctx"}

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		state, err := sm.FireEvent(StateInit, EventPlaceOrder, ctx)
		if err != nil {
			log.Panic(err)
		}
		log.Println(state)
		wg.Done()
	}()
	go func() {
		state, err := sm.FireEvent(StateInit, EventPlaceOrder, ctx)
		if err != nil {
			log.Panic(err)
		}
		log.Println(state)
		wg.Done()
	}()
	go func() {
		state, err := sm.FireEvent(StateInit, EventPlaceOrder, ctx)
		if err != nil {
			log.Panic(err)
		}
		log.Println(state)
		wg.Done()
	}()
	wg.Wait()
}
