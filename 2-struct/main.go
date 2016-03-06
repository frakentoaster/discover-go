package main

import "fmt"
import "time"

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func main() {
	form := "January 2, 2006"
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	fmt.Printf("Hello %s\n", u.Name)
	t, _ := time.Parse(form, u.DOB)
	r := time.Now()
	result := r.Year() - t.Year()
	fmt.Printf("%s was born in %s would be %d years old today.\n", u.Name, u.City, result)
}
