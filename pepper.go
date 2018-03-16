package pepper

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type Pepper interface {
	Info(messages ...interface{})
	Debug(messages ...interface{})
	Error(messages ...interface{})
}

type pepper struct {
	Config      *Config
	InfoPrefix  string
	DebugPrefix string
	ErrorPrefix string
}

type Config struct {
	Prefix *Prefix
	Output *os.File
}

type Prefix struct {
	FileName     bool
	PackageName  bool
	FunctionName bool
	LineNumber   bool
}

type callInfo struct {
	packageName string
	fileName    string
	funcName    string
	line        int
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

func cleanPrint(objs []interface{}) string {
	newObjs := make([]string, len(objs))
	for i, obj := range objs {
		newObjs[i] = cleanPrintSingle(obj)
	}
	return jsonify(newObjs)
}

func cleanPrintSingle(obj interface{}) string {
	switch v := obj.(type) {
	case string:
		return v
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	case fmt.GoStringer:
		return v.GoString()
	default:
		return jsonify(obj)
	}
}

func jsonify(obj interface{}) string {
	val, _ := json.Marshal(obj)
	return string(val)
}

func formatPrefix(prefix string, ci *callInfo, config *Config) string {
	result := prefix
	if config.Prefix.FileName {
		result = result + " - " + ci.fileName
	}
	if config.Prefix.PackageName {
		result = result + " - " + ci.packageName
		if config.Prefix.FunctionName {
			result = result + "."
		}
	}
	if config.Prefix.FunctionName {
		result = result + ci.funcName
	}
	if config.Prefix.LineNumber {
		result = result + " Line " + strconv.Itoa(ci.line)
	}
	return result + ": "
}

func (p *pepper) Info(messages ...interface{}) {
	callInfo := retrieveCallInfo()
	fmt.Fprintln(p.Config.Output, fmt.Sprintf("%s%s", formatPrefix(p.InfoPrefix, callInfo, p.Config), cleanPrint(messages)))
}

func (p *pepper) Error(messages ...interface{}) {
	callInfo := retrieveCallInfo()
	fmt.Fprintln(p.Config.Output, fmt.Sprintf("%s%s", formatPrefix(p.ErrorPrefix, callInfo, p.Config), cleanPrint(messages)))
}

func (p *pepper) Debug(messages ...interface{}) {
	callInfo := retrieveCallInfo()
	fmt.Fprintln(p.Config.Output, fmt.Sprintf("%s%s", formatPrefix(p.DebugPrefix, callInfo, p.Config), cleanPrint(messages)))
}

func New(config *Config) Pepper {
	return &pepper{
		Config:      config,
		InfoPrefix:  "Info",
		DebugPrefix: "Debug",
		ErrorPrefix: "Error",
	}
}

func NewDefault() Pepper {
	config := &Config{
		&Prefix{true, true, true, true},
		os.Stdout,
	}
	return &pepper{
		Config:      config,
		InfoPrefix:  "Info",
		DebugPrefix: "Debug",
		ErrorPrefix: "Error",
	}
}
