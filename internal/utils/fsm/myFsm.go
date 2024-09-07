package fsm

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 状态枚举
type State string
type Event string

// 定义状态和事件
const (
	PayNeed State = "PayNeed"

	PayPaying State = "PayPaying"

	PaySuccess State = "PaySuccess"

	PayFail State = "PayFail"

	PayClose State = "PayClose"

	PayStart State = "PayStart"

	PaySuccessEvent Event = "PaySuccessEvent"

	PayFailEvent Event = "PayFailEvent"

	PayCloseEvent Event = "PayCloseEvent"

	PayInitEvent Event = "PayInitEvent"

	PayInvokeEvent Event = "PayInvokeEvent"
)

// 定义事件和对应的状态转换
var events = fsm.Events{
	{Name: string(PayCloseEvent), Src: []string{string(PaySuccess)}, Dst: string(PayClose)},
	{Name: string(PayCloseEvent), Src: []string{string(PayFail)}, Dst: string(PayClose)},
	{Name: string(PayFailEvent), Src: []string{string(PayPaying)}, Dst: string(PayFail)},
	{Name: string(PayInitEvent), Src: []string{string(PayStart)}, Dst: string(PayNeed)},
	{Name: string(PayInvokeEvent), Src: []string{string(PayNeed)}, Dst: string(PayPaying)},
	{Name: string(PaySuccessEvent), Src: []string{string(PayPaying)}, Dst: string(PaySuccess)},
}

// 定义状态机回调...
var callbacks = fsm.Callbacks{

	"enter_" + string(PayNeed): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},

	"enter_" + string(PayPaying): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},

	"enter_" + string(PaySuccess): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},

	"enter_" + string(PayFail): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},

	"enter_" + string(PayClose): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},

	"enter_" + string(PayStart): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
}

// FsmContext 结构体...
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
	err := c.myFsm.Event(context.Background(), event)
	if err != nil {
		return "", err
	}
	return c.myFsm.Current(), nil
}
