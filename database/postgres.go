package database

import (
	"database/sql"
	"fmt"
	"log"
	"map/config"
	// "map/point"
	// "time"
	
	// "github.com/lib/pq"
)

type PostgresDB struct {
	conf config.Config
	db *sql.DB
}

func (p *PostgresDB) UsedConn() (int, int, int) {
	idle := p.db.Stats().Idle
	use := p.db.Stats().InUse
	open := p.db.Stats().OpenConnections
	return use, open, idle
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
	// go func() {
	// 	for {
	// 		time.Sleep(100 * time.Millisecond)
	// 		use, open, idle := p.UsedConn()
	// 		fmt.Println(use, open, idle)
	// 	}
	// }()
}

func (p *PostgresDB) Close() {
	p.db.Close()
}



