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
type Session struct {
	ids Ids
}
func (session Session)login(username, passwd string)(bool) {
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
	unmarshalErr := json.Unmarshal(contents, &response)
	dealWith(unmarshalErr)
	session.ids=response
	return true
}
func (session Session)getSession() Ids {
	return session
}
