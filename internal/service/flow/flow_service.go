package flow

import (
	"context"
	"fmt"
	engagementsvcpb "github.com/MoeGolibrary/moego-api-definitions/out/go/moego/service/engagement/v1"
	flow_temp "visual-state-machine/internal/generate/flow"
)

type FlowService struct {
	*engagementsvcpb.UnimplementedFlowServiceServer
}

func NewFlowService() *FlowService {
	return &FlowService{}
}

func (s *FlowService) CreateFlow(ctx context.Context, req *engagementsvcpb.CreateFlowRequest) (
	*engagementsvcpb.CreateFlowResponse, error) {

	fmt.Println(req.GetStates())
	fmt.Println(req.GetEvents())
	latestStatus, err := flow_temp.GenerateFSM(req.GetStates()).Trans(req.GetEvents())
	if err != nil {
		return nil, err
	}
	return &engagementsvcpb.CreateFlowResponse{States: latestStatus}, nil
}
