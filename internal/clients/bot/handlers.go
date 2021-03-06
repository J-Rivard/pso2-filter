package bot

import (
	"fmt"

	"github.com/J-Rivard/pso2-filter/internal/logging"

	"github.com/bwmarrin/discordgo"
)

const (
	Prefix        = "$"
	Subscribe     = "SUBSCRIBE"
	ReadChannelId = "741095315968491604"
)

// This function will be called every time a new
// message is created on any channel that the authenticated bot has access to.
func (b *Bot) messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	// Handle commands from all channels
	b.handleCommand(msg)

	// If its the channel we're reading from
	if msg.ChannelID == ReadChannelId {
		b.handleEvent(session, msg)
	}
}

// This function will be called every time a
// message is updated on any channel that the authenticated bot has access to.
func (b *Bot) messageUpdate(session *discordgo.Session, msg *discordgo.MessageUpdate) {
	if msg.Author == nil || msg.Author.ID == session.State.User.ID {
		return
	}

	for _, embed := range msg.Embeds {
		for _, field := range embed.Fields {
			b.Log.LogDebug(logging.FormattedLog{
				"action":   "messageUpdate",
				"metadata": fmt.Sprintf("%s\n", field.Value),
			})
		}
	}
}
