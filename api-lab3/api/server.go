package api

import (
	"fmt"
	"github.com/leandroribeiro/go-labs/api-lab3/api/controllers"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

var server = controllers.Server{}

func Run()  {
	var err error
	err = gotenv.Load()
	if err != nil{
		log.Fatalf("Error getting env, not comming through %v", err)
	}else{
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	//TODO
	//seed.Load(server.DB)

	server.Run(":8080")
}