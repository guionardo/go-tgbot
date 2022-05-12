package tgbot

import "testing"

func TestSetupLevel_Set(t *testing.T) {
	var sl SetupLevel

	type args struct {
		flag SetupLevel
	}
	tests := []struct {
		name    string
		cmd     func()
		test    func() bool
		failMsg string
	}{
		{
			name:    "Initial",
			cmd:     func() { sl = Set(sl, Instance) },
			test:    func() bool { return Has(sl, Instance) },
			failMsg: "Initial state should be 0",
		},
		{
			name:    "Set Configuration",
			cmd:     func() { sl = Set(sl, Configuration) },
			test:    func() bool { return Has(sl, Configuration) },
			failMsg: "Configuration should be set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cmd()
			if !tt.test() {
				t.Errorf("%s - %v", tt.failMsg, sl)
			}
		})
	}
}
