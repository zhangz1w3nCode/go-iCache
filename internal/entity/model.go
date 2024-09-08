package entity

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

type StatusPair struct {
	SourceState string
	TargetState string
}

type TemplateParam struct {
	Package  string
	States   []string
	Events   []string
	Relation map[string][]*StatusPair
}
