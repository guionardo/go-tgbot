package helpers

import "strings"

type (
	BotMenuOption struct {
		Command     string
		Caption     string
		Value       string
		IsLineBreak bool
	}
)

func (b BotMenuOption) String() string {
	if b.IsLineBreak {
		return "-"
	}
	if len(b.Value) == 0 {
		b.Value = b.Caption
	}
	if len(b.Command) > 0 {
		return b.Caption + ":" + b.Command + "|" + b.Value
	}
	return b.Caption + ":" + b.Value
}

func (b BotMenuOption) MessageValue() string {
	if len(b.Command) > 0 {
		return b.Command + "|" + b.Value
	} else {
		return b.Value
	}
}

func ParseBotMenuOption(option string) BotMenuOption {
	if option == "-" {
		return BotMenuOption{
			IsLineBreak: true,
		}
	}

	caption, value, found := strings.Cut(option, ":")
	if !found {
		caption = option
		value = option
	}
	command, newValue, found := strings.Cut(value, "|")
	if !found {
		command = ""
	} else {
		value = newValue
	}

	return BotMenuOption{
		Caption: caption,
		Command: command,
		Value:   value,
	}
}
