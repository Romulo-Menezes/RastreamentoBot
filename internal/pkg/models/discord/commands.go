package discord

import (
	"RastreioBot/internal/pkg/database"
	"RastreioBot/internal/pkg/models"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "add-package",
			Description: "Adicionar novo pacote para ser rastreado",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "tracking-code",
					Description: "Código de rastreamento",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Nome para identificar o pacote",
					Required:    true,
				},
			},
		},
		{
			Name:        "resume",
			Description: "Mostra o último status atualizado do seu pacote",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Nome que foi registrado o pacote",
					Required:    true,
				},
			},
		},
		{
			Name:        "remove",
			Description: "Remover um pacote",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Nome que foi registrado o pacote",
					Required:    true,
				},
			},
		},
		{
			Name:        "history",
			Description: "Histórico completo do pacote",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Nome que foi registrado o pacote",
					Required:    true,
				},
			},
		},
		{
			Name:        "list",
			Description: "Lista todos os seus pacotes registrados",
		},
		{
			Name:        "clear-chat",
			Description: "Limpa seu chat privado",
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"add-package": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.User == nil {
				privateAlert(s, i.Interaction)
				return
			}

			code := i.ApplicationCommandData().Options[0].StringValue()
			name := i.ApplicationCommandData().Options[1].StringValue()

			if !models.CheckCode(code) {
				errorMessage("O código de rastreio é inválido, verifique se você digitou certo!", s, i.Interaction)
				return
			}
			id := database.InsertPackage(i.User.ID, strings.ToUpper(name), strings.ToUpper(code))

			_, name, code = database.SelectByID(id)
			title, description := models.GetResume(name, code)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       title,
							Description: "Pacote adicionado com sucesso!",
							Color:       5763719,
						},
						{
							Title:       title,
							Description: description,
							Color:       5773779,
						},
					},
				},
			})
		},
		"resume": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.User == nil {
				privateAlert(s, i.Interaction)
				return
			}
			name := i.ApplicationCommandData().Options[0].StringValue()
			find, code := database.SelectByName(strings.ToUpper(name))
			if find {
				title, description := models.GetResume(name, code)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							{
								Title:       title,
								Description: description,
								Color:       5763719,
							},
						},
					},
				})
			} else {
				errorMessage("Ocorreu um erro ao tentar encontrar o pacote!", s, i.Interaction)
			}
		},
		"remove": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.User == nil {
				privateAlert(s, i.Interaction)
				return
			}
			name := i.ApplicationCommandData().Options[0].StringValue()

			if database.DeleteByName(strings.ToUpper(name)) {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							{
								Title:       "Pacote excluido!",
								Description: "Seu pacote foi excluido com sucesso!",
								Color:       5763719,
							},
						},
					},
				})
			} else {
				errorMessage("Ocorreu um erro ao tentar excluir o pacote!", s, i.Interaction)
			}
		},
		"history": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.User == nil {
				privateAlert(s, i.Interaction)
				return
			}
			fmt.Println("Histórico do pacote...")
			emConstrucao(s, i.Interaction)
		},
		"list": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.User == nil {
				privateAlert(s, i.Interaction)
				return
			}
			fmt.Println("listar pacotes...")
			emConstrucao(s, i.Interaction)
		},
		"clear-chat": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.User == nil {
				privateAlert(s, i.Interaction)
				return
			}
			emConstrucao(s, i.Interaction)
		},
	}
)

func privateAlert(s *discordgo.Session, i *discordgo.Interaction) {
	err := s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Chama no privado",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("Ocorreu um erro ao mandar o alerta privado: %v", err)
	}
}

func emConstrucao(s *discordgo.Session, i *discordgo.Interaction) {
	err := s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Em construção...",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("Ocorreu um erro ao mandar o alerta privado: %v", err)
	}
}

func errorMessage(description string, s *discordgo.Session, i *discordgo.Interaction) {
	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Ops, ocorreu um erro!",
					Description: description,
					Color:       15548997,
				},
			},
		},
	})
}
