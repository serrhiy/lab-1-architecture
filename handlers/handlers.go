package handlers

import (
	"encoding/json"
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
	answer, err := json.Marshal(TimeResponse{currentTime})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(answer)
}

var Handlers = map[string]Handler{
	"/time": {[]string{http.MethodGet}, timeHandler},
}
