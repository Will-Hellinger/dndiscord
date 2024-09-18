package discord

import (
	"log"
	"strconv"
	"strings"

	"github.com/Will-Hellinger/dndiscord/internal/die"
	"github.com/bwmarrin/discordgo"
)

func Run(token string, commandPrefix string) {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(func(discord *discordgo.Session, message *discordgo.MessageCreate) {
		NewMessage(discord, message, commandPrefix)
	})

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")
	select {}
}

func NewMessage(discord *discordgo.Session, message *discordgo.MessageCreate, commandPrefix string) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if !strings.HasPrefix(message.Content, commandPrefix) {
		return
	}

	content := strings.TrimPrefix(message.Content, commandPrefix)
	args := strings.Fields(content)

	// Ensure there is at least one argument after the prefix
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "ping":
		discord.ChannelMessageSend(message.ChannelID, "Pong!")
	case "roll":
		// Handle the roll command
		if len(args) < 2 {
			discord.ChannelMessageSend(message.ChannelID, "Invalid roll command. Please provide the dice (e.g., d10, d20).")
			return
		}

		// Split all dice arguments by commas and spaces
		diceArgs := strings.Split(strings.Join(args[1:], " "), ",")

		// Trim spaces and roll dice for each valid argument
		for _, diceArg := range diceArgs {
			diceArg = strings.TrimSpace(diceArg) // Remove extra spaces

			if !strings.HasPrefix(diceArg, "d") || len(diceArg) < 2 {
				discord.ChannelMessageSend(message.ChannelID, "Invalid dice format: "+diceArg+". Use the format dX, where X is the dice size (e.g., d10).")
				continue
			}

			diceSize, err := strconv.Atoi(strings.TrimPrefix(diceArg, "d"))
			if err != nil || diceSize <= 0 {
				discord.ChannelMessageSend(message.ChannelID, "Invalid dice size for: "+diceArg+". Please provide a valid number after 'd'.")
				continue
			}

			roll := die.Roll(diceSize)
			discord.ChannelMessageSend(message.ChannelID, "You rolled a "+strconv.Itoa(roll)+" for "+diceArg)
		}
	}
}
