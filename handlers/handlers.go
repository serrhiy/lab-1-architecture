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

func timeHandler(response http.ResponseWriter, request *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	fmt.Println(currentTime)
}
