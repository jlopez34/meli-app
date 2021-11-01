package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jlopez34/meli-app/pkg/finding"
	"github.com/jlopez34/meli-app/pkg/rest"
	"github.com/jlopez34/meli-app/pkg/storage"
)

func main() {
	r, err := storage.SetupStorage()
	if err != nil {
		fmt.Println(err)
	}

	as := finding.NewService(r)
	router := rest.Handler(as)

	fmt.Println("Starting server on port 8080......")
	log.Fatal(http.ListenAndServe(":8080", router))
}
