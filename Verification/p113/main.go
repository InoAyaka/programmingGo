package main

import (
	"fmt"
)

type employee struct {
	ID   int
	name string
}

func main() {

	var dilbert employee

	fmt.Println(employeeByID(dilbert.ID).name)

	id := dilbert.ID
	employeeByID(id).name = "huga"

}

func employeeByID(id int) employee {

	emp := employee{
		ID:   1,
		name: "hoge",
	}

	return emp
}
