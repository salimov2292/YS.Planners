package bot

import (
	"planners/supabase"
	"planners/types"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
)

type Bot struct {
	client *supabase.Client
	bot    *telebot.Bot
}

func NewBot(botToken string, supabaseClient *supabase.Client) (*Bot, error) {
	b, err := telebot.NewBot(telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	return &Bot{
		client: supabaseClient,
		bot:    b,
	}, nil
}

func (b *Bot) Start() {
	b.bot.Handle(telebot.OnText, func(c telebot.Context) error {
		message := c.Message().Text
		tasks := strings.Split(message, "\n")

		for _, taskLine := range tasks {
			if strings.Contains(taskLine, " ") {
				msgParts := strings.SplitN(taskLine, " ", 2)
				tag := msgParts[0]
				if !strings.HasPrefix(tag, "#") {
					continue
				}
				description := msgParts[1]

				task := types.Tab{
					Tag:         tag,
					Description: description,
				}

				err := b.client.InsertTask(task)
				if err != nil {
					return c.Send("Ошибка при создании задачи: " + err.Error())
				}
			}
		}

		return c.Send("Задачи были успешно сохранены!")
	})

	b.bot.Start()
}
