package main

import "os"

func dealWith(err error) {
	if err != nil {
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
func getApp() app {
	ap, err := newApp()
	if err != nil {
		os.Exit(0)
	}
	return ap
}