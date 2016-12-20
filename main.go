package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job, goGroup *sync.WaitGroup) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	goGroup.Done()
}

func main() {
	goGroup := new(sync.WaitGroup)
	fmt.Println("Starting")

	hello := new(Job)
	hello.i = 0
	hello.max = 2
	hello.text = "Hello"

	world := new(Job)
	world.i = 0
	world.max = 2
	world.text = "World"

	go outputText(hello, goGroup)
	go outputText(world, goGroup)

	goGroup.Add(2)
	goGroup.Wait()
}
