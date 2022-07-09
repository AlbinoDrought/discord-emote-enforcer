package main

import (
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/bwmarrin/discordgo"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

var (
	guildID, channelID string
)

var (
	exprEmotes     = regexp.MustCompile(`(<a?:[^:]+:\d+>)`)
	exprWhitespace = regexp.MustCompile(`(\s+)`)
)

func main() {
	authenticationToken := os.Getenv("DEC_TOKEN")
	guildID = os.Getenv("DEC_GUILD_ID")
	channelID = os.Getenv("DEC_CHANNEL_ID")

	if authenticationToken == "" || guildID == "" || channelID == "" {
		log.Fatal("require DEC_TOKEN, DEC_GUILD_ID, DEC_CHANNEL_ID")
	}

	session, err := discordgo.New("Bot " + authenticationToken)
	if err != nil {
		log.Fatal("failed to create discord session: ", err)
	}
	session.AddHandler(ready)
	session.AddHandler(messageCreate)

	session.Identify.Intents = discordgo.IntentsGuildMessages

	if err := session.Open(); err != nil {
		log.Fatal("failed to open discord session: ", err)
	}
	defer session.Close()

	log.Println("I'm running 😊")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	log.Println("I'm closing 😢")
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "hi")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.GuildID != guildID || m.ChannelID != channelID {
		return
	}

	hasAttachments := len(m.Attachments) > 0
	hasNonEmoteText := textContainsNonEmotes(m.Content)

	shouldBeRemoved := hasAttachments || hasNonEmoteText

	if !shouldBeRemoved {
		return
	}

	log.Printf("deleting message %v with %v attachments, content %v", m.ID, len(m.Attachments), m.Content)
	if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
		log.Printf("failed to delete message %v: %v", m.ID, err)
	}
}

func textContainsNonEmotes(messageText string) bool {
	// remove all discord emotes
	messageText = exprEmotes.ReplaceAllString(messageText, "")
	// remove all emojis
	messageText = emoji.RemoveAll(messageText)
	// remove all whitespace
	messageText = exprWhitespace.ReplaceAllString(messageText, "")

	// if there's anything left, this message had other text in it
	return messageText != ""
}
