package main

import (
	"context"
	"os"
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
)

func main()  {
	os.Setenv("TOKEN", "Your-Discord-Bot-Token")

	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("TOKEN"),
	})

	defer client.Gateway().StayConnectedUntilInterrupted()

	fmt.Println("The bot has connected!")

	content, _ := std.NewMsgFilter(context.Background(), client)
	content.SetPrefix("ping")

	client.Gateway().WithMiddleware(content.HasPrefix).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
		_, _ = h.Message.Reply(context.Background(), s, "Pong!")
	})
}