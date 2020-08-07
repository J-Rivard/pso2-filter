package bot

import (
	"fmt"

	"github.com/J-Rivard/pso2-filter/internal/logging"

	"github.com/bwmarrin/discordgo"
)

// This function will be called every time a new
// message is created on any channel that the authenticated bot has access to.
func (b *Bot) messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if msg.Author.ID == session.State.User.ID {
		return
	}

	var embeds *discordgo.MessageEmbed
	if len(msg.Embeds) > 0 {
		embeds = msg.Embeds[0]
	}
	newMsg := &discordgo.MessageSend{
		Content: msg.Content,
		Embed:   embeds,
	}
	session.ChannelMessageSendComplex(msg.ChannelID, newMsg)

	for _, embed := range msg.Embeds {
		b.Log.LogDebug(logging.FormattedLog{
			"action":   "messageCreate",
			"metadata": fmt.Sprintf("%s/%s/%s\n", embed.Author, embed.Title, embed.Description),
		})
	}
}

func (b *Bot) messageUpdate(session *discordgo.Session, msg *discordgo.MessageUpdate) {
	if msg.Author == nil || msg.Author.ID == session.State.User.ID {
		return
	}

	fmt.Println(msg.ChannelID)

	for _, embed := range msg.Embeds {
		for _, field := range embed.Fields {
			b.Log.LogDebug(logging.FormattedLog{
				"action":   "messageUpdate",
				"metadata": fmt.Sprintf("%s\n", field.Value),
			})
		}
	}
}
