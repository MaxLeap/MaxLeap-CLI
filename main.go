package main

import (
	"fmt"
)

func main() {
	fmt.Print("user name:");
	userName:=""
	fmt.Scanln(userName)
	session:=Session{}
	for i := 0; i < 3; i++ {
		passwd, err := GetPass("enter password:")
		if err != nil {
			fmt.Println("can't get password")
			return
		}
		if session.login(userName, passwd) {
			start(session)
			break
		} else {
			if i < 2 {
				fmt.Println("Permission denied, please try again.")
			} else {
				fmt.Println("Permission denied")
			}
		}
	}
}
