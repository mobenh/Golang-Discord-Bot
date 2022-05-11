package main

import (
	"fmt"

	"github.com/mobenh/Golang-Discord-Bot/bot"
	"github.com/mobenh/Golang-Discord-Bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
