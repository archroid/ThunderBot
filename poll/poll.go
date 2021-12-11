package poll

import (
	"github.com/bwmarrin/discordgo"
)

type Poll struct {
	topic             string
	options           []string
	durationInMinutes int
}

var numberToEmoji = map[int]string{
	1:  "1️⃣",
	2:  "2️⃣",
	3:  "3️⃣",
	4:  "4️⃣",
	5:  "5️⃣",
	6:  "6️⃣",
	7:  "7️⃣",
	8:  "8️⃣",
	9:  "9️⃣",
	10: "🔟",
}

func CreatePoll(s *discordgo.Session, i *discordgo.InteractionCreate) {

	var poll Poll

	topic := i.ApplicationCommandData().Options[0].StringValue()
	duration := i.ApplicationCommandData().Options[1].IntValue()

	optionsNum := len(i.ApplicationCommandData().Options) - 2

	option1 := i.ApplicationCommandData().Options[2].StringValue()
	option2 := i.ApplicationCommandData().Options[3].StringValue()
	optionsArray := []string{option1, option2}

	if optionsNum == 5 {
		optionsArray = append(optionsArray, i.ApplicationCommandData().Options[4].StringValue())
	}

	if optionsNum == 6 {
		optionsArray = append(optionsArray, i.ApplicationCommandData().Options[5].StringValue())
	}
	if optionsNum == 7 {
		optionsArray = append(optionsArray, i.ApplicationCommandData().Options[6].StringValue())
	}
	if optionsNum == 8 {
		optionsArray = append(optionsArray, i.ApplicationCommandData().Options[7].StringValue())
	}
	if optionsNum == 9 {
		optionsArray = append(optionsArray, i.ApplicationCommandData().Options[8].StringValue())
	}
	if optionsNum == 10 {
		optionsArray = append(optionsArray, i.ApplicationCommandData().Options[9].StringValue())
	}
	if optionsNum == 11 {
		optionsArray = append(optionsArray, i.ApplicationCommandData().Options[10].StringValue())
	}
	if optionsNum == 12 {
		optionsArray = append(optionsArray, i.ApplicationCommandData().Options[11].StringValue())
	}

	poll = Poll{topic: topic, durationInMinutes: int(duration), options: optionsArray}

	println(len(poll.options))

}
