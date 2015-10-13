package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"fmt"
)

type userinfo struct {
	Loginid  string `json:"loginid"`
	Password string `json:"password"`
}
type Ids struct {
	SessionToken string
	OrgId        string
	Username     string
	UpdatedAt    time.Time
	CreatedAt    time.Time
	UserType     int
	ObjectId     string
}

func login(username, passwd string) bool {
	clear()
	os.Mkdir(getDir(), 0700)
	os.Chmod(getDir(), 0700)
	if "CN" == region{
		host=CN
	}else if "US"==region{
		host=US
	}else {
		host=region
	}
	result:=login2(username,passwd,host)
	persistHostString()
	return result
}
func login2(username,passwd,url string) bool{
	APIURL="https://"+url
	client := &http.Client{}
	data := userinfo{Loginid: username, Password: passwd}
	bdata, marshalErr := json.Marshal(data)
	dealWith(marshalErr)
	req, reqErr := http.NewRequest("POST", APIURL+LOGIN_PATH, bytes.NewReader(bdata))
	dealWith(reqErr)
	req.Header.Add("Content-Type", "application/json")
	dealWith(reqErr)
	resp, respErr := client.Do(req)
	dealWith(respErr)
	if resp.StatusCode != 200 {
		return false
	}
	var response Ids
	contents, _ := ioutil.ReadAll(resp.Body)
	fileErr := ioutil.WriteFile(getSessionPath(), contents, 0644)
	dealWith(fileErr)
	unmarshalErr := json.Unmarshal(contents, &response)
	dealWith(unmarshalErr)
	return true
}
func getSession() Ids {
	data, ioerr := ioutil.ReadFile(getSessionPath())
	dealWith(ioerr)
	var ids Ids
	unmarshalErr := json.Unmarshal(data, &ids)
	dealWith(unmarshalErr)
	return ids
}

func persistHostString(){
	err := ioutil.WriteFile(getHostPath(),[]byte(host), 0700)
	dealWith(err)
}
func initHostString() {
	if data,ioerr:=ioutil.ReadFile(getHostPath());ioerr==nil{
		host=string(data)
		APIURL="https://"+host
	}else {
		fmt.Println("pls login first")
	}


}
