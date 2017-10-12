package conf

import (
	"log"
	"os"
)

type SinaJson struct {
	IsChanged  bool   `json:"isChanged"`
	Timestamp  string `json:"timestamp"`
	RsyslogUrl string `json:"rsyslogUrl"`
	LogUrl     string `json:"LogUrl"`
}

var Filename = []string{"/etc/rsyslog.conf", "/etc/rsyslog.d/log.conf"}

const BaseUrl string = "http://***.***.***.***:8001/"

var logFile = "sina_agent.log"
var DebugLog *log.Logger

func init() {
	logFile, _ := os.Create(logFile)
	DebugLog = log.New(logFile, "[ERROR]", log.LstdFlags)
}
