package handlers

import (
	"net/http"
)

type Handler struct {
	Methods []string
	Controller func(http.ResponseWriter, *http.Request)
}
