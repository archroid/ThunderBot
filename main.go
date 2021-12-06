package main

import (
	"archroid/ElProfessorBot/bot"
	"os"
)

func main() {

	token := os.Getenv("DISCORD_BOT_TOKEN")

	bot.Start(token)

	// <-make(chan struct{})
	return
}
