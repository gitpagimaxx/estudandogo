package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func main() {
	fmt.Println("HTTP")

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuario)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
