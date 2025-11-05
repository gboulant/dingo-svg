package main

import "log"

func main() {
	if err := demo01(); err != nil {
		log.Fatal(err)
	}
}
