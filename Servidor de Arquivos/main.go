package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string{
	if user == "user"{
		return "$1$NJb3xbck$S0pO/8BdnoeW1pnn3EAwc1"
	}
	return ""
	
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run <Caminho Diretório> <porta>")
		os.Exit(1)
	}
	httpdir := os.Args[1]
	porta := os.Args[2]
	http.FileServer(http.Dir(httpdir))

	authenticator := auth.NewBasicAuthenticator("server.com", Secret )
	http.HandleFunc("/", authenticator.Wrap(func (w http.ResponseWriter, r *auth.AuthenticatedRequest){
		http.FileServer(http.Dir(httpdir)).ServeHTTP(w, &r.Request)
	}))


	log.Println("Servidor rodando na porta", porta)
	log.Fatal(http.ListenAndServe(":"+porta, nil))
}
