package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/leandroribeiro/golang-labs/api-lab4-error-handling/infrastructure"
	"github.com/leandroribeiro/golang-labs/api-lab4-error-handling/interfaces"
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

type rootHandler func(http.ResponseWriter, *http.Request) error

func loginHandler(w http.ResponseWriter, r *http.Request) error {

	if r.Method != http.MethodPost {
		return infrastructure.NewHTTPError(nil, 405, "Method not allowed.")
	}

	// Read request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("Request body read error: %v", err)
	}

	// Parse body as json.
	var schema loginSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		return infrastructure.NewHTTPError(err, 400, "Bad request: invalid JSON")
	}

	ok, err := loginUser(schema.Username, schema.Password)
	if err != nil {
		return fmt.Errorf("loginUser DB error: %v", err)
	}

	if !ok { // Authentication failed.
		return infrastructure.NewHTTPError(nil, 401, "Wrong password or username")
	}

	w.WriteHeader(200) // Successfully authenticated. Return access token?
	return nil
}

func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	err := fn(w, r) // Call handler function
	if err == nil {
		return
	}

	// This is where our error handling logic starts
	log.Printf("An error accured: %v", err)// Log the error

	clientError, ok := err.(interfaces.ClientError) // Check if it is a ClientError.
	if !ok{
		// If the error is not ClientError, assume that it is ServerError.
		w.WriteHeader(500)// return 500 Internal Server Error
		return
	}

	body, err := clientError.ResponseBody()
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		return
	}

	status, headers := clientError.ResponseHeaders()
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)
	w.Write(body)

}

func main() {
	// Notice rootHandler
	http.Handle("/login/", rootHandler(loginHandler))

	fmt.Println("Server is ready and listening :)")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
