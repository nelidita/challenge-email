package utils

import (
	"net/http"
	"os"
)

func SetBasicHeaders(req *http.Request) {
	username := os.Getenv("ZS_USER")
	password := os.Getenv("ZS_PASSWORD")

	if username == "" || password == "" {
		panic("ZS_USER and ZS_PASSWORD no exists")
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)
}
