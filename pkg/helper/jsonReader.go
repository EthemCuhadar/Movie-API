package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/EthemCuhadar/Movie-API/pkg/models"
)

func ReadJSON(filename string) ([]models.Movie, error) {
	MovieList := models.MovieList{}
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err = json.Unmarshal(byteValue, &MovieList); err != nil {
		return nil, err
	}
	return MovieList, nil
}
