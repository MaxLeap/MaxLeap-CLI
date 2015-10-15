package main

import (
	"fmt"
	"os"
	"time"
)

func dealWith(err error) {
	if err != nil {
		fmt.Println()
		fmt.Println("error:  "+err.Error())
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
		fmt.Println("miss arguments,find help with --help")
		os.Exit(0)
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
