package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t0 := time.Now()
	defer func() { log.Println("done in ", time.Since(t0)) }()

	createdCh := make(chan string, 1)
	deletedCh := make(chan string, 1)

	// no need to set waitgroups here
	go func() { createdCh <- GetString("created", 3*time.Second) }()
	go func() { deletedCh <- GetString("deleted", 8*time.Second) }()

	created := <-createdCh
	deleted := <-deletedCh

	fmt.Println(created, deleted)
}

func GetString(str string, wait time.Duration) string {
	time.Sleep(wait)

	return str
}
