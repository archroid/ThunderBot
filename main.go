package main

import (
	"fmt"
	"test/Discord-Template/bot"
	"test/Discord-Template/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
