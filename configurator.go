package main

import "log"

var (
	filename = "samples/System.ini"
)

func main() {
	err := gui()
	if err != nil {
		log.Panicf("caught error: %v", err)
	}
}
