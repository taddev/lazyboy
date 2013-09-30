package lazyboy

import (
	"errors"
	"github.com/robfig/revel"
)

var (
	lbUrl      string
	lbPort     string
	lbUsername string
	lbPassword string
	lbName     string
)

func AppInit() {
	var found bool
	if lbUrl, found = revel.Config.String("couchdb.url"); !found {
		err := errors.New("lazyboy: Database URL not defined in app.conf")
		revel.ERROR.Panic(err)
	}

	if lbPort, found = revel.Config.String("couchdb.port"); !found {
		lbPort = "5984"
	}

	if lbUsername, found = revel.Config.String("couchdb.username"); !found {
		lbUsername = ""
	}

	if lbPassword, found = revel.Config.String("couchdb.password"); !found {
		lbPassword = ""
	}

	if lbName, found = revel.Config.String("couchdb.name"); !found {
		err := errors.New("lazyboy: Database name not defined in app.conf")
		revel.ERROR.Panic(err)
	}
}

type LazyboyController struct {
	*revel.Controller
	
}
