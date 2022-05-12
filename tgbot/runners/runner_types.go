package runners

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type (
	Runner struct {
		Name        string
		Logger      *log.Entry
		StopChannel chan bool
		runnerLoop  RunnerLoop
		Data        interface{}
		customLoop  bool
	}

	// RunnerLoop is the function that will be called by the runner Start() method
	// The inner loop must check the context.Done() channel to stop the runner
	RunnerLoop func(ctx context.Context, runner *Runner) error
	
	Runners    struct {
		runners map[string]*Runner
		logger  *log.Entry
	}
)
