## lazyboy
Package to wrap the CouchDB initialization for inclusion into the revel framework.

**Table of Contents**

1. [Usage](#usage)
    1. [conf/app.conf](#confappconf)
    2. [app/init.go](#appinitgo)
    3. [controllers/init.go](#controllersinitgo)
2. [Acknowledgments](#acknowledgments)
3. [License](#license)


## Usage
To clone this repository and play with the code of source your would have to *get* it.

`go get github.com/taddev/lazyboy`

Otherwise you can just *install* the package into your $GOPATH.

`go install github.com/taddev/lazyboy`

### conf/app.conf
Add this package as a module to your revel application and configure it for use with the connection information as follows.

```
module.lazyboy=github.com/taddev/lazyboy

couchdb.https=false         #(bool)   optional, default=false
couchdb.url=localhost       #(string) required
couchdb.port=5984           #(string) optional, default=5984
couchdb.name=lazyboy        #(string) required
couchdb.username=admin      #(string) optional, default=""
couchdb.password=password   #(string) optional, default=""
```

### controllers/init.go
Initialize the module by importing the package and calling `lazyboy.AppInit` in the `revel.OnAppStart()` function.

```go
import (
	"github.com/revel/revel"
	"github.com/taddev/lazyboy"
)

//call AppInit at the bottom of init.go like so
revel.OnAppStart(lazyboy.AppInit)
```

## Acknowledgments 
Thanks to [Rob Figueiredo][1] for his fine work on [Revel][2]. It is a very interesting and fun framework to work with.

## Change Log
**20130929** Initial upload

**20131001** I've rearranged things quite a bit after talking with Andy R. on the Revel Groups list. I realized that by having things run through an interceptor I was essentially repeating a lot of static items on every request. I moved everything into the `AppInit()` function and have it calling the `NewDatabaseByURL()` only once, by my understanding of the `revel.OnAppStart` function. 

## License
Copyright (C) 2013-2014  Tad DeVries <tad@splunk.net>
http://tad.mit-license.org/2013


<!-- Links -->
[1]: https://github.com/robfig "Rob Figueiredo"
[2]: https://github.com/robfig/revel "Revel Framework"
