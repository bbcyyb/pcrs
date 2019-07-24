package common

import (
	"path"
	"runtime"
)

func BuildRunningPath(filename string) string {
	_, callerPath, _, _ := runtime.Caller(1)
	dir := path.Dir(callerPath)
	configPath := path.Join(dir, filename)
	return configPath
}
