package main

import (
	"log"
	"os"
	"planners/bot"
	"planners/supabase"
)

func main() {
	// Удалено использование godotenv

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	botToken := os.Getenv("BOT_TOKEN")

	supabaseClient := supabase.NewClient(supabaseUrl, supabaseKey)
	myBot, err := bot.NewBot(botToken, supabaseClient)
	if err != nil {
		log.Fatalf("failed to create bot: %v", err)
	}

	myBot.Start()
}
