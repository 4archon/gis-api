package database

import ()

func (p *PostgresDB) GetAuth(email string, password string) (int, string, error) {
	res := p.db.QueryRow(`select id, role from users where login = $1 and password = $2
	and active = 't';`,email, password)
	var id int
	var role string
	err := res.Scan(&id, &role)
	if err != nil {
		return 0, "", err
	}
	return id, role, err
}

func (p *PostgresDB) CheckActiveAuth(id int, role string) bool {
	res := p.db.QueryRow(`select count(*) from users where id = $1 and role = $2
	and active = 't';`, id, role)
	var count int
	err := res.Scan(&count)
	return err == nil
}