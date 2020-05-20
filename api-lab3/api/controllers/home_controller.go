package controllers

import (
	"github.com/leandroribeiro/golang-labs/api-lab3/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request)  {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}