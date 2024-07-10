package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar página de Usuários"))
	})

	http.HandleFunc("/contatos", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar contatos"))
	})

	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Exit"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}

/*
-> HTTP é um protocolo de comunicação - base da comunicação web
-> Cliente(faz a requisição) - Servidor(Processa a requisição e envia a resposta)
-> Request - Response
-> Rotas
-> URI - Identificador do Recurso
-> Método - GET, POST, PUT, DELETE
*/
