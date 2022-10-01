package discord

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"os"
)

type token struct {
	Token string `json:"discordToken"`
}

func DiscordConnect() (s *discordgo.Session, err error) {

	configJson, err := os.Open("config.json")
	if err != nil {
		log.Printf("Ocorreu um erro ao abrir o arquivo! %s\n", err.Error())
	}

	configBytes, err := ioutil.ReadAll(configJson)
	if err != nil {
		log.Printf("Ocorreu um erro ao converter para bytes! %s\n", err.Error())
	}

	var config token

	json.Unmarshal(configBytes, &config)

	return discordgo.New("Bot " + config.Token)
}