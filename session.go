package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func login(username, passwd string) {
	checkStrArg(username)
	checkStrArg(passwd)
	clear()
	os.Mkdir(getDir(), 0700)
	os.Chmod(getDir(), 0700)
	client := &http.Client{}
	data := userinfo{Loginid: username, Password: passwd}
	bdata, marshalErr := json.Marshal(data)
	dealWith(marshalErr)
	fmt.Println(string(bdata))
	req, reqErr := http.NewRequest("POST", APIURL+LOGIN_PATH, bytes.NewReader(bdata))
	dealWith(reqErr)
	fmt.Println(APIURL + LOGIN_PATH)
	req.Header.Add("Content-Type", "application/json")
	dealWith(reqErr)
	resp, respErr := client.Do(req)
	dealWith(respErr)
	var response Ids
	fmt.Println(resp.StatusCode)
	contents, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("is %s\n", string(contents))
	fileErr := ioutil.WriteFile(getSessionPath(), contents, 0644)
	dealWith(fileErr)
	unmarshalErr := json.Unmarshal(contents, &response)
	dealWith(unmarshalErr)
}
func getSession() Ids {
	data, ioerr := ioutil.ReadFile(getSessionPath())
	dealWith(ioerr)
	var ids Ids
	fmt.Println(string(data))
	unmarshalErr := json.Unmarshal(data, &ids)
	dealWith(unmarshalErr)
	fmt.Println("ids:" + ids.SessionToken)
	return ids
}
