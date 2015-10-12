package main

import (
	"fmt"
	"os"
	"os/user"
)

func getSessionPath() string {
	return getDir() + "/.session"
}
func getDir() string {
	usr, userErr := user.Current()
	dealWith(userErr)
	return usr.HomeDir + "/.zcc"
}
func getAppPath() string {
	return getDir() + "/.app"
}
func getHostPath() string{
	return getDir()+"/.host"
}
func clear() {
	err := os.RemoveAll(getDir())
	if err != nil {
		fmt.Println(err)
	}
}
