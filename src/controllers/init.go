// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package controllers

import (
	"github.com/Pedro-lmso-erp/erp/src/server"
	"github.com/Pedro-lmso-erp/erp/src/tools/logging"
)

var log logging.Logger

// BootStrap creates the actual controllers from the controllers registry.
// This function must be called before starting the http server.
func BootStrap() {
	Registry.createRoutes(server.GetServer().Group("/"))
}

func init() {
	log = logging.GetLogger("controllers")
	Registry = newGroup("/")
}
