package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	BotToken string
	//OpenWeatherToken string
)

func NewMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	log.Printf("AuthorID: %s,UserID:%s \n", message.Author.ID, discord.State.User.ID)
	//忽略机器人的消息
	if message.Author.ID == discord.State.User.ID {
		return
	}
	fmt.Println(message.Content)
	//消息回复
	switch {
	case strings.Contains(message.Content, "weather"):
		discord.ChannelMessageSend(message.ChannelID, "I can help with that!")
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Hi there!")
	}

}

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Println("error creating Discord session,", err)
	}

	discord.Identify.Intents = discordgo.IntentDirectMessages

	discord.AddHandler(NewMessage)
	if err := discord.Open(); err != nil {
		log.Println("error opening connection,", err)
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	defer func() {
		if ok := discord.Close(); ok != nil {
			log.Println("discord close err:", err)
		}
	}()
}
