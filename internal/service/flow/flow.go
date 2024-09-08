package flow

import (
	"context"
	"fmt"
	engagementsvcpb "github.com/MoeGolibrary/moego-api-definitions/out/go/moego/service/engagement/v1"
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

	return &engagementsvcpb.CreateFlowResponse{}, nil
}
