package bot

import (
	"fmt"
	"strings"

	"github.com/J-Rivard/pso2-filter/internal/logging"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handleCommand(msg *discordgo.MessageCreate) {
	if len(msg.Content) < 1 {
		return
	}

	if string(msg.Content[0]) == Prefix {
		tokenized := strings.Split(msg.Content, Prefix)
		if len(tokenized) < 2 {
			return
		}

		switch strings.ToUpper(tokenized[1]) {
		case Subscribe:
			err := b.addSubscription(msg)
			if err != nil {
				b.Log.LogDebug(logging.FormattedLog{
					"action":   "db_update",
					"error":    err.Error(),
					"metadata": fmt.Sprintf("Failed to insert channelID %s", msg.ChannelID),
				})
			}
		}
	}
}

func (b *Bot) handleEvent(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.Username != BotAuthor {
		return
	}

	for _, event := range b.Database.Events {
		for _, embed := range msg.Embeds {
			if strings.Contains(embed.Description, *event) {
				newMsg := &discordgo.MessageSend{
					Content: msg.Content,
					Embed:   embed,
				}

				err := b.notifySubscriptions(session, newMsg)
				if err != nil {
					b.Log.LogDebug(logging.FormattedLog{
						"action":   "handleEvent",
						"error":    err.Error(),
						"metadata": "Failed to notify",
					})
				}

				return
			}
		}
	}
}

func (b *Bot) addSubscription(msg *discordgo.MessageCreate) error {
	err := b.Database.AddSubscription(msg.ChannelID)
	if err != nil {
		return err
	}

	b.Log.LogDebug(logging.FormattedLog{
		"action":   "db_update",
		"metadata": fmt.Sprintf("Inserted channelID %s", msg.ChannelID),
	})

	return nil
}

func (b *Bot) notifySubscriptions(session *discordgo.Session, msg *discordgo.MessageSend) error {
	subs, err := b.Database.FetchSubscriptions()
	if err != nil {
		return err
	}

	for _, sub := range subs {
		_, err := session.ChannelMessageSendComplex(*sub.Channel, msg)
		if err != nil {
			return err
		}
	}

	return nil
}
