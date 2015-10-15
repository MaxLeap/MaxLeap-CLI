package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
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
var session Ids
var on bool=true
func login(username, passwd string) bool {
	if "CN" == region{
		host=CN
	}else if "US"==region{
		host=US
	}else {
		host=region
	}
	result:=login2(username,passwd,host)
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
	contents, _ := ioutil.ReadAll(resp.Body)
	unmarshalErr := json.Unmarshal(contents, &session)
	dealWith(unmarshalErr)
	return true
}
func getSession() Ids {
	return session
}


