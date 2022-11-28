package controller

import (
	"encoding/json"
	"net/http"
	"serv/internal/entity"
	"serv/internal/repo"
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

	routes = append(routes, httpserver.Route{Name: "Get all cities", Method: http.MethodGet, Pattern: "/city",
		HandlerFunc: lc.getAll})

	routes = append(routes, httpserver.Route{Name: "Get one city", Method: http.MethodGet, Pattern: "/city/{id}",
		HandlerFunc: lc.getOne})

	routes = append(routes, httpserver.Route{Name: "Create new city", Method: http.MethodPost, Pattern: "/city",
		HandlerFunc: lc.create})

	routes = append(routes, httpserver.Route{Name: "Edit city", Method: http.MethodPut, Pattern: "/city",
		HandlerFunc: lc.edit})

	routes = append(routes, httpserver.Route{Name: "Delete city", Method: http.MethodDelete, Pattern: "/city",
		HandlerFunc: lc.delete})
	return routes
}

func (lc *Controller) getAll(w http.ResponseWriter, r *http.Request) {
	l := r.URL.Query().Get("limit")
	o := r.URL.Query().Get("offset")
	if len(l) == 0 || len(o) == 0 {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Invalid params for get"))
	}

	lim, err := strconv.Atoi(l)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Invalid limit for get"))
	}

	offs, err := strconv.Atoi(o)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Invalid offset for get"))
	}

	cities, err := repo.GetList(lim, offs)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusInternalServerError, err)
	}

	httpserver.WriteResponse(w, http.StatusOK, cities)
}

func (lc *Controller) getOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 0)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Invalid id"))
	}

	city, err := repo.GetOne(id)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusInternalServerError, err)
	}

	httpserver.WriteResponse(w, http.StatusOK, city)
}

func (lc *Controller) create(w http.ResponseWriter, r *http.Request) {
	var city entity.City
	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Err of body in Post method, invalid model"))
		return
	}

	err = repo.Create(&city)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusInternalServerError, httpserver.NewError(httpserver.BadRequestError, "Err of Create Reindexer"))
		return
	}

	httpserver.WriteResponse(w, http.StatusOK, "New city was created")
}

func (lc *Controller) edit(w http.ResponseWriter, r *http.Request) {
	var city entity.City
	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Err of body in Post method, invalid model"))
		return
	}

	res, err := repo.Edit(city)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusInternalServerError, err)
		return
	}

	mess := "City updated"

	if res == 0 {
		mess = "No such city"
	}

	httpserver.WriteResponse(w, http.StatusOK, mess)
}

func (lc *Controller) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusBadRequest, httpserver.NewError(httpserver.BadRequestError, "Invalid id"))
	}

	num, err := repo.Delete(id)
	if err != nil {
		httpserver.WriteResponse(w, http.StatusInternalServerError, err)
	}

	mess := "Deleted"

	if num == 0 {
		mess = "No such city"
	}

	httpserver.WriteResponse(w, http.StatusOK, mess)
}
