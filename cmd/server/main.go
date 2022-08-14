package main

import (
	"Ictus-Backend/pkg/api"
	"Ictus-Backend/pkg/db"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	pgDb, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	router := api.NewApi(pgDb)
	log.Println("Up and running")
	port := os.Getenv("PORT")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("Error from router: %s\n", err)
	}

	//pgdb, err := db.NewDB()
	//if err != nil {
	//	panic(err)
	//}
	//router := api.NewApi(pgdb)
	//port := os.Getenv("PORT")
	//err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)

}
