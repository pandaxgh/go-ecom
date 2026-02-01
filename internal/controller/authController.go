package controller

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Endpoint")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Signup Endpoint")
}

func Me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Me Endpoint")
}
