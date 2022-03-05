package utils

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dickmanben/story-bot/types"
)

func CompareDiscordMessageWithEventMessage(dg *discordgo.Session, dm *discordgo.Message, em types.Message) (bool, error) {
	if em.RequiredRole != "" {
		if t, err := compareRoles(dg, dm, em); t == false || err != nil {
			return false, errors.New("missing role")
		}
	}

	// if strings.ToLower(em.Content) != strings.ToLower(dm.Content) {
	if strings.Index(em.Content, dm.Content) > 0 {
		return false, errors.New("wrong message")
	}
	return true, nil
}

func compareRoles(dg *discordgo.Session, dm *discordgo.Message, em types.Message) (bool, error) {
	member, err := dg.GuildMember(dm.GuildID, dm.Author.ID)
	if err != nil {
		return false, err
	}
	g, err := dg.Guild(dm.GuildID)
	if err != nil {
		return false, err
	}
	roleMap := map[string]string{}
	for _, role := range g.Roles {
		roleMap[role.ID] = role.Name
	}
	for _, role := range member.Roles {
		if roleMap[role] == em.RequiredRole {
			return true, nil
		}
	}
	return false, nil
}
