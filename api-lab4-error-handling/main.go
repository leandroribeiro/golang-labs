package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type loginSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginUser(username string, password string) (bool, error) {

	// Just for testing purpose
	if username == "dartvader" && password == "darkside" {
		return true, nil
	}

	if username == "dartvader" && password != "" {
		return false, nil
	}

	return false, errors.New("Jedis are trying invade!")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}
	// Read request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}

	// Parse body as json.
	var schema loginSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		return
	}

	ok, err := loginUser(schema.Username, schema.Password)
	if err != nil {
		log.Printf("Login user DB error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		return
	}

	if !ok {
		log.Printf("Unauthorized access for user: %v", schema.Username)
		w.WriteHeader(401) // Wrong password or username, Return 401.
		return
	}
	w.WriteHeader(200) // Successfully logged in.
}

func main() {
	http.HandleFunc("/login/", loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
