package database

import (
	"log"
	"fmt"
	"map/config"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	conf config.Config
	db *sql.DB
}

func (p *PostgresDB) Init(conf config.Config) {
	p.conf = conf
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s default_transaction_isolation=%s",
	conf.PostgresUser, conf.PostgresPassword, conf.PostgresDbName, conf.PostgresSSL, conf.PostgresIsolationLevel)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err.Error())
		return
	}
	p.db = db
}

func (p *PostgresDB) Close() {
	p.db.Close()
}

