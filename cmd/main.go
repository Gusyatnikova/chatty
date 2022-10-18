package main

import (
	"context"

	"chatty/chatty/app"
)

func main() {
	chatService := app.NewChatService(context.Background())
	//todo: run or
	chatService.Run()
	chatService.ListenForShutdown()
}
