## lazyboy
Package to wrap the CouchDB initialization for inclusion into the revel framework.

## Usage
To use this package you'll need to *get* it from here and add it as a module to your revel application.

`go get github.com/taddevries/lazyboy`

### conf/app.conf
Add this package as a module to your revel application and configure it for use with the connection information as follows.

```
module.lazyboy=github.com/taddevries/lazyboy

couchdb.https=false         #(bool)   optional, default=false
couchdb.url=localhost       #(string) required
couchdb.port=5984           #(string) optional, default=5984
couchdb.database=lazyboy    #(string) required
couchdb.username=admin      #(string) optional, default=""
couchdb.password=password   #(string) optional, default=""
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

## Acknowledgments 
Thanks to [Jeff Graham][1] and [Herman Schaaf][2] for their work on [various][3] [revmgo][4] packages. These were very helpful in understanding how to build a custom module in revel. I realize that MangoDB and CouchDB are quite different in how they handle their connections, but their work was still helpful.

Thanks to [Rob Figueiredo][6] for his fine work on [Revel][5]. It is a very interesting and fun framework to work with.

Thanks to Andy R. on the Revel group list for helping me better understand how to make things work correctly.


## Change Log
**20130929** Initial upload

**20131001** I've rearranged things quite a bit after talking with Andy R. on the Revel Groups list. I realized that by having things run through an interceptor I was essentially repeating a lot of static items on every request. I moved everything into the `AppInit()` function and have it calling the `NewDatabaseByURL()` only once, by my understanding of the `revel.OnAppStart` function. 

## License
Copyright (C) 2013  Tad DeVries <tad@splunk.net>

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see [http://www.gnu.org/licenses/].

<!-- Links -->
[1]: https://github.com/jgraham909 "Jeff Graham"
[2]: https://github.com/hermanschaaf "Herman Schaaf"
[3]: https://github.com/jgraham909/revmgo "Jeff's revmgo"
[4]: https://github.com/hermanschaaf/revmgo "Herman's revmgo"
[5]: https://github.com/robfig/revel "Revel Framework"
[6]: https://github.com/robfig "Rob Figueiredo"