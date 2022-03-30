package main

import (
	"fmt"
	"reflect"

	utils "github.com/EthemCuhadar/Movie-API/pkg/helper"
)

func main() {
	MovieList, _ := utils.ReadJSON("C://Users/iethe/go/src/Movie-API/movies.json")
	fmt.Println(reflect.TypeOf(MovieList[99]))
}
