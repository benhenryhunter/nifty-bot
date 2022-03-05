package discord

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/dickmanben/story-bot/types"
)

var channelEvents = map[string]types.Event{}

func getSession() (*discordgo.Session, error) {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		return nil, err
	}
	return dg, nil
}

func NewSession(c chan types.Event) {
	dg, err := getSession()
	if err != nil {
		fmt.Printf("error with discord session: %v", err)
		return
	}
	go AddHandler(c, dg)

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if _, ok := channelEvents[m.ChannelID]; ok && m.Author.ID != s.State.User.ID {
			event := channelEvents[m.ChannelID]
			switch t := event.Type; t {
			case "RandomNumber":
				finished := RandomNumberEvent(dg, event, m)
				if finished {
					delete(channelEvents, event.Channel)
				}
			case "CountUp":
				event, finished := CountUp(dg, event, m)
				if finished {
					delete(channelEvents, event.Channel)
				} else {
					channelEvents[event.Channel] = event
				}
			case "ConsecutiveEvent":
				event, finished := ConsecutiveEvent(dg, event, m)
				if finished {
					delete(channelEvents, event.Channel)
				} else {
					channelEvents[event.Channel] = event
				}
			}
		}
	})
	// dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMembers)

	// Open a websocket connection to Discord and begin listening.
	// defer dg.Close()
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
}

func AddHandler(c chan types.Event, dg *discordgo.Session) {
	for {
		select {
		case event := <-c:
			if event.Type == "RandomNumber" {
				event.Value = fmt.Sprintf("%v", rand.Intn(100))
				channelEvents[event.Channel] = event
				dg.ChannelMessageSend(event.Channel, "New Event Starting!  Guess the number between 0 and 100")
			}
			if event.Type == "CountUp" {
				channelEvents[event.Channel] = event
				dg.ChannelMessageSend(event.Channel, fmt.Sprintf("New Event Starting!  Count up to %v", event.Value))
			}
			if event.Type == "ConsecutiveEvent" {
				channelEvents[event.Channel] = event
				dg.ChannelMessageSend(event.Channel, event.StartAnnouncementContent)
			}
		}
	}
}
