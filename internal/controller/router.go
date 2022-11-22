package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"serv/internal/usecase/entity"
	"serv/internal/usecase/repo"
	"serv/pkg/httpserver"
	"serv/pkg/logger"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct {
	routs []httpserver.Route
}

func NewController() *Controller {
	logger.Debug("Controller initialized")
	return &Controller{}
}

func (lc *Controller) GetRoutes() []httpserver.Route {
	routes := make([]httpserver.Route, 0)

	routes = append(routes, httpserver.Route{Name: "Get all cities", Method: http.MethodGet, Pattern: "/getlist/{lim}",
		HandlerFunc: lc.getlist})

	routes = append(routes, httpserver.Route{Name: "Get city", Method: http.MethodGet, Pattern: "/getcity/{id}",
		HandlerFunc: lc.getCity})

	routes = append(routes, httpserver.Route{Name: "Create new city", Method: http.MethodPost, Pattern: "/createCity",
		HandlerFunc: lc.createCity})

	routes = append(routes, httpserver.Route{Name: "Create new carbase", Method: http.MethodPost, Pattern: "/createCarbase",
		HandlerFunc: lc.createCarbase})

	routes = append(routes, httpserver.Route{Name: "Edit collection", Method: http.MethodPatch, Pattern: "/edit",
		HandlerFunc: lc.editСollection})

	routes = append(routes, httpserver.Route{Name: "Delete collection", Method: http.MethodDelete, Pattern: "/delete/{id}",
		HandlerFunc: lc.deleteСollection})
	return routes
}

func (lc *Controller) getlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lim, err := strconv.ParseInt(vars["lim"], 10, 0)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Invalid lim"))
	}
	its, err := repo.GetList(lim)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusInternalServerError, err)
	}
	httpserver.WriteResponse(w, http.StatusOK, its)
}

func (lc *Controller) getCity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 0)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Invalid id"))
	}

	city, err := repo.GetCity(id)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusInternalServerError, err)
	}

	httpserver.WriteResponse(w, http.StatusOK, city)
}

func (lc *Controller) createCity(w http.ResponseWriter, r *http.Request) {
	var mod entity.City
	err := json.NewDecoder(r.Body).Decode(&mod)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Err of body in Post method, invalid model"))
		return
	}

	err = repo.CreateCity(&mod)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Err of Create Reindexer"))
		return
	}

	httpserver.WriteResponse(w, http.StatusOK, fmt.Sprintf("New city was created"))
}

func (lc *Controller) createCarbase(w http.ResponseWriter, r *http.Request) {
	var mod entity.Carbase
	err := json.NewDecoder(r.Body).Decode(&mod)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Err of body in Post method, invalid model"))
		return
	}

	err = repo.CreateCarbase(&mod)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Err of Create Reindexer"))
		return
	}

	httpserver.WriteResponse(w, http.StatusOK, "New carbase was created")
}

func (lc *Controller) editСollection(w http.ResponseWriter, r *http.Request) {

}

func (lc *Controller) deleteСollection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Invalid id"))
	}

	err = repo.Delete(id)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusInternalServerError, err)
	}

	w.WriteHeader(http.StatusOK)
}
