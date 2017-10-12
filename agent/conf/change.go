package conf

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"
)

func ChangeFile(str string, filename string) {
	if exist := CheckFileExists(filename); exist == true {
		// 先删除文件，然后再创建
		err := os.Remove(filename)
		CheckErr("remove conf file is failed, the reason is:", err)
	}
	file, err := os.Create(filename)
	CheckErr("create conf file is failed, the reason is:", err)
	defer file.Close()
	file.WriteString(str)
}

func RestartRsyslog(timestamp string, myIP string) {
	cmd := exec.Command("service", "rsyslog", "restart")
	var stderr bytes.Buffer
	var url string

	cmd.Stderr = &stderr
	var vis bool = false
	if err := cmd.Run(); err != nil {
		vis = true
	}

	str := stderr.String()

	version := GetRsyslogdVersion()

	// 重启没有出错
	if len(str) == 0 && vis == false {
		effectTime := GetNowTime()
		url = "report_restart_error/?ip=" + myIP + "&restart_success=true&version=" + version + "&timestamp=" + timestamp + "&effectTime=" + effectTime
	} else {
		url = "report_restart_error/?ip=" + myIP + "&restart_success=false&version=" + version + "&timestamp=" + "&effectTime="
	}
	response, err := http.Get(BaseUrl + url)
	CheckErr("restart rsyslog is failed, the reason is:", err)
	defer response.Body.Close()
}
