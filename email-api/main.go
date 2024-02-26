package main

import (
	"email-challenge/controllers"
	"email-challenge/utils"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	utils.LoadEnvVars()
	port := os.Getenv("PORT")

	app := controllers.App{}

	r := app.Routes()

	if port == "" {
		log.Println("Env variables PORT are not set.")
	}

	log.Printf("App is running on: http://localhost:%s\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatalf("error in server: %v", err)
	}
}
