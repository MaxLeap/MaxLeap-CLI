package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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

func login(username, passwd string) bool {
	clear()
	os.Mkdir(getDir(), 0700)
	os.Chmod(getDir(), 0700)
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
