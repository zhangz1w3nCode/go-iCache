package fsm

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

StateStart State = "StateStart"

StateHandle State = "StateHandle"

StateFinish State = "StateFinish"



DoEvent Event = "DoEvent"

FinishEvent Event = "FinishEvent"

)

var events = fsm.Events{
	{Name: string(DoEvent), Src: []string{string(StateStart)}, Dst: string(StateHandle)},
	{Name: string(FinishEvent), Src: []string{string(StateHandle)}, Dst: string(StateFinish)},
}

var callbacks = fsm.Callbacks{
	
	"enter_" + string(StateStart): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	
	"enter_" + string(StateHandle): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	
	"enter_" + string(StateFinish): func(_ context.Context, e *fsm.Event) {
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
