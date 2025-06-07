package routers

import (
	//"apigolang/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	SetupRouterProdutos(router)   //<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	router.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(
			http.Dir("./static/"))))

	return router
}

//pegar apenas essa linha e adicionar no router.go que já está no diretorio
//colar a linha abaixo das demais rotas


