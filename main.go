package main

import (
	"archroid/ElProfessorBot/bot"
	"archroid/ElProfessorBot/database"
	"os"
)

func main() {

	token := os.Getenv("DISCORD_BOT_TOKEN")

	db := database.Start()
	bot.Start(token, db)

	// <-make(chan struct{})
	// return
}
