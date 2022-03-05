package handlers

import (
	"net/http"

	"github.com/dickmanben/story-bot/types"
	"github.com/dickmanben/story-bot/utils"
	"github.com/gorilla/mux"
)

type route struct {
	Name           string
	Pattern        string
	Method         string
	Handler        http.HandlerFunc
	SpecialHandler func(c chan types.Event) http.HandlerFunc
}

var routes = []route{
	{
		Name:    "healthCheck",
		Pattern: "/",
		Method:  "GET",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			utils.JSONSuccessResponder(w)
		},
	},
}

var extraSpecialRoutes = []route{
	{
		Name:           "New Event",
		Pattern:        "/event",
		Method:         "POST",
		SpecialHandler: NewEvent,
	},
}

//
//NewRouter returns a new mux router
//
func NewRouter(c chan types.Event) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, specialRoute := range extraSpecialRoutes {
		router.
			Methods(specialRoute.Method).
			Path(specialRoute.Pattern).
			Name(specialRoute.Name).
			Handler(specialRoute.SpecialHandler(c))
	}

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
