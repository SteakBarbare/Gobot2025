package main

import (
	"fmt"

	"github.com/SteakBarbare/Gobot2025/bot"
	"github.com/SteakBarbare/Gobot2025/config"
)

func main() {
	err := config.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

}
