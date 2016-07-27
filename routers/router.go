package routers

import (
	//"net/http"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	
	// Routes for the User entity
	router = SetUserRoutes(router)
	
	// Routes for the Task entity
	router = SetTaskRoutes(router)
	
	// Routes for the TaskNote entity
	router = SetNoteRoutes(router)
	
	// Routes for the Site entity
	router = SetSiteRoutes(router)

	return router

}
