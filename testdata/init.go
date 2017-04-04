package testdata

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ederavilaprado/golang-web-architecture-template/app"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // initialize posgresql for test
)

var (
	DB *sqlx.DB
)

func init() {
	// the test may be started from the home directory or a subdirectory
	err := app.LoadConfig("./config", "../config")
	if err != nil {
		panic(err)
	}
	DB, err = sqlx.Connect("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
}

// ResetDB re-create the database schema and re-populate the initial data using the SQL statements in db.sql.
// This method is mainly used in tests.
func ResetDB() *sqlx.DB {
	if err := runSQLFile(DB, getSQLFile()); err != nil {
		panic(fmt.Errorf("Error while initializing test database: %s", err))
	}
	return DB
}

func getSQLFile() string {
	if _, err := os.Stat("testdata/db.sql"); err == nil {
		return "testdata/db.sql"
	}
	return "../testdata/db.sql"
}

func runSQLFile(db *sqlx.DB, file string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	lines := strings.Split(string(s), ";")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if _, err := db.Exec(line); err != nil {
			fmt.Println(line)
			return err
		}
	}
	return nil
}
