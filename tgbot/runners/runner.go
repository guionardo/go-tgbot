package runners

import (
	"context"
	"time"

	"github.com/guionardo/go-tgbot/tgbot/infra"
)



func (runner *Runner) Start(ctx context.Context, doneChannel chan string) {
	if runner.customLoop {
		runner.startCustomLoop(ctx, doneChannel)
	} else {
		runner.startCommonLoop(ctx, doneChannel)
	}
	runner.Logger.Infof("removing runner %s", runner.Name)
	doneChannel <- runner.Name
}

func (runner *Runner) startCustomLoop(ctx context.Context, doneChannel chan string) {
	if err := runner.runnerLoop(ctx, runner); err != nil {
		runner.Logger.Error(err)
	}
}
func (runner *Runner) startCommonLoop(ctx context.Context, doneChannel chan string) {
	canRun := true
	for canRun {
		select {
		case <-ctx.Done():
			runner.Logger.Info("context done")
			canRun = false
		case <-runner.StopChannel:
			runner.Logger.Info("stopped by channel")
			canRun = false
		default:
			if err := runner.runnerLoop(ctx, runner); err != nil {
				runner.Logger.Error(err)
				canRun = false
			}
		}
	}
}

func (runner *Runner) Stop() {
	runner.Logger.Info("stopping...")
	runner.StopChannel <- true
}

func NewRunnerCollection() *Runners {
	return &Runners{
		runners: make(map[string]*Runner),
		logger:  infra.GetLogger("Runners"),
	}
}

func (runners *Runners) CreateRunner(name string, actionLoop RunnerLoop, data interface{}) *Runner {
	runner := &Runner{
		Name:        name,
		Logger:      infra.GetLogger(name),
		StopChannel: make(chan bool, 1),
		runnerLoop:  actionLoop,
		Data:        data,
	}
	runners.runners[runner.Name] = runner
	return runner
}

func (runners *Runners) CreateRunnerCustomLoop(name string, actionLoop RunnerLoop, data interface{}) *Runner {
	runner := &Runner{
		Name:        name,
		Logger:      infra.GetLogger(name),
		StopChannel: make(chan bool, 1),
		runnerLoop:  actionLoop,
		Data:        data,
		customLoop:  true,
	}
	runners.runners[runner.Name] = runner
	return runner
}

func (runners *Runners) RunAll(ctx context.Context) {
	lastRunnersCount := len(runners.runners)
	stoppedChannel := make(chan string, lastRunnersCount)
	for _, runner := range runners.runners {
		go runner.Start(ctx, stoppedChannel)
	}
	for {
		select {
		case runnerName := <-stoppedChannel:
			delete(runners.runners, runnerName)
			runners.logger.Infof("Removed runner %s", runnerName)
		case <-ctx.Done():
			runners.logger.Info("context done")
			return
		default:
			if len(runners.runners) == 0 {
				runners.logger.Info("all runners stopped")
				return
			}

			if len(runners.runners) != lastRunnersCount {
				runnersNames := make([]string, len(runners.runners))
				index := 0
				for k := range runners.runners {
					runnersNames[index] = k
					index++
				}
				runners.logger.Info("waiting for runners to stop", runnersNames)
				lastRunnersCount = len(runners.runners)
			}
			time.Sleep(time.Millisecond * 1000)

		}
	}
}
