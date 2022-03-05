package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dickmanben/story-bot/discord"
	"github.com/dickmanben/story-bot/handlers"
	"github.com/dickmanben/story-bot/types"
	"github.com/dickmanben/story-bot/utils"
	"github.com/subosito/gotenv"
)

var EventChan = make(chan types.Event)

func init() {
	gotenv.Load()
}

func main() {

	go discord.NewSession(EventChan)

	// EventChan <- types.Event{
	// 	Channel: "267030923386683393",
	// 	Type:    "RandomNumber",
	// }
	r := handlers.NewRouter(EventChan)

	srv := &http.Server{
		Handler: utils.CorsHandler(r),
		Addr:    ":" + os.Getenv("PORT"),
	}
	srv.SetKeepAlivesEnabled(false)

	log.Fatal(srv.ListenAndServe())
}
