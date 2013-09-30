package lazyboy

import (
	"errors"
	"github.com/robfig/revel"
)

var (
	Url      string
	Port     string
	Username string
	Password string
	Name     string
)

func AppInit() {
	var found bool
	if Url, found = revel.Config.String("couchdb.url"); !found {
		err := errors.New("lazyboy: Database URL not defined in app.conf")
		revel.ERROR.Panic(err)
	}

	if Port, found = revel.Config.String("couchdb.port"); !found {
		Port = "5984"
	}

	if Username, found = revel.Config.String("couchdb.username"); !found {
		Username = ""
	}

	if Password, found = revel.Config.String("couchdb.password"); !found {
		Password = ""
	}

	if Name, found = revel.Config.String("couchdb.name"); !found {
		err := errors.New("lazyboy: Database name not defined in app.conf")
		revel.ERROR.Panic(err)
	}
}
