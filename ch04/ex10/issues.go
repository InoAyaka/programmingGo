package main

import (
	"Book_gopl/exercise/ch04/ex10/github"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {

	//現日時を取得
	today := time.Now()
	month := today.AddDate(0, -1, 0)
	year := today.AddDate(-1, 0, 0)

	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	//日付順で出力するため、ソート
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
	})

	fmt.Printf("%d issues:\n", result.TotalCount)

	withinM, withinY, more := true, true, true

	for _, item := range result.Items {

		switch {
		case item.CreatedAt.Before(year) && more:
			fmt.Println("more")
			more = false
		case item.CreatedAt.After(year) && item.CreatedAt.Before(month) && withinY:
			fmt.Println("Within 1 year")
			withinY = false
		case item.CreatedAt.After(month) && withinM:
			fmt.Println("Within 1 month")
			withinM = false
		}

		fmt.Printf("  #%-5d %9.9s %v %.55s\n", item.Number, item.User.Login, item.CreatedAt.Format("2006/01/02"), item.Title)
	}

}
