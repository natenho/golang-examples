package main

import (
	"log"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Ocorreu panic que chegou até o main %v", r)
		}
	}()
	x := 7
	y := 0

	for  {
		foo(x, y)
		time.Sleep(1 * time.Second)
	}
}

func foo(x int, y int) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Ocorreu panic mas tudo vai continuar funcionando!! %v \n", r)
		}
	}()
	divisionByZero := x / y
	log.Printf("%v", divisionByZero)
}
