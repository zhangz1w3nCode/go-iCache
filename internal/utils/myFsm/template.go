package myFsm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

// StateEventRelationship 定义了状态和事件的关系
type StateEventRelationship struct {
	States   []string            `json:"states"`
	Events   []string            `json:"events"`
	Relation map[string][]string `json:"relation"`
}

//type Node struct {
//	id   string
//	text Text
//}
//type Text struct {
//	value string
//	x     int
//	y     int
//}
//type Edge struct {
//	sourceNodeId string
//	targetNodeId string
//}

// GraphConfigData 定义图配置数据
type GraphConfigData struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

// Node 定义节点的结构
type Node struct {
	ID   *string `json:"id,omitempty"`
	Text struct {
		Value string `json:"value"`
	} `json:"text"`
}

// Edge 定义边的结构
type Edge struct {
	ID           string `json:"id"`
	SourceNodeId string `json:"sourceNodeId"`
	TargetNodeId string `json:"targetNodeId"`
	Text         struct {
		Value string `json:"value"`
	} `json:"text"`
}

// FsmGoFileTemplate 是生成的 Go 文件的模板
const FsmGoFileTemplate = `package {{.Package}}

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
    "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//状态枚举
type State string
type Event string

// 定义状态和事件
const(
{{range.States}}
{{.}} State = "{{.}}"
{{end}}

{{range.Events}}
{{.}} Event = "{{.}}"
{{end}}
)

// 定义事件和对应的状态转换
var events = fsm.Events{
	{{- range $eventName, $pairs := .Relation}}
	{{- range $pair := $pairs}}
	{Name: string({{$eventName}}), Src: []string{string({{$pair.SourceState}})}, Dst: string({{$pair.TargetState}})},
	{{- end}}
	{{- end}}
}



// 定义状态机回调...
var callbacks = fsm.Callbacks{
	{{range .States}}
	"enter_" + string({{.}}): func(_ context.Context, e *fsm.Event) {
		fmt.Printf("状态改变为: %s\n", e.FSM.Current())
		return
	},
	{{end}}
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
	err := c.myFsm.Event(context.Background(),event)
	if err != nil {
		return "", err
	}
	return c.myFsm.Current(), nil
}
`

// 定义 HTTP 处理函数
func GenerateFsmHandler(w http.ResponseWriter, r *http.Request) {
	// 设置 CORS 响应头
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// 处理 OPTIONS 预检请求
	if r.Method == "OPTIONS" {
		return
	}

	// 读取请求体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// 解析请求参数
	var ser StateEventRelationship
	if err := json.Unmarshal(body, &ser); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// 解析请求参数
	var req GraphConfigData
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	status := ExtractStates(req.Nodes)
	events := ExtractEvents(req.Edges)
	extractedRelationships := ExtractRelationships(req.Nodes, req.Edges)

	fmt.Printf("状态集合: %v\n", status)
	fmt.Printf("事件集合: %v\n", events)
	fmt.Printf("状态和事件的关系: %v\n", extractedRelationships)

	CreateTemplate(status, events, extractedRelationships)
}

// ExtractStates 提取状态集合
func ExtractStates(nodes []Node) []string {
	states := make(map[string]bool)
	for _, node := range nodes {
		if node.Text.Value != "" {
			states[node.Text.Value] = true
		}
	}
	var stateList []string
	for state := range states {
		stateList = append(stateList, state)
	}
	return stateList
}

// ExtractEvents 提取事件集合
func ExtractEvents(edges []Edge) []string {
	events := make(map[string]bool)
	for _, edge := range edges {
		if edge.Text.Value != "" {
			events[edge.Text.Value] = true
		}
	}
	var eventList []string
	for event := range events {
		eventList = append(eventList, event)
	}
	return eventList
}

type StatusPair struct {
	SourceState string
	TargetState string
}

// ExtractRelationships 提取状态和事件的关系
func ExtractRelationships(nodes []Node, edges []Edge) map[string][]*StatusPair {
	relationships := make(map[string][]*StatusPair)
	for _, edge := range edges {
		sourceState := ""
		targetState := ""
		for _, node := range nodes {
			if *node.ID == edge.SourceNodeId {
				sourceState = node.Text.Value
			}
			if *node.ID == edge.TargetNodeId {
				targetState = node.Text.Value
			}
		}
		if _, exists := relationships[edge.Text.Value]; !exists {
			relationships[edge.Text.Value] = []*StatusPair{{sourceState, targetState}}
		} else {
			relationships[edge.Text.Value] = append(relationships[edge.Text.Value], &StatusPair{sourceState, targetState})
		}
	}
	return relationships
}

func CreateTemplate(states []string, events []string, relation map[string][]*StatusPair) error {
	data := struct {
		Package  string
		States   []string
		Events   []string
		Relation map[string][]*StatusPair
	}{
		Package:  "fsm",
		States:   states,
		Events:   events,
		Relation: relation,
	}
	// 渲染模板
	t, err := template.New("myFsm").Parse(FsmGoFileTemplate)
	if err != nil {
		return err
	}
	// 创建 Go 文件
	file, err := os.Create("internal/utils/fsm/myFsm.go")
	if err != nil {
		return err
	}
	defer file.Close()
	if err := t.Execute(file, data); err != nil {
		return err
	}
	// 返回成功响应
	fmt.Printf("Go file 'myFsm.go' generated successfully")
	return nil
}

//func main() {
//	// 设置 HTTP 路由
//	http.HandleFunc("/generate_fsm", generateFsmHandler)
//
//	// 启动 HTTP 服务器
//	fmt.Println("Server is running on port 8080...")
//	http.ListenAndServe(":8080", nil)
//}
