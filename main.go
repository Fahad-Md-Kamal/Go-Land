package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}


func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2097016465109-3807901051427-HoaIRYB1SLJGv9zx4gpSS6sz")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03PRTVU8FL-3807950379874-b5920314fe8a4103137680c2944f2083ba7127cdef64f41944ec3c30d490e21c")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is", &slacker.CommandDefinition{
		Description: "youb calculator",
		Example: "my yob is 2022",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("err")
			}
			age := 2022-yob
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