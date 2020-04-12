package controllers

import (
	"github.com/leandroribeiro/go-labs/api-lab3/api/responses"
	"net/http"
)

func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	//TODO
	responses.JSON(w, http.StatusCreated, nil)
}

func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	//TODO
	responses.JSON(w, http.StatusCreated, nil)
}

func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	//TODO
	responses.JSON(w, http.StatusCreated, nil)
}

func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	//TODO
	responses.JSON(w, http.StatusCreated, nil)
}