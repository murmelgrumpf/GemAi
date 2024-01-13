package utils

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var EmbedColorInfo = 0x2894a0
var EmbedColorError = 0xff1900

func BoolButtonStye(enabled bool) discordgo.ButtonStyle {
	if enabled {
		return discordgo.SuccessButton
	}
	return discordgo.DangerButton
}

func DefaultEmoji(emoji string) discordgo.ComponentEmoji {
	return discordgo.ComponentEmoji{
		ID:   "",
		Name: emoji,
	}
}

func createErrorEmbed(errMsg string) *[]*discordgo.MessageEmbed {
	stacktrace := "```\n" + errMsg + "\n```\n"
	pc := make([]uintptr, 20) // at least 1 entry needed
	stackLen := runtime.Callers(2, pc)
	for iter := 0; iter < stackLen; iter++ {
		f := runtime.FuncForPC(pc[iter])
		file, line := f.FileLine(pc[iter])
		if !strings.Contains(file, "GemAi") {
			break
		}
		fnSlice := strings.Split(f.Name(), "/")
		stacktrace += "**" + fnSlice[len(fnSlice)-1] + "**\n" + strings.Split(file, "GemAi/")[1] + ":" + strconv.Itoa(line) + "\n\n"
	}

	embed := []*discordgo.MessageEmbed{
		{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Function could not be executed",
			Description: stacktrace,
			Color:       EmbedColorError,
		},
	}
	return &embed
}

func SendError(s *discordgo.Session, i *discordgo.InteractionCreate, errMsg string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: *createErrorEmbed(errMsg),
		},
	})
}

func EditMessageError(s *discordgo.Session, i *discordgo.InteractionCreate, errMsg string) {
	empty := ""
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &empty,
		Embeds:  createErrorEmbed(errMsg),
	})
}

func InteractionRespond(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	resp *discordgo.InteractionResponse,
	options ...discordgo.RequestOption,
) {
	err := s.InteractionRespond(i.Interaction, resp, options...)
	if err != nil {
		SendError(s, i, err.Error())
	}
}

func ChannelMessageEditComplex(s *discordgo.Session, i *discordgo.InteractionCreate, me *discordgo.MessageEdit) *discordgo.Message {
	msg, err := s.ChannelMessageEditComplex(me)
	if err != nil {
		SendError(s, i, err.Error())
	}
	return msg
}
