package tgbot

import (
	"context"
	"fmt"

	"github.com/guionardo/go-tgbot/tgbot/runners"
)

func BotWorkerRunnerAction(ctx context.Context, runner *runners.Runner) error {
	svc := GetBotService(ctx)
	wrk := svc.worker
	if wrk.schedules.Count() == 0 {
		runner.Stop()
		return fmt.Errorf("no schedules to run")
	}

	nextSchedule := wrk.schedules.GetNextSchedule()
	nextSchedule.WaitUntilNextRunRound()
	nextSchedule.DoAction(ctx)
	return nil
}
