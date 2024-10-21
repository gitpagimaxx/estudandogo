package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Criando usuário")
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	r.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	r.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	r.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)

	fmt.Println("Server is running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
