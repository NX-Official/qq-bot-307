package svc

import (
	"github.com/sashabaranov/go-openai"
	"qq-bot-307/service/http/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	GPTClient *openai.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		GPTClient: openai.NewClient(c.ChatGPTApiKey),
	}
}
