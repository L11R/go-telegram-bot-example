package main // import "awesome-project"

import (
	"awesome-project/config"
	"awesome-project/database"
	"awesome-project/handlers"
	"flag"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"os/signal"
)

var (
	path = flag.String(
		"config",
		"",
		"enter path to config file",
	)
)

func main() {
	// Parse at first startup
	flag.Parse()

	// Read config
	conf, err := config.NewConfig(*path)
	if err != nil {
		fmt.Println("Incorrect path or config itself! See help.")
		os.Exit(2)
	}

	bot, err := tgbotapi.NewBotAPI("677581835:AAF1wrTOns5-XHJoePMC44Xa7NEwD3okMsk")
	if err != nil {
		fmt.Println("Telegram bot cannot be initialized! See, error:")
		panic(err)
	}

	fmt.Printf("Authorized on account @%s\n", bot.Self.UserName)

	// Init database
	db, err := database.NewDatabase(conf)
	if err != nil {
		fmt.Println("Database cannot be initialized! See, error:")
		panic(err)
	}

	// Try auto migration for first start
	err = db.AutoMigrate()
	if err != nil {
		fmt.Println("Cannot auto migrate! See, error:")
		panic(err)
	}

	var h handlers.Handler
	h.Bot = bot
	h.DB = db.Conn

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Graceful shutdown
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, os.Kill)

	go func() {
		<-s
		updates.Clear()
		os.Exit(1)
	}()

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		switch update.Message.Command() {
		case "ping":
			h.Pong(update)
		}
	}
}
