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
	System(messages ...interface{})
	Alert(messages ...interface{})
	Critical(messages ...interface{})
	Error(messages ...interface{})
	Warning(messages ...interface{})
	Notice(messages ...interface{})
	Info(messages ...interface{})
	Debug(messages ...interface{})
}

const (
	System = 0
	Alert = 1
	Critical = 2
	Error = 3
	Warning = 4
	Notice = 5
	Info = 6
	Debug = 7
)

type levelPrefixes struct {
	system   string
	alert    string
	critical string
	error    string
	warning  string
	notice   string
	info     string
	debug    string
}

func getLevelPrefixes() *levelPrefixes {
	return &levelPrefixes{
		"System",   	// 0
		"Alert",    	// 1
		"Critical",	// 2
		"Error",		// 3
		"Warning",	// 4
		"Notice",		// 5
		"Info",		// 6
		"Debug",		// 7
	}
}

type pepper struct {
	Config        *Config
	levelPrefixes *levelPrefixes
}

type Config struct {
	Prefix *Prefix
	Output *os.File
	Level  int
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

func (p *pepper) printMsg(prefix string, callInfo *callInfo, messages []interface{}) {
	fmt.Fprintln(p.Config.Output, fmt.Sprintf("%s%s", formatPrefix(prefix, callInfo, p.Config), cleanPrint(messages)))
}

func (p *pepper) System(messages ...interface{}) {
	callInfo := retrieveCallInfo()
	p.printMsg(p.levelPrefixes.system, callInfo, messages)
}

func (p *pepper) Alert(messages ...interface{}) {
	if p.Config.Level >= Alert {
		callInfo := retrieveCallInfo()
		p.printMsg(p.levelPrefixes.alert, callInfo, messages)
	}
}

func (p *pepper) Critical(messages ...interface{}) {
	if p.Config.Level >= Critical {
		callInfo := retrieveCallInfo()
		p.printMsg(p.levelPrefixes.critical, callInfo, messages)
	}
}

func (p *pepper) Error(messages ...interface{}) {
	if p.Config.Level >= Error {
		callInfo := retrieveCallInfo()
		p.printMsg(p.levelPrefixes.error, callInfo, messages)
	}
}

func (p *pepper) Warning(messages ...interface{}) {
	if p.Config.Level >= Warning{
		callInfo := retrieveCallInfo()
		p.printMsg(p.levelPrefixes.warning, callInfo, messages)
	}
}

func (p *pepper) Notice(messages ...interface{}) {
	if p.Config.Level >= Notice {
		callInfo := retrieveCallInfo()
		p.printMsg(p.levelPrefixes.notice, callInfo, messages)
	}
}

func (p *pepper) Info(messages ...interface{}) {
	if p.Config.Level >= Info {
		callInfo := retrieveCallInfo()
		p.printMsg(p.levelPrefixes.info, callInfo, messages)
	}
}

func (p *pepper) Debug(messages ...interface{}) {
	if p.Config.Level >= Debug {
		callInfo := retrieveCallInfo()
		p.printMsg(p.levelPrefixes.debug, callInfo, messages)
	}
}

func New(config *Config) Pepper {
	p := getLevelPrefixes()
	return &pepper{
		Config:        config,
		levelPrefixes: p,
	}
}

func NewDefaultLevel(level int) Pepper {
	config := &Config{
		&Prefix{true, true, true, true},
		os.Stdout,
		level,
	}
	p := getLevelPrefixes()
	return &pepper{
		Config:        config,
		levelPrefixes: p,
	}
}

func NewDefault() Pepper {
	config := &Config{
		&Prefix{true, true, true, true},
		os.Stdout,
		7,
	}
	p := getLevelPrefixes()
	return &pepper{
		Config:        config,
		levelPrefixes: p,
	}
}
