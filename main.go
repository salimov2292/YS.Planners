package main

import (
	"log"
	"os"
	"planners/bot"
	"planners/supabase"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	botToken := os.Getenv("BOT_TOKEN")

	supabaseClient := supabase.NewClient(supabaseUrl, supabaseKey)
	myBot, err := bot.NewBot(botToken, supabaseClient)
	if err != nil {
		panic("failed to create bot: " + err.Error())
	}

	myBot.Start()
}
