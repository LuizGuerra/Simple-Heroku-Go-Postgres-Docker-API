package db

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"os"
)

func NewDB() (*pg.DB, error) {
	var opts *pg.Options
	var err error
	if os.Getenv("ENV") == "PROD" {
		opts, err = pg.ParseURL(os.Getenv("DATABASE_URL"))
		if err != nil {
			return nil, err
		}
	} else {
		opts = &pg.Options{
			Addr:     "db:5432",
			User:     "postgres",
			Password: "admin",
		}
	}

	// connect to db
	db := pg.Connect(opts)

	// run migrations
	collection := migrations.NewCollection()
	err = collection.DiscoverSQLMigrations("migrations")
	if err != nil {
		return nil, err
	}
	_, _, err = collection.Run(db, "init")
	if err != nil {
		return nil, err
	}
	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		return nil, err
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}

	// return the db connection
	return db, nil
}
