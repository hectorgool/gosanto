package controllers

import (
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}
