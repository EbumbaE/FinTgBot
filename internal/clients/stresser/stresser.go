package stresser

import (
	"context"
	"sync"
	"time"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/logger"
)

type Command struct {
	Text, Command, Arguments string
}

var commands = []Command{
	{Text: "/start", Command: "start"},
	{Text: "/setBudget", Command: "setBudget", Arguments: "10.2022 10000 USD"},
	{Text: "/setNote", Command: "setNote", Arguments: "12.10.2022 food 10"},
	{Text: "/setNote", Command: "setNote", Arguments: "31.10.2022 school 10"},
	{Text: "/getStatistic", Command: "getStatistic", Arguments: "week"},
	{Text: "/getStatistic", Command: "getStatistic", Arguments: "month"},
	{Text: "/getStatistic", Command: "getStatistic", Arguments: "year"},
}
var texts = []string{
	"USD",
	"RUB",
	"ebumba",
}

func StressTest(ctx context.Context, msgModel *messages.Model) {

	timeTicker := time.NewTicker(10 * time.Microsecond)
	go func() {
		logger.Info("stresser begin")
		for {
			select {
			case <-timeTicker.C:
				for _, text := range texts {
					for id := 1; id <= 20; id++ {
						msgModel.IncomingMessage(ctx, messages.Message{Text: text, UserID: int64(id)})
					}
				}
				for _, msg := range commands {
					for id := 1; id <= 20; id++ {
						msgModel.IncomingCommand(ctx,
							messages.Message{
								Text:      msg.Text,
								Command:   msg.Command,
								Arguments: msg.Arguments,
								UserID:    int64(id),
							})
					}
				}
			case <-ctx.Done():
				defer ctx.Value("allDoneWG").(*sync.WaitGroup).Done()
				logger.Info("stresser end")
				return
			}
		}
	}()
}
