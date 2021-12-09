package main

import "log"

func main() {

	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for _, animal := range animals {
		log.Println(animal)
	}

	person := make(map[string]string)
	person["abi"] = "aam"
	person["mama"] = "mela"
	person["anak1"] = "sakha"
	person["anak2"] = "shahim"

	for whois, people := range person {
		log.Println(whois, people)
	}

	var firstLine = "Once upon a midnight dreary"
	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Melati Faridatul", "Ilmi", "melatiilmi@gmail.com", 31})
	users = append(users, User{"M Kholiq", "Amrullah", "hallurma@gmail.com", 36})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
