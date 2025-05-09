package main

import (
	_ "embed"
	"fmt"

	bot "example.com/roboanna/Bot"
)

//go:embed token.txt
var token string

func main() {
	//read token from file :P
	// dat, err := os.ReadFile("token.txt")
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Token: ", token)
	bot.BotToken = string(token)
	bot.Run()
}
