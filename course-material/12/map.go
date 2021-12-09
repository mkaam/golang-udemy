package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// maps
	myMap := make(map[string]User)

	myMap2 := make(map[string]interface{})

	myMap2["mela"] = 1

	aam := User{
		FirstName: "Aam",
		LastName:  "Amrullah",
	}

	myMap["aam"] = aam

	log.Println(myMap2)
}
