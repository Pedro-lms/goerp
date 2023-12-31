// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Pedro-lmso-erp/erp/src/actions"
	"github.com/Pedro-lmso-erp/erp/src/controllers"
	"github.com/Pedro-lmso-erp/erp/src/menus"
	"github.com/Pedro-lmso-erp/erp/src/models"
	"github.com/Pedro-lmso-erp/erp/src/reports"
	"github.com/Pedro-lmso-erp/erp/src/server"
	"github.com/Pedro-lmso-erp/erp/src/templates"
	"github.com/Pedro-lmso-erp/erp/src/tools/logging"
	"github.com/Pedro-lmso-erp/erp/src/views"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var driver, user, password, prefix, debug string

// RunTests initializes the database, run the tests given by m and
// tears the database down.
//
// It is meant to be used for modules testing. Initialize your module's
// tests with:
//
//	    import (
//	        "testing"
//	        "github.com/Pedro-lmso-erp/erp/src/tests"
//	    )
//
//	    func TestMain(m *testing.M) {
//		       tests.RunTests(m, "my_module")
//	    }
func RunTests(m *testing.M, moduleName string, preHookFnct func()) {
	var res int
	defer func() {
		TearDownTests(moduleName)
		if r := recover(); r != nil {
			panic(r)
		}
		os.Exit(res)
	}()
	InitializeTests(moduleName)
	if preHookFnct != nil {
		preHookFnct()
	}
	res = m.Run()

}

// InitializeTests initializes a database for the tests of the given module.
// You probably want to use RunTests instead.
func InitializeTests(moduleName string) {
	fmt.Printf("Initializing tests for module %s\n", moduleName)
	driver = os.Getenv("erp_DB_DRIVER")
	if driver == "" {
		driver = "postgres"
	}
	user = os.Getenv("erp_DB_USER")
	if user == "" {
		user = "erp"
	}
	password = os.Getenv("erp_DB_PASSWORD")
	if password == "" {
		password = "erp"
	}
	prefix = os.Getenv("erp_DB_PREFIX")
	if prefix == "" {
		prefix = "erp"
	}
	dbName := fmt.Sprintf("%s_%s_tests", prefix, moduleName)
	debug = os.Getenv("erp_DEBUG")
	logTests := os.Getenv("erp_LOG")

	viper.Set("LogLevel", "panic")
	if logTests != "" {
		viper.Set("LogLevel", "info")
		viper.Set("LogStdout", true)
	}
	if debug != "" {
		viper.Set("Debug", true)
		viper.Set("LogLevel", "debug")
		viper.Set("LogStdout", true)
	}
	logging.Initialize()

	db := sqlx.MustConnect(driver, fmt.Sprintf("dbname=postgres sslmode=disable user=%s password=%s", user, password))
	keepDB := os.Getenv("erp_KEEP_TEST_DB") != ""
	var dbExists bool
	err := db.Get(&dbExists, fmt.Sprintf("SELECT TRUE FROM pg_database WHERE datname = '%s'", dbName))
	if err != nil {
		fmt.Println(err)
	}
	if !dbExists || !keepDB {
		fmt.Println("Creating database", dbName)
		db.MustExec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName))
		db.MustExec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	}
	db.Close()

	server.PreInit()
	models.DBConnect(models.ConnectionParams{
		Driver:   driver,
		DBName:   dbName,
		User:     user,
		Password: password,
		SSLMode:  "disable",
	})
	models.BootStrap()
	resourceDir, _ := filepath.Abs(filepath.Join(".", "res"))
	server.ResourceDir = resourceDir
	server.LoadInternalResources(resourceDir)
	if !dbExists || !keepDB {
		fmt.Println("Upgrading schemas in database", dbName)
		models.SyncDatabase()
		fmt.Println("Loading resources into database", dbName)
		server.LoadDataRecords(resourceDir)
		server.LoadDemoRecords(resourceDir)
	}
	views.BootStrap()
	templates.BootStrap()
	actions.BootStrap()
	reports.BootStrap()
	controllers.BootStrap()
	menus.BootStrap()
	server.PostInit()
}

// TearDownTests tears down the tests for the given module
func TearDownTests(moduleName string) {
	models.DBClose()
	keepDB := os.Getenv("erp_KEEP_TEST_DB")
	if keepDB != "" {
		return
	}
	fmt.Printf("Tearing down database for module %s...", moduleName)
	dbName := fmt.Sprintf("%s_%s_tests", prefix, moduleName)
	db := sqlx.MustConnect(driver, fmt.Sprintf("dbname=postgres sslmode=disable user=%s password=%s", user, password))
	db.MustExec(fmt.Sprintf("DROP DATABASE %s", dbName))
	db.Close()
	fmt.Println("Ok")
}
