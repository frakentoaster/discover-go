package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getJSON(url string, movie interface{}) error {
	r, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(movie)
}

func main() {
	movie := new(Movie)
	rating := int(movie.ImdbRating)
	getJSON("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json", movie)
	fmt.Printf("The movie : %s was released in %+v - the IMDB rating is %+v%% with %+v votes\n", movie.Title, movie.Year, movie.ImdbRating, movie.ImdbVotes)
	fmt.Printf("Rating is %+v%%", rating)
}
