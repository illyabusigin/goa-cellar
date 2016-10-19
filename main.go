// +build !appengine,!appenginevm

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/kit"
	"github.com/goadesign/goa/middleware"
	"github.com/illyabusigin/goa-cellar/app"
	"github.com/illyabusigin/goa-cellar/controllers"
	"github.com/illyabusigin/goa-cellar/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DB is the application database instance
	DB *gorm.DB
)

func main() {
	// Setup logger
	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)

	// Configure/connect to DB
	dsn, dsnErr := normalizeDSN(os.Getenv("DATABASE_URL"))
	if dsnErr != nil {
		fmt.Printf("Unable to parse DSN!! Error: %v\n", dsnErr.Error())
		os.Exit(1)
	}

	if dbErr := connectToDB(dsn); dbErr != nil {
		fmt.Println("Unable to connect to DB! :( Error:", dbErr)
		os.Exit(1)
	}

	DB.LogMode(true)

	// Perform automigrations.
	DB.DropTable(&models.Account{}, &models.Bottle{})
	DB.AutoMigrate(&models.Account{}, &models.Bottle{})

	// Create goa service
	service := goa.New("cellar")
	service.WithLogger(goakit.New(logger))

	// Setup basic middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount account controller onto service
	ac := controllers.NewAccount(service, DB)
	app.MountAccountController(service, ac)

	// Mount bottle controller onto service
	bc := controllers.NewBottle(service, DB)
	app.MountBottleController(service, bc)

	// Mount public controller onto service
	pc := controllers.NewPublic(service)
	app.MountPublicController(service, pc)

	// Mount js controller onto service
	jc := controllers.NewJs(service)
	app.MountJsController(service, jc)

	// Mount swagger controller onto service
	sc := controllers.NewSwagger(service)
	app.MountSwaggerController(service, sc)

	// Run service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError(err.Error())
	}
}

func connectToDB(dsn string) error {
	ordinal := func(x int) string {
		suffix := "th"
		switch x % 10 {
		case 1:
			if x%100 != 11 {
				suffix = "st"
			}
		case 2:
			if x%100 != 12 {
				suffix = "nd"
			}
		case 3:
			if x%100 != 13 {
				suffix = "rd"
			}
		}
		return strconv.Itoa(x) + suffix
	}

	maxDBAttempts := 20

	dbConnect := func() error {
		db, err := gorm.Open("mysql", dsn)
		if err != nil {
			return err
		}

		DB = db

		return nil
	}

	var dbError error

	for attempt := 1; attempt <= maxDBAttempts; attempt++ {
		dbError = dbConnect()
		if dbError == nil && DB.DB().Ping() == nil {
			fmt.Printf("Successfully connected to database on %s attempt!\n", ordinal(attempt))
			break
		}

		// Exponential backoff
		waitTime := time.Duration(attempt) * time.Second

		fmt.Printf("Database connection Error: <%s>. Attempt %d/%d. Trying again in %d seconds. ", dbError.Error(), attempt, maxDBAttempts, waitTime/time.Second)
		time.Sleep(waitTime)
	}

	return dbError
}
