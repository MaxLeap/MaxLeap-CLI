package main

const APIVERSION = "/2.0"
var host string=US
var region string="US"
const US string = "api.appcube.io"
const CN string = "api.maxleap.cn"
var APIURL string = "https://"+host
const LOGIN_PATH = APIVERSION + "/orgUsers/login2"
const UPLOAD_PATH = APIVERSION + "/cloudcode" + "/upload"
const LIST_APPS_PATH = APIVERSION + "/apps"
const LOG_PATH = APIVERSION + "/console/logs"
const DEPLOY_PATH = APIVERSION + "/cloudcode" + "/deploy"
const UNDEPLOY_PATH = APIVERSION + "/cloudcode" + "/unDeploy"
const LIST_VERSION = APIVERSION + "/cloudcode" + "/versions"
const SERVER_VERSION = "1.0"
