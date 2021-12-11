package poll

import (
	"fmt"

	embed "archroid/ElProfessorBot/utils"

	"github.com/bwmarrin/discordgo"
)

func CreatePoll(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// var poll Poll

	topic := i.ApplicationCommandData().Options[0].StringValue()

	optionsNum := len(i.ApplicationCommandData().Options) - 2

	option1 := i.ApplicationCommandData().Options[2].StringValue()
	option2 := i.ApplicationCommandData().Options[3].StringValue()

	pollContent := fmt.Sprintf("📊** %v ** \n \n 1️⃣ %v \n 2️⃣ %v", topic, option1, option2)

	if optionsNum >= 3 {
		pollContent = pollContent + fmt.Sprintf("\n3️⃣ %v ", i.ApplicationCommandData().Options[4].StringValue())
		// println(i.ApplicationCommandData().Options[4].StringValue())
		// println(pollContent)
	}
	if optionsNum >= 4 {
		pollContent = pollContent + fmt.Sprintf("\n4️⃣ %v ", i.ApplicationCommandData().Options[5].StringValue())
	}
	if optionsNum >= 5 {
		pollContent = pollContent + fmt.Sprintf("\n5️⃣ %v ", i.ApplicationCommandData().Options[6].StringValue())
	}
	if optionsNum >= 6 {
		pollContent = pollContent + fmt.Sprintf("\n6️⃣ %v ", i.ApplicationCommandData().Options[7].StringValue())

	}
	if optionsNum >= 7 {
		pollContent = pollContent + fmt.Sprintf("\n7️⃣ %v ", i.ApplicationCommandData().Options[8].StringValue())

	}
	if optionsNum >= 8 {
		pollContent = pollContent + fmt.Sprintf("\n8️⃣ %v ", i.ApplicationCommandData().Options[9].StringValue())

	}
	if optionsNum >= 9 {
		pollContent = pollContent + fmt.Sprintf("\n9️⃣ %v ", i.ApplicationCommandData().Options[10].StringValue())

	}
	if optionsNum == 10 {
		pollContent = pollContent + fmt.Sprintf("\n🔟 %v ", i.ApplicationCommandData().Options[11].StringValue())

	}

	embed := embed.NewEmbed().
		SetColor(0x66F442).
		SetAuthor(i.Member.User.Username, i.Member.User.AvatarURL("24")).
		SetDescription(pollContent).
		MessageEmbed

	embeds := []*discordgo.MessageEmbed{embed}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
		},
	})

	// time.Sleep(time.Duration(duration))
	// println(time.Duration(duration).Minutes())

}
