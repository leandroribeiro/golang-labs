package controllers

import (
	"github.com/leandroribeiro/go-labs/api-lab3/api/responses"
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	//TODO
	responses.JSON(w, http.StatusCreated, nil)
}

func (server *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	//TODO
	responses.JSON(w, http.StatusCreated, )
}