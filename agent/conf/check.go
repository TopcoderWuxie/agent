package conf

import (
	"os"
)

func CheckErr(str string, err error) {
	if err != nil {
		DebugLog.Println(str, err)
	}
}

// true 文件存在，false 文件不存在
func CheckFileExists(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
