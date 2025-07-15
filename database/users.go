package database

import (
	"map/business"
	"log"
)

func (p *PostgresDB) GetUsersInfo() (business.UsersInfo, error) {
	var result business.UsersInfo
	
	rows, err := p.db.Query(`select id, login, role, active,
	name, surname, patronymic, tg_id,
	subgroup, trust from users`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.UserInfo
		err := rows.Scan(&res.ID, &res.Login, &res.Role, &res.Active,
			&res.Name, &res.Surname, &res.Patronymic, &res.TgID,
			&res.Subgroup, &res.Trust)
		if err != nil {
			log.Println(err)
			return result, err
		}
		result.Info = append(result.Info, res)
	}
	return result, nil
}

func (p *PostgresDB) GetUserInfo(id int) (business.UserInfo, error) {
	var res business.UserInfo
	
	row := p.db.QueryRow(`select id, login, role, active,
	name, surname, patronymic, tg_id,
	subgroup, trust
	from users where id = $1`, id)
	err := row.Scan(&res.ID, &res.Login, &res.Role, &res.Active,
		&res.Name, &res.Surname, &res.Patronymic, &res.TgID,
		&res.Subgroup, &res.Trust)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}

func (p *PostgresDB) CreateNewUser(user business.User) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	row := tx.QueryRow(`insert into users values(default,
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning id`,
	user.Login, user.Password, user.Role, user.Active,
	user.Name, user.Surname, user.Patronymic, user.TgID,
	user.Subgroup, user.Trust)
	var id int
	err = row.Scan(&id)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Println(err2)
			return 0, err2
		}
		log.Println(err)
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}

func (p *PostgresDB) ChangeUser(user business.User) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	if *user.Password != "" {
		_, err = tx.Exec(`update users set login = $1, password = $2, role = $3, active = $4,
		name = $5, surname = $6, patronymic = $7, tg_id = $8,
		subgroup = $9, trust = $10 where id = $11`,
		user.Login, user.Password, user.Role, user.Active,
		user.Name, user.Surname, user.Patronymic, user.TgID,
		user.Subgroup, user.Trust, user.ID)
	} else {
		_, err = tx.Exec(`update users set login = $1, role = $2, active = $3,
		name = $4, surname = $5, patronymic = $6, tg_id = $7,
		subgroup = $8, trust = $9 where id = $10`,
		user.Login, user.Role, user.Active,
		user.Name, user.Surname, user.Patronymic, user.TgID,
		user.Subgroup, user.Trust, user.ID)
	}
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Println(err2)
			return err2
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

func (p *PostgresDB) ChangeUserProfile(user business.User) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	if *user.Password != "" {
		_, err = tx.Exec(`update users set password = $1,
		name = $2, surname = $3, patronymic = $4, tg_id = $5 where id = $6`,
		user.Password,
		user.Name, user.Surname, user.Patronymic, user.TgID, user.ID)
	} else {
		_, err = tx.Exec(`update users set
		name = $1, surname = $2, patronymic = $3, tg_id = $4 where id = $5`,
		user.Name, user.Surname, user.Patronymic, user.TgID, user.ID)
	}
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Println(err2)
			return err2
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

func (p *PostgresDB) GetUserSubgroup(id int) (string, error) {
	var result *string
	row := p.db.QueryRow(`select subgroup from users where id = $1`, id)
	err := row.Scan(&result)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if result == nil {
		return "", nil
	}

	return *result, nil
}

func (p *PostgresDB) GetUserTrust(id int) (bool, error) {
	var result *bool
	row := p.db.QueryRow(`select trust from users where id = $1`, id)
	err := row.Scan(&result)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if result == nil {
		return false, nil
	}

	return *result, nil
}