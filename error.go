package main

import (
	"log"
)

// Check for quick error handling (including logging)
func Check(e error) {
    if e != nil {
        log.Fatal(e)
        panic(e)
    }
}