package httpserver

import (
	"encoding/json"
	"net/http"
	"serv/pkg/logger"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(http.ResponseWriter, *http.Request)
}

type Controller interface {
	GetRoutes() []Route
}

func WriteResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			logger.Err("Could not marshal task to json", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error in marshalling"))
			return
		}
		w.WriteHeader(statusCode)
		w.Write(jsonBody)
	}
}

const contextRoot string = "/service"

func NewHTTPHandler(routesHandler ...Controller) http.Handler {
	muxHandler := mux.NewRouter().StrictSlash(true)
	for _, routeHandler := range routesHandler {
		routes := routeHandler.GetRoutes()
		for _, route := range routes {
			muxHandler.PathPrefix(contextRoot).Methods(route.Method).Path(route.Pattern).
				Name(route.Name).HandlerFunc(route.HandlerFunc)
		}
	}
	return muxHandler
}
