package commands

import (
	"archroid/ElProfessorBot/pkg/fetch"
	"archroid/ElProfessorBot/static"
	"archroid/ElProfessorBot/utils"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/zekroTJA/shireikan"
)

type CmdClear struct {
}

func (c *CmdClear) GetInvokes() []string {
	return []string{"clear", "c", "delete"}
}

func (c *CmdClear) GetDescription() string {
	return "Clear messages in a channel."
}

func (c *CmdClear) GetHelp() string {
	return "`clear` - delete last 100 messages\n" +
		"`clear <n>` - clear an amount of messages\n" +
		"`clear <n> <userResolvable>` - clear an amount of messages by a specific user\n"

}

func (c *CmdClear) GetGroup() string {
	return shireikan.GroupModeration
}

func (c *CmdClear) GetDomainName() string {
	return ""
}

func (c *CmdClear) GetSubPermissionRules() []shireikan.SubPermission {
	return nil
}

func (c *CmdClear) IsExecutableInDMChannels() bool {
	return true
}

func (c *CmdClear) Exec(ctx shireikan.Context) error {
	var msglist []*discordgo.Message
	var err error

	if len(ctx.GetArgs()) == 0 {
		msglist, err = ctx.GetSession().ChannelMessages(ctx.GetChannel().ID, 100, "", "", "")
	} else {

		var memb *discordgo.Member
		n, err := ctx.GetArgs().Get(0).AsInt()
		if err != nil {
			return utils.SendEmbedError(ctx.GetSession(), ctx.GetChannel().ID,
				"Sorry, but the member can not be found on this guild. :cry:").
				DeleteAfter(8 * time.Second).Error()
		} else if n < 0 || n > 99 {
			return utils.SendEmbedError(ctx.GetSession(), ctx.GetChannel().ID,
				"Number of messages is invald and must be between *(including)* 0 and 100.").
				DeleteAfter(8 * time.Second).Error()
		}

		// Account for command message itself
		n++

		if len(ctx.GetArgs()) >= 2 {
			memb, err = fetch.FetchMember(ctx.GetSession(), ctx.GetGuild().ID, ctx.GetArgs().Get(1).AsString())
			if err != nil {
				return utils.SendEmbedError(ctx.GetSession(), ctx.GetChannel().ID,
					"Sorry, but the member can not be found on this guild. :cry:").
					DeleteAfter(8 * time.Second).Error()
			}
		}
		msglistUnfiltered, err := ctx.GetSession().ChannelMessages(ctx.GetChannel().ID, n, "", "", "")
		if err != nil {
			return err
		}

		if memb != nil {
			for _, m := range msglistUnfiltered {
				if m.Author.ID == memb.User.ID {
					msglist = append(msglist, m)
				}
			}
		} else {
			msglist = msglistUnfiltered
		}
	}

	if err != nil {
		return err
	}

	var messageIds []string
	for _, strings := range msglist {
		messageTimestamp := strings.Timestamp.Unix()

		twoWeeksTimestampUnix := time.Now().AddDate(0, 0, -14).Unix()

		if messageTimestamp >= twoWeeksTimestampUnix {
			messageIds = append(messageIds, strings.Reference().MessageID)
		}

	}

	err = ctx.GetSession().ChannelMessagesBulkDelete(ctx.GetChannel().ID, messageIds)
	if err != nil {
		return utils.SendEmbedError(ctx.GetSession(), ctx.GetChannel().ID, "You can only bulk delete messages that are under 14 days old.").
			DeleteAfter(12 * time.Second).Error()
	}

	return utils.SendEmbed(ctx.GetSession(), ctx.GetChannel().ID,
		fmt.Sprintf("Deleted %d %s.", len(messageIds)-1, utils.Pluralize(len(messageIds)-1, "message")), "", static.ColorEmbedUpdated).
		DeleteAfter(6 * time.Second).Error()
}
