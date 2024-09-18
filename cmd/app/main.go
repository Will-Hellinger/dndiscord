package main

import (
	"log"

	"github.com/Will-Hellinger/dndiscord/internal/config"
	"github.com/Will-Hellinger/dndiscord/internal/discord"
)

func main() {
	log.Println("Hello, World!")
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	discord.Run(config.Token, config.CommandPrefix)
}
