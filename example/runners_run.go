package example

import (
	"context"
	"time"

	"github.com/guionardo/go-tgbot/tgbot/runners"
)

func ActionLoop1(ctx context.Context, runner *runners.Runner) error {
	runner.Logger.Info("action loop ", runner.Data)
	time.Sleep(time.Second * 2)
	counter := 0
	switch runner.Data.(type) {
	case int:
		counter = runner.Data.(int)
	}
	counter++
	if counter > 10 {
		runner.Stop()
	}
	runner.Data = counter
	return nil // fmt.Errorf("Exiting runner %s", runner.Name)
}

func ActionLoop2(ctx context.Context, runner *runners.Runner) error {
	runner.Logger.Info("action loop ", runner.Data)
	time.Sleep(time.Second * 1)
	counter := 0
	switch runner.Data.(type) {
	case int:
		counter = runner.Data.(int)
	}
	counter++
	if counter > 5 {
		runner.Stop()
	}
	runner.Data = counter
	return nil //fmt.Errorf("Exiting runner %s", runner.Name)
}
func ActionLoop3(ctx context.Context, runner *runners.Runner) error {
	runner.Logger.Info("action loop ", runner.Data)
	time.Sleep(time.Second * 1)
	counter := 0
	switch runner.Data.(type) {
	case int:
		counter = runner.Data.(int)
	}
	counter++
	if counter > 7 {
		runner.Stop()
	}
	runner.Data = counter
	return nil //fmt.Errorf("Exiting runner %s", runner.Name)
}
func RunRunners() {
	runners := runners.NewRunnerCollection()

	runners.CreateRunner("runner1", ActionLoop1, 0)
	runners.CreateRunner("runner2", ActionLoop2, 0)
	runners.CreateRunner("runner3", ActionLoop3, 0)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
	defer cancel()
	runners.RunAll(ctx)

}
