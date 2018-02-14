package pepper

import (
	"runtime"
	"path"
	"strings"
	"fmt"
	"strconv"
)

type Pepper interface {
	Info(message string)
	Debug(message string)
	Error(message string)
}

type pepper struct{
	InfoPrefix string
	DebugPrefix string
	ErrorPrefix string

}

type callInfo struct{
	packageName string
	fileName string
	funcName string
	line int
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
		fileName:    fileName,
		funcName:    funcName,
		line:        line,
	}
}

func formatPrefix(prefix string, ci *callInfo) string{
	return prefix + " - " + ci.fileName + " - " + ci.packageName + "." + ci.funcName + " Line " + strconv.Itoa(ci.line) + ": "
}

func (p *pepper) Info(message string){
	callInfo := retrieveCallInfo()
	fmt.Println(formatPrefix(p.InfoPrefix, callInfo) + message)
}

func (p *pepper) Error(message string){
	callInfo := retrieveCallInfo()
	fmt.Println(formatPrefix(p.ErrorPrefix, callInfo) + message)
}

func (p *pepper) Debug(message string){
	callInfo := retrieveCallInfo()
	fmt.Println(formatPrefix(p.DebugPrefix, callInfo) + message)
}

func New() Pepper {
	return &pepper{
		InfoPrefix: "Info",
		DebugPrefix: "Debug",
		ErrorPrefix: "Error",
	}
}