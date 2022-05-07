package tgbot

const (
	RunnerStop = "stop"
)

type (
	InternalChannel chan InternalMessage

	InternalMessage struct {
		source  string
		message string
	}
)

func NewInternalChannel() InternalChannel {
	return make(chan InternalMessage)
}

func (channel InternalChannel) Stop(source string) {	
	channel <- InternalMessage{source, RunnerStop}
}
