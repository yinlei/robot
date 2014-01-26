package logger

import (
	"fmt"
	"strings"
	"runtime"
	"path"
)

type callInfo struct {
	packageName string
	fileName	string
	funcName	string
	line		int
}

func (ci *callInfo) shortFormat() string {
	return fmt.Sprintf("[%s]", ci.packageName)
}


func (ci *callInfo) verboseFormat() string {
	return fmt.Sprintf("[%s] (%s:%s:%d)", ci.packageName, ci.fileName, ci.funcName, ci.line)
}

func retrieveCallInfo() *callInfo {
	pc, file, line, _ := runtime.Caller(2)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &callInfo{
		packageName: packageName,
		fileName:	 fileName,
		funcName:    funcName,
		line:		 line,
	}
}
