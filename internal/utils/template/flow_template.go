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
	{Name: {{$eventName}}.GetString(), Src: []string{string({{$pair.SourceState}})}, Dst: string({{$pair.TargetState}})},
	{{- end}}
	{{- end}}
}

var callbacks = fsm.Callbacks{
	{{range .States}}
	"enter_" + {{.}}.GetString(): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	{{end}}
}

type StatusContext struct {
	myFsm *fsm.FSM
}

func GenerateFSM(status string) *StatusContext {
	fsmCtx := &StatusContext{}
	fsmCtx.myFsm = fsm.NewFSM(status, events, callbacks)
	return fsmCtx
}

func (s State) GetString() string {
	return string(s)
}

func (e Event) GetString() string {
	return string(e)
}

func (c *StatusContext) Trans(ctx context.Context, event Event) (string, error) {
	if !c.myFsm.Can(event.GetString()) {
		return "", status.Errorf(codes.InvalidArgument,
			"Tans status invalid! current status:%s  "+
				"current event: %s", c.myFsm.Current(), event)
	}
	err := c.myFsm.Event(ctx, event.GetString())
	if err != nil {
		return "", err
	}
	return c.myFsm.Current(), nil
}
`
