package discord

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dickmanben/story-bot/utils"
)

func memberUpdate(s *discordgo.Session, u *discordgo.GuildMemberUpdate) {
	if u.Member.User.Bot {
		return
	}
	if u.Member.Nick == "" {
		s.ChannelMessageSend(os.Getenv("WELCOME_CHANNEL_ID"), "NOT A JEPH?!?!?!")
		name := fmt.Sprintf("jeph - %v", utils.StringWithCharset(5))
		s.GuildMemberNickname(u.GuildID, u.User.ID, name)
		s.ChannelMessageSend(os.Getenv("WELCOME_CHANNEL_ID"), fmt.Sprintf("I dub you %v", name))
		s.GuildMemberRoleAdd(u.GuildID, u.User.ID, os.Getenv("ARMY_ROLE_ID"))
		return
	}
	if !strings.Contains(strings.ToLower(u.Member.Nick), "jeph") {
		name := fmt.Sprintf("traitor jeph - %v", utils.StringWithCharset(5))
		s.ChannelMessageSend(os.Getenv("GENERAL_CHANNEL_ID"), fmt.Sprintf("Nuh uh uhhh! <@%v> We are many. We are one.  You shall now be: %v", u.User.ID, name))
		s.GuildMemberNickname(u.GuildID, u.User.ID, name)
		s.GuildMemberRoleRemove(u.GuildID, u.User.ID, os.Getenv("ARMY_ROLE_ID"))
		s.GuildMemberRoleAdd(u.GuildID, u.User.ID, os.Getenv("TRAITOR_ID"))
	}
}

func messageSent(s *discordgo.Session, m *discordgo.MessageCreate) {
	// check if an event is live right now

}
