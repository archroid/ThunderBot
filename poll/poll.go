package poll

import (
	"log"

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

	// var poll Poll

	// topic := i.ApplicationCommandData().Options[0].StringValue()
	// duration := i.ApplicationCommandData().Options[1].IntValue()

	optionsNum := len(i.ApplicationCommandData().Options) - 2

	option1 := i.ApplicationCommandData().Options[2].StringValue()
	option2 := i.ApplicationCommandData().Options[3].StringValue()
	option3 := i.ApplicationCommandData().Options[4].StringValue()
	
	option4 := i.ApplicationCommandData().Options[5].StringValue()
	option5 := i.ApplicationCommandData().Options[6].StringValue()
	option6 := i.ApplicationCommandData().Options[7].StringValue()
	option7 := i.ApplicationCommandData().Options[8].StringValue()
	option8 := i.ApplicationCommandData().Options[9].StringValue()
	option9 := i.ApplicationCommandData().Options[10].StringValue()
	option10 := i.ApplicationCommandData().Options[11].StringValue()

	optionsArray := [10]string{option1, option2, option3, option4, option5, option6, option7, option8, option9, option10}
	optionsSlice := optionsArray[2:optionsNum]
	log.Println(len(optionsSlice))

}
