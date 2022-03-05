package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dickmanben/story-bot/types"
	"github.com/dickmanben/story-bot/utils"
)

func NewEvent(eventChan chan types.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newEvent := types.Event{}
		if err := json.NewDecoder(r.Body).Decode(&newEvent); err != nil {
			utils.ErrorResponder(w, 400, err)
			return
		}

		eventChan <- newEvent
		utils.JSONSuccessResponder(w)
	}
}
