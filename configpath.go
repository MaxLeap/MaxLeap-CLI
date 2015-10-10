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
func clear() {

}
