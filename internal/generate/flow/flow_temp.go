package flow_temp

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
    "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type State string
type Event string

const(

Init State = "Init"

Queue State = "Queue"

Ring State = "Ring"

Answering State = "Answering"

Answered State = "Answered"

Unanswered State = "Unanswered"



CALLING_FINISH Event = "CALLING_FINISH"

CALLING_DEQUEUE Event = "CALLING_DEQUEUE"

CALLING_ACCEPT Event = "CALLING_ACCEPT"

CALLING_QUEUE_HANGUP Event = "CALLING_QUEUE_HANGUP"

CALLING_HANGUP Event = "CALLING_HANGUP"

CALLING_ENQUEUE Event = "CALLING_ENQUEUE"

)

var events = fsm.Events{
	{Name: string(CALLING_ACCEPT), Src: []string{string(Ring)}, Dst: string(Answering)},
	{Name: string(CALLING_DEQUEUE), Src: []string{string(Queue)}, Dst: string(Ring)},
	{Name: string(CALLING_ENQUEUE), Src: []string{string(Init)}, Dst: string(Queue)},
	{Name: string(CALLING_FINISH), Src: []string{string(Answering)}, Dst: string(Answered)},
	{Name: string(CALLING_HANGUP), Src: []string{string(Ring)}, Dst: string(Unanswered)},
	{Name: string(CALLING_QUEUE_HANGUP), Src: []string{string(Queue)}, Dst: string(Unanswered)},
}

var callbacks = fsm.Callbacks{
	
	"enter_" + string(Init): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	
	"enter_" + string(Queue): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	
	"enter_" + string(Ring): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	
	"enter_" + string(Answering): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	
	"enter_" + string(Answered): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	
	"enter_" + string(Unanswered): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	
}

type FsmContext struct {
	myFsm *fsm.FSM
}

func GenerateFSM(status string) *FsmContext {
	fsmCtx := &FsmContext{}
	fsmCtx.myFsm = fsm.NewFSM(status, events, callbacks)
	return fsmCtx
}

func (c *FsmContext) Trans(event string) (string, error) {
	if !c.myFsm.Can(event) {
		return "", status.Errorf(codes.InvalidArgument, "invalid event: %s", event)
	}
	err := c.myFsm.Event(context.Background(),event)
	if err != nil {
		return "", err
	}
	return c.myFsm.Current(), nil
}
