package main

import (
	"log"
	"time"
	"runner/runner"
	"os"
)

const timeout = 3 * time.Second

func main() {
	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("timeout")
			os.Exit(1)
			
		case runner.ErrInterrupt:
			log.Println("ErrInterrupt")
			os.Exit(2)
		}
	}

	log.Println("end.")
}

func createTask() func(int) {
	return func(id int) {
		log.Println("Processor - Task #%d.", id)

		time.Sleep(time.Duration(id) * time.Second)
	}
}