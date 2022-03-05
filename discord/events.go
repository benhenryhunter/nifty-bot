package discord

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/dickmanben/story-bot/types"
	"github.com/dickmanben/story-bot/utils"
)

func RandomNumberEvent(dg *discordgo.Session, event types.Event, m *discordgo.MessageCreate) bool {
	channelID := event.Channel
	if m.ChannelID == channelID && m.Message.Content == event.Value {
		dg.ChannelMessageSend(channelID, "You guessed the right number!")
		return true
	}
	return false
}

func ConsecutiveEvent(dg *discordgo.Session, event types.Event, m *discordgo.MessageCreate) (types.Event, bool) {
	channelID := event.Channel
	currentValue, err := strconv.Atoi(event.CurrentValue)
	if err != nil {
		currentValue = 0
	}
	nextMessage := event.MessageProgression[currentValue]
	// if m.Message.Author.Username == event.LastSender {
	// 	dg.MessageReactionAdd(m.ChannelID, m.Message.ID, "❌")
	// 	dg.MessageReactionAdd(m.ChannelID, m.Message.ID, "person_facepalming")
	// 	dg.ChannelMessageSend(channelID, event.RestartContent)
	// 	event.CurrentValue = "0"
	// 	event.LastSender = ""
	// 	return event, false
	// }

	pass, err := utils.CompareDiscordMessageWithEventMessage(dg, m.Message, nextMessage)
	if err != nil {
		switch err.Error() {
		case "missing role":
			dg.MessageReactionAdd(m.ChannelID, m.Message.ID, "❌")
			dg.ChannelMessageSend(channelID, fmt.Sprintf("missing required role %v", event.RestartContent))
			event.CurrentValue = "0"
			event.LastSender = ""
			return event, false
		case "wrong message":
			dg.MessageReactionAdd(m.ChannelID, m.Message.ID, "❌")
			dg.ChannelMessageSend(channelID, fmt.Sprintf("wrong message %v", event.RestartContent))
			event.CurrentValue = "0"
			event.LastSender = ""
			return event, false
		}
	}

	if pass {
		dg.MessageReactionAdd(m.ChannelID, m.Message.ID, "✅")
		if currentValue == len(event.MessageProgression)-1 {
			dg.ChannelMessageSend(channelID, event.EndAnnouncementContent)
			return event, true
		}
		event.CurrentValue = fmt.Sprintf("%v", currentValue+1)
		event.LastSender = m.Message.Author.Username
		return event, false
	}
	return event, false
}

func CountUp(dg *discordgo.Session, event types.Event, m *discordgo.MessageCreate) (types.Event, bool) {
	channelID := event.Channel
	currentValue, err := strconv.Atoi(event.CurrentValue)
	if err != nil {
		currentValue = 0
	}
	rightNextValue := fmt.Sprintf("%v", currentValue+1)
	if m.Message.Author.Username == event.LastSender {
		dg.MessageReactionAdd(m.ChannelID, m.Message.ID, "❌")
		dg.MessageReactionAdd(m.ChannelID, m.Message.ID, "person_facepalming")
		dg.ChannelMessageSend(channelID, fmt.Sprintf("%v!!!! You can't send twice during the count.  Starting back at 0!", event.LastSender))
		event.CurrentValue = "0"
		event.LastSender = ""
		return event, false
	}
	if m.Message.Content != rightNextValue {
		dg.MessageReactionAdd(m.ChannelID, m.Message.ID, "❌")
		dg.ChannelMessageSend(channelID, fmt.Sprintf("%v hecking learn to count, starting over.", event.LastSender))
		event.CurrentValue = "0"
		event.LastSender = ""
		return event, false
	}
	if m.ChannelID == channelID && rightNextValue == m.Message.Content {
		dg.MessageReactionAdd(m.ChannelID, m.Message.ID, "✅")
		if rightNextValue == event.Value {
			dg.ChannelMessageSend(channelID, fmt.Sprintf("YOU DID IT!  WE COUNTED ALL THE WAY UP TO %v", event.Value))
			return event, true
		}
		event.CurrentValue = fmt.Sprintf("%v", currentValue+1)
		event.LastSender = m.Message.Author.Username
		return event, false
	}
	return event, false
}
