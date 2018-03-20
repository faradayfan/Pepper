## Pepper

A simple logger for go.
 
### Features

* RFC 5424 Loggin levels
* Create logger with config struct indicating log prefix verbosity.
* In addition to log level prefixes, log information can includes
  * package name
  * file name
  * function name
  * line number
* Can specify an output stream
* JSONifies message structures
* Intelligently handles error objects and other objects that stringify itself.

### Setup

There are three ways to start using pepper.

#### The Easy Way
```go
package test

import (
	"github.com/faradayfan/Pepper"
	)
func main(){
	log := pepper.NewDefault()
	
	log.Info("this is a message")
}

```

#### The Log Level Way
```go
package test

import (
	"github.com/faradayfan/Pepper"
	)
func main(){
	// log levels are 0 - 7 ( RFC 5424 )
	log := pepper.NewDefaultLevel(pepper.Debug)
	
	log.Info("this is a message")
}

```

#### The Configurable Way
```go
package test

import (
	"github.com/faradayfan/Pepper"
	)
func main(){


    prefix := &Prefix{
            FileName: true,
            PackageName: true,
            FunctionName: true,
            LineNumber: true,
    }
    
    config := &Config{
        Prefix: prefix,
        Output: os.Stdout,
    }
    
    log := pepper.New(config)
    
    // a simple string message
    log.Info("this is a message")
    
    
    // a complex structure
    type testStruct struct{
        Message string
        Data string
    }
    var test = testStruct{
        "this is just a test obj message.",
        "here is some data",
    }
    
    log.Debug(test)
    
    log.Error(config)
    
}
```

