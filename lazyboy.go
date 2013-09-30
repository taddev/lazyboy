/*
   Package to wrap the CouchDB initialization for inclusion into the revel framework.
   Copyright (C) 2013  Tad DeVries <tad@splunk.net>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see [http://www.gnu.org/licenses/].
*/

package lazyboy

import (
	"code.google.com/p/couch-go"
	"errors"
	"github.com/robfig/revel"
)

var (
	Url      string
	Port     string //optional, will be set to 5984 by default
	Username string
	Password string
	Database string
	Https    bool //optional, will be set to false by default
)

func AppInit() {
	var found bool
	if Url, found = revel.Config.String("couchdb.url"); !found {
		err := errors.New("lazyboy: couchdb.url not defined in app.conf")
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

	if Database, found = revel.Config.String("couchdb.database"); !found {
		err := errors.New("lazyboy: couchdb.database not defined in app.conf")
		revel.ERROR.Panic(err)
	}

	if Https, found = revel.Config.Bool("couchdb.https"); !found {
		Https = false
	}
}

func ControllerInit() {
	revel.InterceptMethod((*CouchDBController).Begin, revel.BEFORE) //function to open the database at the start
	//revel.InterceptMethod((*CouchDBController).End, revel.FINALLY) //function to close the database at the end
}

type CouchDBController struct {
	*revel.Controller
	Database couch.Database
	DBUrl    string
}

func (c *CouchDBController) Begin() revel.Result {
	var credentials string
	var secure string
	if Username != "" && Password != "" {
		credentials = Username + ":" + Password + "@"
	} else {
		credentials = ""
	}

	if Https {
		secure = "https://"
	} else {
		secure = "http://"
	}

	c.DBUrl = secure + credentials + Url + ":" + Port + "/" + Database
	c.Database, _ = couch.NewDatabaseByURL(c.DBUrl)

	return nil
}
