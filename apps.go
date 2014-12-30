package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func listApps() []app {
	resp, err := sessionRequest("GET", APIURL+LIST_APPS_PATH, "application/json", nil)
	dealWith(err)
	contents, ioerr := ioutil.ReadAll(resp.Body)
	dealWith(ioerr)
	apps := make([]app, 0)
	json.Unmarshal(contents, &apps)
	return apps
}
func showApps() {
	apps := listApps()
	if len(apps) <= 0 {
		println("no apps")
		return
	}
	fmt.Println()
	fmt.Print("appid")
	for i := 0; i < len(apps[0].ObjectId)-5; i++ {
		fmt.Print(" ")
	}
	fmt.Println(" appname")
	fmt.Println()
	for i := range apps {
		fmt.Println(apps[i].ObjectId + ":" + apps[i].Name)
	}
}
