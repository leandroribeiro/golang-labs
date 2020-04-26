package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leandroribeiro/go-labs/api-lab3/api/models"
	"github.com/prometheus/common/log"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreateUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	samples := []struct {
		inputJSON    string
		statusCode   int
		nickname     string
		email        string
		errorMessage string
	}{
		{
			inputJSON:    `{}`,
			statusCode:   0,
			nickname:     "",
			email:        "",
			errorMessage: "",
		},
		{
			inputJSON:    `{}`,
			statusCode:   0,
			nickname:     "",
			email:        "",
			errorMessage: "",
		},
		{
			inputJSON:    `{}`,
			statusCode:   0,
			nickname:     "",
			email:        "",
			errorMessage: "",
		},
		{
			inputJSON:    `{}`,
			statusCode:   0,
			nickname:     "",
			email:        "",
			errorMessage: "",
		},
		{
			inputJSON:    `{}`,
			statusCode:   0,
			nickname:     "",
			email:        "",
			errorMessage: "",
		},
		{
			inputJSON:    `{}`,
			statusCode:   0,
			nickname:     "",
			email:        "",
			errorMessage: "",
		},
		{
			inputJSON:    `{}`,
			statusCode:   0,
			nickname:     "",
			email:        "",
			errorMessage: "",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateUser)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 201 {
			assert.Equal(t, responseMap["nickname"], v.nickname)
			assert.Equal(t, responseMap["email"], v.email)
		}
		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestGetUsers(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetUsers)
	handler.ServeHTTP(rr, req)

	var users []models.User
	err = json.Unmarshal([]byte(rr.Body.String()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(users), 2)
}

func TestGetUserByID(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}
	userSample := []struct {
		id           string
		statusCode   int
		nickname     string
		email        string
		errorMessage string
	}{
		{
			id:         strconv.Itoa(int(user.ID)),
			statusCode: 200,
			nickname:   user.Nickname,
			email:      user.Email,
		},
		{
			id:         "unknow",
			statusCode: 400,
		},
	}

	for _, v := range userSample {

		req, err := http.NewRequest("GET", "/users", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": v.id})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.GetUser)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			log.Fatalf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, user.Nickname, responseMap["nickname"])
			assert.Equal(t, user.Email, responseMap["email"])
		}
	}
}

func TestUpdateUser(t *testing.T) {
	//TODO
}

func TestDeleteUser(t *testing.T) {
	//TODO
}
