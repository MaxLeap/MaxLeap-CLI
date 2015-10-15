package main

import (
	"fmt"
	"os"
	"time"
	"errors"
)

func dealWith(err error) {
	if err != nil {
		fmt.Println()
		panic(err)
	}
}
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func println(content string) {
	fmt.Println(content)
}
func checkStrArg(arg string) {
	if arg == "" {
		panic(errors.New("miss arguments,find help with --help"))
	}

}
func showProgress(ch chan int) {
	for {
		select {
		case status := <-ch:
			if 200 == status {
				fmt.Println("success")
			} else {
				fmt.Println("failed")
			}
			return
		default:
			fmt.Print(".")
		}
		time.Sleep(time.Second)
	}
}
func startWithProgress(fn func() int) {
	chann := make(chan int)
	go func() {
		status := fn()
		chann <- status
	}()
	showProgress(chann)
}
