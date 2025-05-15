package main

import (
	"asyasocute/pills/config"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	config.Load()
	pref := tele.Settings{
		Token:  config.C.BotApiToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
