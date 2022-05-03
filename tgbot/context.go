package tgbot

import (
	"context"
)

const ctxKey = "svc"

func CreateBotContext(svc *GoTGBotSevice) (ctx context.Context, cancel context.CancelFunc) {
	cancelCtx, cancel := context.WithCancel(context.WithValue(context.Background(), ctxKey, svc))

	return cancelCtx, cancel
}

func GetBotService(ctx context.Context) *GoTGBotSevice {
	return ctx.Value(ctxKey).(*GoTGBotSevice)
}
