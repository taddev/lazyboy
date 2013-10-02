/*
   Package to wrap the CouchDB initialization for inclusion into the revel
   framework.

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

// Package lazyboy wraps a CouchDB connection into a module for use with the
// Revel framework.
package lazyboy

import (
	"code.google.com/p/couch-go"
	"github.com/robfig/revel"
	"fmt"
)

var Database couch.Database //couchdb database object

// AppInit pulls the configuration options out of the app.conf file in a Revel
// application and stores them locally for use in the database connection.
// For optional configurations their defaults are set if they are not
// configured. For required configurations and error message is created and a
// Panic is send to Revel to kill the application.
func AppInit() {
	var found bool //logic processing

	//get url from app.conf or error out
	url := ""
	if url, found = revel.Config.String("couchdb.url"); !found {
		revel.ERROR.Panic("lazyboy: couchdb.url not defined in app.conf")
	}

	//get database name from app.conf or error out
	name := ""
	if name, found = revel.Config.String("couchdb.name"); !found {
		revel.ERROR.Panic("lazyboy: couchdb.name not defined in app.conf")
	}

	//pull configurations from app.conf or set defaults if necessary
	port := revel.Config.StringDefault("coucdb.port", "5984")
	username := revel.Config.StringDefault("couchdb.username", "")
	password := revel.Config.StringDefault("couchdb.password", "")
	https := revel.Config.BoolDefault("couchdb.https", false)

	//build credentials if username and password are set
	credentials := ""
	if username != "" && password != "" {
		credentials = fmt.Sprintf("%s:%s@", username, password)
	}

	//configure https if requested
	secure := "http"
	if https {
		secure = "https"
	}

	//build DBurl and setup couchdb connection
	dbUrl := fmt.Sprintf("%s://%s%s:%s/%s", secure, credentials, url, port, name)

	// attempt to connect to the database
	var err error
	if Database, err = couch.NewDatabaseByURL(dbUrl); err != nil {
		revel.ERROR.Panic("lazyboy: error connecting to database")
	}
}
