package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Before 'sig := <-sigs' ")
		sig := <-sigs
		fmt.Println("After 'sig := <-sigs' ")
		fmt.Println(sig)
		fmt.Println("Value of sig = ", sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	fmt.Println("The Done value is", <-done)
	//<-done
	fmt.Println("exiting")
}
