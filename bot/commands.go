package bot

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
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

		{
			Name:        "join",
			Description: "Join to the voice channel.",
		},

		{
			Name:        "disconnect",
			Description: "Disconnect from the voice channel.",
		},
		{
			Name:        "play",
			Description: "Play the youtube video.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "music",
					Description: "Enter music name or URL",
					Required:    true,
				},
			},
		},
		{
			Name:        "stop",
			Description: "Stop the current playing song",
		},

		{
			Name:        "search",
			Description: "Search for video from youtube",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "search",
					Description: "search query",
					Required:    true,
				},
			},
		},

		{
			Name:        "poll",
			Description: "Create a new poll",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "topic",
					Description: "Description of the poll",
					Required:    true,
				},

				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "duration",
					Description: "Poll duration in minutes (0 for unlimited)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "option1",
					Description: "Option 1",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "option2",
					Description: "Option 2",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "option3",
					Description: "Option 3",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "option4",
					Description: "Option 4",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "option5",
					Description: "Option 5",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "option6",
					Description: "Option 6",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "option7",
					Description: "Option 7",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "option8",
					Description: "Option 8",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "option9",
					Description: "Option 9",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "option10",
					Description: "Option 10",
					Required:    false,
				},
			},
		},

		{
			Name:        "notes",
			Description: "Get all notes",
		},
		{
			Name:        "get-note",
			Description: "Get a note",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "note",
					Description: "Note name",
					Required:    true,
				},
			},
		},
		{
			Name:        "delete-note",
			Description: "Delete a note",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "note",
					Description: "Note name",
					Required:    true,
				},
			}},
		{
			Name:        "add-note",
			Description: "Add a note",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "note-name",
					Description: "Note name",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "note-content",
					Description: "Note content",
					Required:    true,
				},
			},
		},
	}
)

func addCommands(session *discordgo.Session, commands []*discordgo.ApplicationCommand) {
	var commandNum = 0
	for _, v := range commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		} else {
			commandNum++
		}
	}
	log.Printf("Created %v commands", commandNum)

}

func deleteAllCommands(session *discordgo.Session) {
	commands, _ := session.ApplicationCommands(session.State.User.ID, "")
	var commandNum = 0
	for _, v := range commands {
		err := session.ApplicationCommandDelete(session.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		} else {
			commandNum++
		}
	}
	log.Printf("Removed %v commands", commandNum)
}
