package main

const APIVERSION = ""

const APIURL string = "http://10.10.10.176:8080"

//const APIURL string = "https://api.appcube.io"
const LOGIN_PATH = APIVERSION + "/orgUsers/login2"
const UPLOAD_PATH = APIVERSION + "/cloudcode" + "/upload"
const LIST_APPS_PATH = APIVERSION + "/apps"
const LOG_PATH = APIVERSION + "/console/logs"
const DEPLOY_PATH = APIVERSION + "/cloudcode" + "/deploy"
const UNDEPLOY_PATH = APIVERSION + "/cloudcode" + "/unDeploy"
const LIST_VERSION = APIVERSION + "/cloudcode" + "/versions"
