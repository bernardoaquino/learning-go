package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Livro struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

var Livros []Livro = []Livro{
	{
		Id:     1,
		Titulo: "O livro",
		Autor:  "Autor do livro",
	},
	{
		Id:     2,
		Titulo: "O livro 2",
		Autor:  "Autor do livro 2",
	},
	{
		Id:     3,
		Titulo: "O livro 3",
		Autor:  "Autor do livro 3",
	},
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func listarLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

func cadastrarLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!")
	}

	var novoLivro Livro
	json.Unmarshal(body, &novoLivro)
	novoLivro.Id = len(Livros) + 1
	Livros = append(Livros, novoLivro)

	encoder := json.NewEncoder(w)
	encoder.Encode(novoLivro)
}

func rotearLivros(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listarLivros(w, r)
	} else if r.Method == "POST" {
		cadastrarLivros(w, r)
	}
}

func buscarLivro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	partes := strings.Split(r.URL.Path, "/")
	if len(partes) > 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, _ := strconv.Atoi(partes[2])
	for _, livro := range Livros {
		if livro.Id == id {
			json.NewEncoder(w).Encode(livro)
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/livros", rotearLivros)
	http.HandleFunc("/livros/", buscarLivro)
}

func configurarServidor() {
	configurarRotas()
	fmt.Println("Servidor rodando na porta 1337")
	log.Fatal(http.ListenAndServe(":1337", nil)) //DefaultServerMux
}

func main() {
	configurarServidor()
}
