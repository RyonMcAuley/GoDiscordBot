package main

import (
	_ "embed"
	"strings"

	bot "example.com/roboanna/Bot"
)

//go:embed token.txt
var token string

func main() {
	token = strings.TrimSpace(token)
	bot.BotToken = token
	bot.Run()
}
