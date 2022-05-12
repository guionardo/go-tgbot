package tgbot

func (svc *GoTGBotService) AddHandlers(handlers ...*ListenerFilteredHandler) *GoTGBotService {
	for _, handler := range handlers {
		svc.listener.AddHandler(handler.Title, handler.Filter, handler.Func)
	}
	svc.setupLevel = Set(svc.setupLevel, Handlers)
	return svc
}

func (svc *GoTGBotService) AddCommandHandlers(handlers ...*ListenerCommandHandler) *GoTGBotService {
	for _, handler := range handlers {
		svc.listener.AddCommandHandler(handler.Title, handler.Command, handler.Func)
	}
	svc.setupLevel = Set(svc.setupLevel, Handlers)
	return svc
}
