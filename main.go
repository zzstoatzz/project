package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// read env var
func load_token() string {
	return os.Getenv("DISCORD_CLIENT_SECRET")
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + load_token())
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

    messageCreate := func(s *discordgo.Session, m *discordgo.MessageCreate) {
        // Ignore all messages created by the bot itself
        if m.Author.ID == s.State.User.ID {
            return
        }

		if _, err := s.ChannelMessageSend(m.ChannelID, "Hello"); err != nil {
			fmt.Println("error sending handling", err)
		}
	}
    

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}
