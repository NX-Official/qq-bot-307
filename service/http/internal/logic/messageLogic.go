package logic

import (
	"context"
	"qq-bot-307/common/gpt"
	"qq-bot-307/service/http/internal/svc"
	"qq-bot-307/service/http/internal/types"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageLogic {
	return &MessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageLogic) Message(req *types.MessageRequest) (resp *types.MessageReply, err error) {

	trigger := "[CQ:at,qq=2165526145] "

	if req.PostType == "message" && req.MessageType == "group" && strings.HasPrefix(req.Message, trigger) {
		l.Logger.Info(req)
		gptReply := gpt.Chat(strings.TrimPrefix(req.Message, trigger), l.svcCtx.GPTClient)
		l.Logger.Info(gptReply)
		return &types.MessageReply{
			Reply: gptReply,
		}, nil
	}

	return nil, nil
}
