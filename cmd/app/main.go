package main

import (
	"log"
	"os"
)

func main() {
	log := log.New(os.Stdout, "APP: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(log); err != nil {
		log.Println("errror: ", err)
		os.Exit(1)
	}

}

func run(log *log.Logger) *error {
	log.Println(log)
	return nil
}
