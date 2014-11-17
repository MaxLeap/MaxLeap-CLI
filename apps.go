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
	for i := range apps {
		fmt.Println(apps[i].ObjectId + ":" + apps[i].Name)
	}
}
