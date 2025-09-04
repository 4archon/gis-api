package database

import (
	"log"
	// "strconv"
)

func (p *PostgresDB) GetUserLogin(id int) string {
	res := p.db.QueryRow(`select login from users where id = $1;`, id)
	var email *string
	err := res.Scan(&email)
	if err != nil {
		log.Println(err)
		return ""
	}
	if email == nil {
		return ""
	}
	return *email
}

func (p *PostgresDB) ChangeUserPassword(id int, password string) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = tx.Exec(`update users set password = $1 where id = $2`, password, id)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
