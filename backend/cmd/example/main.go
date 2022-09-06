package main

import "github.eapteka.ru/example/internal/worker/example"

func main() {
	worker := example.New()
	worker.Run()
}
