package template

// FsmGoFileTemplate 是生成的 Go 文件的模板
const FsmGoFileTemplate = `package {{.Package}}

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
{{range.States}}
{{.}} State = "{{.}}"
{{end}}

{{range.Events}}
{{.}} Event = "{{.}}"
{{end}}
)

var events = fsm.Events{
	{{- range $eventName, $pairs := .Relation}}
	{{- range $pair := $pairs}}
	{Name: string({{$eventName}}), Src: []string{string({{$pair.SourceState}})}, Dst: string({{$pair.TargetState}})},
	{{- end}}
	{{- end}}
}

var callbacks = fsm.Callbacks{
	{{range .States}}
	"enter_" + string({{.}}): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	{{end}}
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
`
