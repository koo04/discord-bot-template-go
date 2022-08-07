package commands

import (
	"github.com/bwmarrin/discordgo"
)

// HelloWorld command handler
func HelloWorld(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello World!",
			Flags:   uint64(discordgo.MessageFlagsEphemeral),
		},
	}); err != nil {
		panic(err)
	}
}
