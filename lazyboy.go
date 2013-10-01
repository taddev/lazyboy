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

// Package lazyboy wraps a CouchDB connection into a module for use with the
// Revel framework.
package lazyboy

import (
	"code.google.com/p/couch-go"
	"github.com/robfig/revel"
)

var (
	DBUrl    string         //url for couchdb connection
	Database couch.Database //couchdb database object
)

// AppInit pulls the configuration options out of the app.conf file in a Revel application
// and stores them locally for use in the database connection. For optional configurations
// their defaults are set if they are not configured. For required configurations and error
// message is created and a Panic is send to Revel to kill the application.
func AppInit() {
	var (
		url         string //required, the URL or IP address to the CouchDB instance
		port        string //optional, the port the CouchDB instance is listening on (default=5984)
		username    string //optional, the username for access to the database (default="")
		password    string //optional, the password for access to the database (default="")
		name        string //required, the name of the database
		https       bool   //optional, uses HTTPS to make the database connections (default="false")
		found       bool
		credentials string
		secure      string
	)

	if url, found = revel.Config.String("couchdb.url"); !found {
		//err := errors.New("lazyboy: couchdb.url not defined in app.conf")
		revel.ERROR.Panic("lazyboy: couchdb.url not defined in app.conf")
	}

	if name, found = revel.Config.String("couchdb.database"); !found {
		//err := errors.New("lazyboy: couchdb.database not defined in app.conf")
		revel.ERROR.Panic("lazyboy: couchdb.database not defined in app.conf")
	}

	/*
		if Port, found = revel.Config.String("couchdb.port"); !found {
			Port = "5984"
		}

		if Username, found = revel.Config.String("couchdb.username"); !found {
			Username = ""
		}

		if Password, found = revel.Config.String("couchdb.password"); !found {
			Password = ""
		}

		if Https, found = revel.Config.Bool("couchdb.https"); !found {
			Https = false
		}
	*/

	port = revel.Config.StringDefault("coucdb.port", "5984")
	username = revel.Config.StringDefault("couchdb.username", "")
	password = revel.Config.StringDefault("couchdb.password", "")
	https = revel.Config.BoolDefault("couchdb.https", false)

	//build credentials if username and password are set
	if username != "" && password != "" {
		credentials = username + ":" + password + "@"
	} else {
		credentials = ""
	}

	//configure https if requested
	if https {
		secure = "https://"
	} else {
		secure = "http://"
	}

	//build DBurl and setup couchdb connection
	DBUrl = secure + credentials + url + ":" + port + "/" + name
	Database, _ = couch.NewDatabaseByURL(DBUrl)
}

/*
// ControllerInit defines the Interceptor to start the database connection at
// the begining of each request.
func ControllerInit() {
	revel.InterceptMethod((*CouchDBController).Begin, revel.BEFORE) //function to open the database at the start
	//revel.InterceptMethod((*CouchDBController).End, revel.FINALLY) //db cleanup at the end of a request
}

// CouchDBController stores information about the Revel application controller
// and adds to it the CouchDB connection object and a string containg the URL
// used in the connection. DBUrl might be removed in the future, it's primarily
// for testing right now.
type CouchDBController struct {
	*revel.Controller
	Database couch.Database
}

// Begin is an interceptor function called at the start of a Revel request.
// If the the username and password are defined they are constructed into a
// credentials string to be added to the database connection URL. The HTTPS
// option is also set if defined. The connection URL string is stored into
// DBUrl and the database object is stored in Database.
func (c *CouchDBController) Begin() revel.Result {

	c.Database, _ = couch.NewDatabaseByURL(DBUrl)

	return nil
}
*/
