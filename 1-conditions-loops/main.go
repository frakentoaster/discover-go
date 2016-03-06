package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100)
	var beautifulWeather bool
	beautifulWeather = true
	school := "Holberton School"
	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}

	if randomNumber > 50 {
		fmt.Println("my random number is ", randomNumber, "and is greater than 50")
	} else {
		fmt.Println("my random number is ", randomNumber, "and is not greater than 50")
	}

	if school == "Holberton School" {
		fmt.Println("I am a student of the ", school)
	} else {
		fmt.Println("I am not a student of the ", school)
	}

	if beautifulWeather {
		fmt.Println("It's beautiful weather!")
	} else {
		fmt.Println("The weather is terrible!")
	}

	for i := 0; i < 3; i++ {
		fmt.Println(holbertonFounders[i], "is a founder")
	}
}
