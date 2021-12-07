package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	guildCommand string = "801840788022624296"

	commands = []*discordgo.ApplicationCommand{
		{
			Name: "ping",

			Description: "Get the Bot's ping.",
		},

		{
			Name:        "set-rules",
			Description: "set server rules",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "rules",
					Description: "type rules",
					Required:    true,
				},
			},
		},

		{
			Name:        "rules",
			Description: "See server rules.",
		},

		{
			Name:        "clear",
			Description: "Removes latest messages.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "number-of-messages",
					Description: "Number of messages to delete(max 100)",
					Required:    false,
				},
			},
		},

		{

			Name:        "set-welcome",
			Description: "Enable welcome message on user join",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "enabled",
					Description: "Enable or disable the welcoming system.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "welcome-channel",
					Description: "The text channel you want your welcome messages send to them.",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "welcome-message",
					Description: "The text you want to send as welcome message.",
					Required:    false,
				},
			},
		},

		{
			Name:        "auto-role",
			Description: "give a special role to anyone that joins the crew!",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "enabled",
					Description: "Enable or disable auto-roling system.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The role you want to set.",
					Required:    false,
				},
			},
		},

		{
			Name:        "roll",
			Description: "rolling dice",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "dice",
					Description: "The dice",
					Required:    true,
				},
			},
		},

		{
			Name:        "help",
			Description: "Help",
		},
	}
)

func addCommands(session *discordgo.Session, commands []*discordgo.ApplicationCommand) {
	var commandNum = 0
	for _, v := range commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, guildCommand, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		} else {
			commandNum++
		}
	}
	log.Printf("Created %v commands", commandNum)

}

func deleteAllCommands(session *discordgo.Session) {
	commands, _ := session.ApplicationCommands(session.State.User.ID, guildCommand)
	var commandNum = 0
	for _, v := range commands {
		err := session.ApplicationCommandDelete(session.State.User.ID, guildCommand, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		} else {
			commandNum++
		}
	}
	log.Printf("Removed %v commands", commandNum)
}
