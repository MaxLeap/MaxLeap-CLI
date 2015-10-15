package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type app struct {
	Name      string
	ObjectId  string
	MasterKey string
}
var currentApp * app

func handle(handler func(a *app)) {
	ap:=getCurrentApp()
	if ap == nil {
		fmt.Println("please choose app first ,by 'use <appname>'")
	}else {
		handler(ap)
	}

}
func (ap app) upload(path string) int {
	//body := createFileForm(path)
	checkStrArg(path)
	headers := make(map[string]string)
	headers["X-ZCloud-AppId"] = ap.ObjectId
	headers["X-ZCloud-MasterKey"] = ap.MasterKey
	return formatResult(postMultiPart("POST", APIURL+UPLOAD_PATH, path, headers))
}

type jversion struct {
	Version string `json:"version"`
}

func (ap app) deploy(v string) int {
	version := jversion{Version: v}
	b, err := json.Marshal(version)
	dealWith(err)
	return formatResult(post(APIURL+DEPLOY_PATH, ap, bytes.NewReader(b)))

}
func formatResult(resp *http.Response, resperr error) int {
	dealWith(resperr)
	body, readErr := ioutil.ReadAll(resp.Body)
	dealWith(readErr)
	fmt.Println()
	fmt.Println(string(body))
	return resp.StatusCode
}
func (ap app) undeploy(v string) int {
	version := jversion{Version: v}
	b, err := json.Marshal(version)
	dealWith(err)
	return formatResult(post(APIURL+UNDEPLOY_PATH, ap, bytes.NewReader(b)))
}
func use(name string) {
	checkStrArg(name)
	apps := listApps()
	notExist := true
	for i := range apps {
		if apps[i].Name == name {
			currentApp=&apps[i]
			notExist = false
		}
	}
	if notExist {
		fmt.Println("app " + name + " is not exist")
	}

}
func getCurrentApp() (*app) {
	return currentApp
}
func (ap app) listAppVersions() string {
	resp, err := get(APIURL+LIST_VERSION, ap, nil)
	dealWith(err)
	var results string=""
	if resp.StatusCode==200 {
		bytesRes, readerr := ioutil.ReadAll(resp.Body)
		results=string(bytesRes)
		dealWith(readerr)
	}
	return string(results)
}
func (ap app) log(level string, number, skip int) {
	type alog struct {
		Message, Level, CreateTime string
	}
	type logarray struct {
		Results []alog
	}
	headers := make(map[string]string)
	limit := strconv.Itoa(number)
	skiped := strconv.Itoa(skip)
	url := APIURL + LOG_PATH + "/" + level + "?limit=" + limit + "&skip=" + skiped
	resp, err := get(url, ap, headers)
	dealWith(err)
	contents, ioerr := ioutil.ReadAll(resp.Body)
	dealWith(ioerr)
	var logs logarray
	jsonErr := json.Unmarshal(contents, &logs)
	dealWith(jsonErr)
	if resp.StatusCode == 200 {
		result := logs.Results
		upbound := len(result) - 1
		for j := range result {
			i := upbound - j
			fmt.Println(result[i].CreateTime + " " + result[i].Level + " " + strings.Replace(result[i].Message, "\n", "", -1))
		}
	}
}
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
