package helpers

import (
	"reflect"
	"testing"
)

func TestParseBotMenuOption(t *testing.T) {
	tests := []struct {
		name string
		args string
		want BotMenuOption
	}{
		{
			name: "line break",
			args: "-",
			want: BotMenuOption{
				IsLineBreak: true,
			},
		},
		{
			name: "Simple option",
			args: "Opção 1",
			want: BotMenuOption{
				Caption: "Opção 1",
				Value:   "Opção 1",
			},
		},
		{
			name: "Value option",
			args: "Opção 1:opcao1",
			want: BotMenuOption{
				Caption: "Opção 1",
				Value:   "opcao1",
			},
		},
		{
			name: "Command Value option",
			args: "Opção 1:MENU|opcao1",
			want: BotMenuOption{
				Caption: "Opção 1",
				Value:   "opcao1",
				Command: "MENU",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseBotMenuOption(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseBotMenuOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBotMenuOption_String(t *testing.T) {
	type fields struct {
		Command     string
		Caption     string
		Value       string
		IsLineBreak bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "line break",
			fields: fields{
				IsLineBreak: true,
			},
			want: "-",
		},
		{
			name: "Simple option",
			fields: fields{
				Caption: "Opção 1",
			},
			want: "Opção 1:Opção 1",
		},
		{
			name: "Value option",
			fields: fields{
				Caption: "Opção 1",
				Value:   "opcao1",
			},
			want: "Opção 1:opcao1",
		},
		{
			name: "Command Value option",
			fields: fields{
				Caption: "Opção 1",
				Value:   "opcao1",
				Command: "MENU",
			},
			want: "Opção 1:MENU|opcao1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BotMenuOption{
				Command:     tt.fields.Command,
				Caption:     tt.fields.Caption,
				Value:       tt.fields.Value,
				IsLineBreak: tt.fields.IsLineBreak,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("BotMenuOption.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
