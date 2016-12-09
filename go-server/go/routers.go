package 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/acme/another-pet/base/",
		Index,
	},

	Route{
		"AddCow",
		"POST",
		"/acme/another-pet/base/cows",
		AddCow,
	},

	Route{
		"CatsGet",
		"GET",
		"/acme/another-pet/base/cats",
		CatsGet,
	},

	Route{
		"AddDog",
		"POST",
		"/acme/another-pet/base/dogs",
		AddDog,
	},

}