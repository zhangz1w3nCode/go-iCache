package flow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"visual-state-machine/internal/entity"
	flowTemp "visual-state-machine/internal/utils/template"
)

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
	var req entity.GraphConfigData
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	status := ExtractStates(req.Nodes)
	events := ExtractEvents(req.Edges)
	extractedRelationships := ExtractRelationships(req.Nodes, req.Edges)

	CreateVisualFlow(status, events, extractedRelationships)
}

// ExtractStates 提取状态集合
func ExtractStates(nodes []entity.Node) []string {
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
func ExtractEvents(edges []entity.Edge) []string {
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

// ExtractRelationships 提取状态和事件的关系
func ExtractRelationships(nodes []entity.Node, edges []entity.Edge) map[string][]*entity.StatusPair {
	relationships := make(map[string][]*entity.StatusPair)
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
			relationships[edge.Text.Value] = []*entity.StatusPair{{sourceState, targetState}}
		} else {
			relationships[edge.Text.Value] = append(relationships[edge.Text.Value], &entity.StatusPair{sourceState, targetState})
		}
	}
	return relationships
}

func CreateVisualFlow(states []string, events []string, relation map[string][]*entity.StatusPair) error {
	data := &entity.TemplateParam{
		Package:  "flow_temp",
		States:   states,
		Events:   events,
		Relation: relation,
	}
	// 渲染模板
	t, err := template.New("flow_temp").Parse(flowTemp.FsmGoFileTemplate)
	if err != nil {
		return err
	}
	// 创建 Go 文件
	file, err := os.Create("internal/generate/flow/flow_temp.go")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	if err := t.Execute(file, data); err != nil {
		return err
	}
	// 返回成功响应
	fmt.Printf("Go file 'flow_temp.go' generated successfully")
	return nil
}
