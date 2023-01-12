package main

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/routes"
	"github.com/bootcamp-go/desafio-go-web/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	str, err := store.NewStore().LoadTicketsFromFile("../tickets.csv")

	if err != nil {
		panic("Imposible abrir archivo")
	}

	router := gin.Default()
	routerInit := routes.NewRouter(router, str)
	routerInit.SetRoutes()

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
