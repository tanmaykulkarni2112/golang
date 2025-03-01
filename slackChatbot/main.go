package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}

func main() {
	//os.Setenv("SLACK_BOT_TOKEN", "xoxb-8533643393714-8535578560724-uNZxlJK59SuuvToRPYSdGZmu")
	//os.Setenv("SLACK_APP_TOKEN", "xapp-1-A08FP6EN9FD-8555378250256-5a8eca92681fe1c60c34e41b4178c907dada8872960c818e2e61f9c7c418027a")

	EnvErr := godotenv.Load()
	if EnvErr != nil {
		log.Fatal("Error loading .env file")
	}
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "YOB calculator",
		Examples: []string{
			"My yob is 2003",
			"My yob is 2008",
			"My yob is 2012",
		},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				print("Error")
			}
			age := 2024 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
