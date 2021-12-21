package main

import (
	"log"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func main() {
	ch := make(chan struct{})
	go func() {
		log.Println("sleep 5s")
		time.Sleep(5 * time.Second)
		close(ch)
	}()
	wait.Until(func() {
		time.Sleep(100 * time.Millisecond)
		sync()
	}, 1000*time.Millisecond, ch)
	log.Println("main exit")
}

func sync() {
	log.Println("test")
}
