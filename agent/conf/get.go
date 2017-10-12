package conf

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetIP() string {
	// 通过访问百度来获取本机IP
	conn, err := net.Dial("udp", "111.13.101.208:80")
	defer conn.Close()
	CheckErr("Failed to get native ip, the reason is:", err)
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	return ip
}

func GetMD5() string {
	buf := make([]byte, 8192)
	var str string

	for _, filename := range Filename {
		if exist := CheckFileExists(filename); exist == true {
			// 文件存在，继续进行以下操作
			file, err := os.Open(filename)
			CheckErr("Open file is failed, the reason is:", err)

			// 读取数据
			for {
				len, _ := file.Read(buf)
				if len == 0 {
					break
				}
				str += string(buf[:len])
			}
		}

	}

	m := md5.New()
	io.WriteString(m, str)
	return hex.EncodeToString(m.Sum(nil))
}

func GetJson(str string) (sina_ptr *SinaJson, err error) {
	var sina_url_data []byte
	if response, err := http.Get(str); err != nil {
		return sina_ptr, err
	} else {
		if sina_url_data, err = ioutil.ReadAll(response.Body); err != nil {
			return sina_ptr, err
		}
	}

	sina_ptr = &SinaJson{}
	err = json.Unmarshal(sina_url_data, sina_ptr)
	return sina_ptr, err
}

func GetConf(url string) (str string) {
	buf := make([]byte, 4096)
	response, err := http.Get(url)
	CheckErr("Get conf url data is failed. the reason is:", err)
	for {
		len, _ := response.Body.Read(buf)
		if len == 0 {
			break
		}
		str += string(buf[:len])
	}
	return
}

// 获取rsyslogd的版本
func GetRsyslogdVersion() string {
	cmd := exec.Command("rsyslogd", "-v")
	var stdout bytes.Buffer

	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		DebugLog.Println("run rsyslogd -v is failed..the reason is:", err)
	}

	return GetVersion(stdout.String())
}

func GetVersion(str string) string {
	reg, err := regexp.Compile("[0-9]+.[0-9]+.[0-9]+")
	if err != nil {
		DebugLog.Println("regex is failed.  the reason is:", err)
	}
	return reg.FindString(str)
}

func GetNowTime() string {
	cur := time.Now()
	timestamp := int64(cur.UnixNano() / 1000000000)
	return strconv.FormatInt(timestamp, 10)
}
