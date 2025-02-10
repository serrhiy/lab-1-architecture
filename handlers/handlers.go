package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type Handler struct {
	Methods    []string
	Controller func(http.ResponseWriter, *http.Request)
}

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(response http.ResponseWriter, request *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	fmt.Println(currentTime)
}
