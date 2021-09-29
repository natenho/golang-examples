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

	for {
		foo(x, y)
		go bar(x, y)
		go bar_fixed(x, y)
		time.Sleep(1 * time.Second)
	}
}

func foo(x int, y int) {
	defer func() {
		log.Printf("(foo) Eu tambem sou executado (ultimo)")
	}()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("(foo) Ocorreu panic mas tudo vai continuar funcionando!! %v \n", r)
		}
	}()
	defer func() {
		log.Printf("(foo) Eu tambem sou executado (segundo)")
	}()
	defer func() {
		log.Printf("(foo) Eu tambem sou executado (primeiro)")
	}()
	divisionByZero := x / y
	log.Printf("%v", divisionByZero)
}

func bar(x int, y int) {
	log.Printf("(bar) go routines nao sobem panic para o main, devem ter seu proprio defer!!")
	divisionByZero := x / y
	log.Printf("%v", divisionByZero)
}

func bar_fixed(x int, y int) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("(bar_fixed) Ocorreu panic mas tudo vai continuar funcionando!! %v \n", r)
		}
	}()
	divisionByZero := x / y
	log.Printf("%v", divisionByZero)
}
