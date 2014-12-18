package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	fmt.Println(ap)
	headers := make(map[string]string)
	headers["X-ZCloud-AppId"] = ap.ObjectId
	headers["X-ZCloud-MasterKey"] = ap.MasterKey
	req, err := postMultiPart("POST", APIURL+UPLOAD_PATH, path, headers)
	dealWith(err)
	fmt.Println(req.StatusCode)
	body, readErr := ioutil.ReadAll(req.Body)
	dealWith(readErr)
	if req.StatusCode == 200 {
		fmt.Println(body)
	}
}

func (ap app) deploy(v string) {
	fmt.Println("deploy...")
	type jversion struct {
		Version string `json:"version"`
	}
	version := jversion{Version: v}
	b, err := json.Marshal(version)
	dealWith(err)
	resp, resperr := post(APIURL+DEPLOY_PATH, ap, bytes.NewReader(b))
	dealWith(resperr)
	body, readErr := ioutil.ReadAll(resp.Body)
	dealWith(readErr)
	if resp.StatusCode == 200 {
		fmt.Println(body)
	}

}
func (ap app) undeploy() {
	fmt.Println("undeploy...")
	resp, resperr := post(APIURL+UNDEPLOY_PATH, ap, nil)
	dealWith(resperr)
	body, readErr := ioutil.ReadAll(resp.Body)
	dealWith(readErr)
	if resp.StatusCode == 200 {
		fmt.Println(body)
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
func getCurrentApp() (app, error) {
	var ap app
	contents, ioerr := ioutil.ReadFile(getAppPath())
	if ioerr != nil {
		return ap, ioerr
	}
	jsonerr := json.Unmarshal(contents, &ap)
	return ap, jsonerr
}
func listAppVersions(appid string) {
}
func (ap app) log(level string, number, skip int) {
	type alog struct {
		Message, Level, CreateTime string
	}
	type logarray struct {
		Results []alog
	}
	headers := make(map[string]string)
	headers["limit"] = strconv.Itoa(number)
	headers["skip"] = strconv.Itoa(skip)
	resp, err := get(APIURL+LOG_PATH+"/"+level, ap, headers)
	dealWith(err)
	fmt.Println(APIURL + LOG_PATH + "/" + level)
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
