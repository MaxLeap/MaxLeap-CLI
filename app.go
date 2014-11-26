package main

import (
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
type alog struct {
	Message, Level, CreateAt string
}
type logarray struct {
	Results []alog
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
	body, readErr := ioutil.ReadAll(req.Body)
	dealWith(readErr)
	if req.StatusCode == 200 {
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
func log(level string, number, skip int) {
	ap := getCurrentApp()
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
	fmt.Println(string(contents))
	if resp.StatusCode == 200 {
		result := logs.Results
		for i := range result {
			fmt.Println(result[i].CreateAt + " " + result[i].Level + " " + strings.Replace(result[i].Message, "\n", "", -1))
		}
	}
}
