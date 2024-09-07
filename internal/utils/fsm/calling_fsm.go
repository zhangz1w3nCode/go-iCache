package fsm

//
//import (
//	"context"
//	"fmt"
//
//	"github.com/looplab/fsm"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//
//	"github.com/MoeGolibrary/moego-svc-engagement/internal/entity"
//	"github.com/MoeGolibrary/moego-svc-engagement/internal/logic/calling"
//)
//
//// CallState 定义状态和事件枚举
//type CallState string
//
//type CallEvent string
//
//const (
//	InitState       CallState = "Init"
//	RingState       CallState = "Ring"
//	AnsweringState  CallState = "Answering"
//	AnsweredState   CallState = "Answered"
//	QueueState      CallState = "Queue"
//	UnansweredState CallState = "Unanswered"
//
//	CallingEnqueueEvent     CallEvent = "CALLING_ENQUEUE"
//	CallingDequeueEvent     CallEvent = "CALLING_DEQUEUE"
//	CallingQueueHangupEvent CallEvent = "CALLING_QUEUE_HANGUP"
//	CallingAcceptEvent      CallEvent = "CALLING_ACCEPT"
//	CallingHangupEvent      CallEvent = "CALLING_HANGUP"
//	CallingFinishEvent      CallEvent = "CALLING_FINISH"
//)
//
//// 定义事件和对应的状态转换
//var events = fsm.Events{
//	{Name: string(CallingEnqueueEvent), Src: []string{string(InitState)}, Dst: string(QueueState)},
//	{Name: string(CallingDequeueEvent), Src: []string{string(QueueState)}, Dst: string(RingState)},
//	{Name: string(CallingQueueHangupEvent), Src: []string{string(QueueState)}, Dst: string(UnansweredState)},
//	{Name: string(CallingAcceptEvent), Src: []string{string(RingState)}, Dst: string(AnsweringState)},
//	{Name: string(CallingHangupEvent), Src: []string{string(RingState)}, Dst: string(UnansweredState)},
//	{Name: string(CallingFinishEvent), Src: []string{string(AnsweringState)}, Dst: string(AnsweredState)},
//}
//
//type CallContext struct {
//	callFsm *fsm.FSM
//	calling *entity.CallingLogStatus
//	//additional fields
//	log *calling.Log
//}
//
//func GenerateCallingFSM(log *entity.CallingLogStatus) *CallContext {
//	callContext := &CallContext{}
//	if log != nil {
//		callContext.calling = log
//	} else {
//		callContext.calling = &entity.CallingLogStatus{
//			Status: string(InitState),
//		}
//	}
//
//	//定义状态机回调
//	var callbacks = fsm.Callbacks{
//		//进入QueueState触发的回调方法
//		"enter_" + string(QueueState): func(_ context.Context, e *fsm.Event) {
//			fmt.Printf("状态改变为: %s\n", e.FSM.Current())
//			callContext.calling.Status = e.FSM.Current()
//			err := callContext.log.UpdateCallingLogStatus(callContext.ctx, callContext.calling)
//			if err != nil {
//				return
//			}
//			return
//		},
//		"enter_" + string(RingState): func(_ context.Context, e *fsm.Event) {
//			fmt.Printf("状态改变为: %s\n", e.FSM.Current())
//			callContext.calling.Status = e.FSM.Current()
//			err := callContext.log.UpdateCallingLogStatus(callContext.ctx, callContext.calling)
//			if err != nil {
//				return
//			}
//			return
//		},
//		"enter_" + string(AnsweringState): func(_ context.Context, e *fsm.Event) {
//			fmt.Printf("状态改变为: %s\n", e.FSM.Current())
//			callContext.calling.Status = e.FSM.Current()
//			err := callContext.log.UpdateCallingLogStatus(callContext.ctx, callContext.calling)
//			if err != nil {
//				return
//			}
//			return
//		},
//		"enter_" + string(AnsweredState): func(_ context.Context, e *fsm.Event) {
//			fmt.Printf("状态改变为: %s\n", e.FSM.Current())
//			callContext.calling.Status = e.FSM.Current()
//			err := callContext.log.UpdateCallingLogStatus(callContext.ctx, callContext.calling)
//			if err != nil {
//				return
//			}
//			return
//		},
//
//		"enter_" + string(UnansweredState): func(_ context.Context, e *fsm.Event) {
//			fmt.Printf("状态改变为: %s\n", e.FSM.Current())
//			callContext.calling.Status = e.FSM.Current()
//			err := callContext.log.UpdateCallingLogStatus(callContext.ctx, callContext.calling)
//			if err != nil {
//				return
//			}
//			return
//		},
//
//		"enter_" + string(CallingFinishEvent): func(_ context.Context, e *fsm.Event) {
//			fmt.Printf("状态改变为: %s\n", e.FSM.Current())
//			callContext.calling.Status = e.FSM.Current()
//			err := callContext.log.UpdateCallingLogStatus(callContext.ctx, callContext.calling)
//			if err != nil {
//				return
//			}
//			return
//		},
//	}
//	callContext.callFsm = fsm.NewFSM(callContext.calling.Status, events, callbacks)
//	return callContext
//}
//
//func (c *CallContext) Trans(event CallEvent) error {
//	if !c.callFsm.Can(string(event)) {
//		return status.Errorf(codes.InvalidArgument, "invalid event: %s", event)
//	}
//	err := c.callFsm.Event(c.ctx, string(event))
//	if err != nil {
//		return err
//	}
//	return nil
//}
