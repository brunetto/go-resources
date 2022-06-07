package main

import (
	"fmt"
	"log"

	"github.com/brunetto/go-resources/testing/div"
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
