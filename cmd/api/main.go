package main

import (
	"fmt"
	"log"

	utils "github.com/EthemCuhadar/Movie-API/pkg/helper"
	database "github.com/EthemCuhadar/Movie-API/pkg/migrations"
)

func main() {
	MovieList, _ := utils.ReadJSON("C://Users/iethe/go/src/Movie-API/movies.json")
	// fmt.Println(reflect.TypeOf(MovieList[99]))

	db, err := database.NewPsqlDB("../../.env")
	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		fmt.Println("db is nil")
	}

	movieRepo := database.NewMovieRepository(db)
	movieRepo.Migrations()
	movieRepo.InsertData(MovieList)
}
