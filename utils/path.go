package utils

import "os"

// IsDirExists 判断目录是否存在
func IsDirExists(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}
