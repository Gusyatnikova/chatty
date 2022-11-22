package main

import (
	"context"

	"chatty/chatty/app"
)

func main() {
	chatService := app.NewChatty(context.Background())
	chatService.Run()
	chatService.ListenForShutdown()
}
