## lazyboy
Package to wrap the CouchDB initialization for inclusion into the revel framework.

## Usage
To pull this package down for use you'll need to *get* it and add it as a module to your revel application.

`go get github.com/taddevries/lazyboy`

### conf/app.conf
Add this package as a module to your revel application and configure it for use with the connection information as follows.

```
module.lazyboy=github.com/taddevries/lazyboy

couchdb.https=false        	#(bool)   optional, default=false
couchdb.url=localhost	  	#(string) required
couchdb.port=5984		  	#(string) optional, default=5984
couchdb.database=lazyboy	#(string) required
couchdb.username=admin     	#(string) optional, default=""
couchdb.password=password  	#(string) optional, default=""
```

### app/init.go
Initialize the module by importing the package and calling the `AppInit()` function.

```go
import (
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
)

//call AppInit at the bottom of init.go like so
revel.OnAppStart(lazyboy.AppInit)
```

### controllers/init.go
You'll have to create this file if you haven't already. This will setup the interceptor to call the `ControllerInit()` function when the application starts. 

```go
package controllers

import (
	"github.com/taddevries/lazyboy"
	//"github.com/robfig/revel"
)

func init() {
	lazyboy.ControllerInit()
}
```

## Acknowledgments 
I'd like to thank https://github.com/hermanschaaf for the inspiration for the layout of this package.



