package main

import (
	"fmt"
	"go-resources/testing/div"
	"log"
)

func main() {
	res, err := div.Div(6, 3)
	dieIf(err)

	fmt.Println(res)
}

func dieIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
