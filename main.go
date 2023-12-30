package main

//https://www.youtube.com/watch?v=i4nX0s8WHEE&list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9&index=11
//Tutorial

import (
	"fmt"

	"github.com/Michalis98/discordbot/bot"
	"github.com/Michalis98/discordbot/config"
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
