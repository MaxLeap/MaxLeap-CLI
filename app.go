package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type app struct {
	Name      string
	ObjectId  string
	MasterKey string
}

func newApp() (app, error) {
	ap, err := getCurrentApp()
	if err != nil {
		fmt.Println("please choose app first ,by 'use <appname>'")
	}
	return ap, err
}
func (ap app) upload(path string) {
	//body := createFileForm(path)
	checkStrArg(path)
	headers := make(map[string]string)
	headers["X-ZCloud-AppId"] = ap.ObjectId
	headers["X-ZCloud-MasterKey"] = ap.MasterKey
	formatResult(postMultiPart("POST", APIURL+UPLOAD_PATH, path, headers))

}
func checkStrArg(arg string) {
	if arg == "" {
		fmt.Println("miss argument,find help with --help")
		os.Exit(0)
	}
}
func (ap app) deploy(v string) {
	fmt.Println("deploy...")
	checkStrArg(v)
	type jversion struct {
		Version string `json:"version"`
	}
	version := jversion{Version: v}
	b, err := json.Marshal(version)
	dealWith(err)
	formatResult(post(APIURL+DEPLOY_PATH, ap, bytes.NewReader(b)))

}
func formatResult(resp *http.Response, resperr error) {
	dealWith(resperr)
	body, readErr := ioutil.ReadAll(resp.Body)
	dealWith(readErr)
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}
func (ap app) undeploy() {
	fmt.Println("undeploy...")
	formatResult(post(APIURL+UNDEPLOY_PATH, ap, nil))
}
func use(name string) {
	checkStrArg(name)
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
func getCurrentApp() (app, error) {
	var ap app
	contents, ioerr := ioutil.ReadFile(getAppPath())
	if ioerr != nil {
		return ap, ioerr
	}
	jsonerr := json.Unmarshal(contents, &ap)
	return ap, jsonerr
}
func (ap app) listAppVersions() {
	resp, err := get(APIURL+LIST_VERSION, ap, nil)
	dealWith(err)
	results, readerr := ioutil.ReadAll(resp.Body)
	dealWith(readerr)
	fmt.Println(resp.Status)
	fmt.Println(string(results))
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
	fmt.Println(url)
	fmt.Println(resp.StatusCode)
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
