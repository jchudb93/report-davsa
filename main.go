package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var BOTID string

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
	token := os.Getenv("BOT_TOKEN")
	dgClient, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("No connection")
	}
	user, err := dgClient.User("@me")
	if err != nil {
		log.Fatal("No user")
	}

	dgClient.AddHandler(messageHandler)
	BOTID = user.ID

	err = dgClient.Open()
	if err != nil {
		log.Fatal("Can't open client")
	}

	fmt.Println("Bot is running!")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BOTID {
		return
	}

	if m.Content == "report davsa" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "AJ TMR REPORT DAVSA")
	}
	fmt.Println(m.Content)
}
