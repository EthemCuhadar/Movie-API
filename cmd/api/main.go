package main

import (
	"fmt"
	"log"

	database "github.com/EthemCuhadar/Movie-API/pkg/migrations"
	"github.com/joho/godotenv"
)

func main() {
	// MovieList, _ := utils.ReadJSON("C://Users/iethe/go/src/Movie-API/movies.json")
	// fmt.Println(reflect.TypeOf(MovieList[99]))
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPsqlDB()
	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		fmt.Println("db is nil")
	}
}
