package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type app struct {
	Name      string
	ObjectId  string
	MasterKey string
}

func deploy(path string) {
	//body := createFileForm(path)
	ap := getCurrentApp()
	fmt.Println(ap)
	headers := make(map[string]string)
	headers["X-ZCloud-AppId"] = ap.ObjectId
	headers["X-ZCloud-MasterKey"] = ap.MasterKey
	req, err := postMultiPart("POST", APIURL+DEPLOY_PATH, path, headers)
	dealWith(err)
	fmt.Println(req.StatusCode)
	io.Copy(os.Stderr, req.Body)
	if req.StatusCode == 200 {
		fmt.Println("success")
	}
}
func use(name string) {
	apps := listApps()
	notExist := true
	for i := range apps {
		if apps[i].Name == name {
			contents, marshalErr := json.Marshal(apps[i])
			dealWith(marshalErr)
			path := getAppPath()
			err := ioutil.WriteFile(path, contents, 0700)
			dealWith(err)
			notExist = false
		}
	}
	if notExist {
		fmt.Println("app " + name + " is not exist")
	}

}
func getCurrentApp() app {
	contents, ioerr := ioutil.ReadFile(getAppPath())
	dealWith(ioerr)
	var ap app
	jsonerr := json.Unmarshal(contents, &ap)
	dealWith(jsonerr)
	return ap
}
func listAppVersions(appid string) {
}
func log(appid string) {
}
