package repo

import "os"

func dbNotExist(name string) bool {
	_, err := os.Stat(name)
	return os.IsNotExist(err)
}
