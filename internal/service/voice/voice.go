package voice

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"visual-state-machine/internal/utils/fsm"

	errorHandler "github.com/MoeGolibrary/go-lib/error-handler"
	"github.com/MoeGolibrary/go-lib/zlog"
	engagementpb "github.com/MoeGolibrary/moego-api-definitions/out/go/moego/models/engagement/v1"
	engagementsvcpb "github.com/MoeGolibrary/moego-api-definitions/out/go/moego/service/engagement/v1"
	"go.uber.org/zap"
)

type HookService struct {
	*engagementsvcpb.UnimplementedTwilioHookServiceServer
}

func NewHookService() *HookService {
	return &HookService{}
}

func (s *HookService) Status(ctx context.Context, req *engagementsvcpb.TwilioHookRequest) (
	*engagementsvcpb.TwilioHookResponse, error) {

	params, err := url.ParseQuery(req.GetContent())
	if err != nil {
		zlog.Error(ctx, "parse query failed", zap.Error(err))
		return nil, errorHandler.WrapError(err)
	}
	status := params.Get("Status")
	event := params.Get("Event")

	currentStatus, err := fsm.GenerateFSM(status).Trans(event)
	if err != nil {
		return nil, err
	}

	fmt.Println(currentStatus)

	return &engagementsvcpb.TwilioHookResponse{Content: currentStatus}, nil
}

func (s *HookService) direction(params url.Values) engagementpb.CallingDirection {
	// From 字段中包含 "client:" 视为呼出
	from := params.Get("From")
	if strings.Contains(from, "client:") {
		return engagementpb.CallingDirection_OUTGOING
	}
	// staff 使用 B-APP 联系用户视为呼出，这种场景下会带有 Digits 字段
	if params.Has("Digits") {
		return engagementpb.CallingDirection_OUTGOING
	}

	return engagementpb.CallingDirection_INCOMING
}
