package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zohaibsoomro/book-server-mongodb/config"
	"github.com/zohaibsoomro/book-server-mongodb/model"
	"github.com/zohaibsoomro/book-server-mongodb/routes"
)

func main() {
	r := httprouter.New()
	routes.RegisterRoutes(r)
	client := config.ConnectDB()
	err := config.PingDB(client)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection successful")
	model.SetClient(client)
	log.Fatal(http.ListenAndServe(":8080", r))
	defer client.Disconnect(context.Background())

}
