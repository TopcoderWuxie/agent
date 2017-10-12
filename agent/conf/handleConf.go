package conf

import (
	"time"
)

func HandleConf() {
	period := time.Duration(60) * time.Second
	ticker := time.NewTicker(period)
	for {
		select {
		case <-ticker.C:
			go Handle()
		}
	}
}

func Handle() {
	ip := GetIP()
	md5 := GetMD5()
	url := BaseUrl + "get_config/?ip=" + ip + "&MD5=" + md5
	structData, err := GetJson(url)
	if err != nil {
		DebugLog.Println("get json is error!", err)
	} else {
		// 配置文件发生变化，更新配置文件
		if structData.IsChanged == true {
			// 更新/etc/rsyslog.conf文件
			ChangeFile(GetConf(BaseUrl+structData.RsyslogUrl), Filename[0])
			// 更新/etc/rsyslog.d/log.conf文件
			ChangeFile(GetConf(BaseUrl+structData.LogUrl), Filename[1])

			RestartRsyslog(structData.Timestamp, ip)
		} else {
			DebugLog.SetPrefix("[INFO]")
			DebugLog.Println("the config is not change")
			DebugLog.SetPrefix("[Error]")
		}
	}
}
