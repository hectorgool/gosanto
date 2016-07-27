package routers

import (
	"github.com/gorilla/mux"
	"github.com/hectorgool/gosanto/controllers"
)

func SetSiteRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.Root).Methods("GET")
	return router
}
