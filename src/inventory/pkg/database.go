package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"

	"github.com/roberthstrand/demo-msa/src/helpers"
)

const ()

type food struct {
	Name string
	Type string
}

type foods []food

func psqlInit() {
	var host string = os.Getenv(POSTGRESQL_HOST)
	var port string = os.Getenv(POSTGRESQL_PORT_NUMBER)
	var user string = os.Getenv(POSTGRESQL_USER)
	var password string = os.Getenv(POSTGRESQL_PASSWORD)
	var database string = os.Getenv(POSTGRESQL_DATABASE)

	psqlconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	db, err := sql.Open("postgres", psqlconn)

}
