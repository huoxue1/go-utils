// Package hooks
// this file is copy from https://github.com/Lyrics-you/sail-logrus-formatter/blob/main/hooks/caller.go
package hooks

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	file string = ""
	line int    = 0
	name string = ""
)

func FindCaller(skip int) string {

	for i := 0; i < 10; i++ {
		file, line, name = getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	s := strings.Split(name, ".")
	name := s[len(s)-1]
	return fmt.Sprintf("%s:%d @%s()", file, line, name)
}

func getCaller(skip int) (string, int, string) {
	pc, file, line, ok := runtime.Caller(skip)
	fnName := runtime.FuncForPC(pc).Name()
	// fmt.Println(file)
	// fmt.Println(line)
	if !ok {
		return "", 0, ""
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line, fnName
}
