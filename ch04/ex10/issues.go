package main

import (
	"Book_gopl/exercise/ch04/ex10/github"
	"fmt"
	"log"
	"os"
)

func main() {

	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %v %.55s\n", item.Number, item.User.Login, item.CreatedAt.Format("2006/01/02"), item.Title)
	}

}
