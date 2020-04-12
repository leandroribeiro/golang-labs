package controllers

import "github.com/leandroribeiro/go-labs/api-lab3/api/middlewares"

func (server *Server) initializeRoutes()  {

	// Home Route
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")
}