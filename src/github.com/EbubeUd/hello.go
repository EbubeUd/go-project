package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	User
	age         uint
	isQualified bool
}

type User struct {
	name  string `required max:"100"`
	email string
}

func main() {

	people := map[string]uint8{
		"Hello": 1,
		"World": 2,
	}

	places := make(map[string]uint8)

	places = map[string]uint8{
		"Enugu": 1,
		"Delta": 2,
	}

	delete(places, "Enugu")
	a := len(people)

	_, out := places["Delta"]

	aDoc := Doctor{}

	aDoc.name = "Ebube"
	aDoc.age = 12
	aDoc.email = "kripsonud@gmail.com"
	aDoc.isQualified = true

	t := reflect.TypeOf(User{})
	field, _ := t.FieldByName("name")

	fmt.Printf("%+v\n", people["Hello"])
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", places)
	fmt.Printf("%+v\n", out)
	fmt.Printf("%+v\n", aDoc)
	fmt.Printf("%+v\n", field)

}
