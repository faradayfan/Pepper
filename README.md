## Pepper

A simple logger for go.
 
### Features

* Different logging prefixes (Info, Debug, Error)
* Create logger with config struct indicating log prefix verbosity.
* Log information can includes
  * package name
  * file name
  * function name
  * line number
* Can specify an output stream
* JSONifies message structures

### Setup

There are two ways to start using pepper.

The Easy Way
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

The Configurable Way
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

