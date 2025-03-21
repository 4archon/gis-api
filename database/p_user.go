package database

import (
	"database/sql"
	"log"
	"map/point"
	"strconv"
)

func (p *PostgresDB) GetUserLogin(id int) string {
	res := p.db.QueryRow(`select login from users where id = $1;`, id)
	var email string
	err := res.Scan(&email)
	if err != nil {
		log.Println(err)
		return ""
	}
	return email
}

func (p *PostgresDB) GetUsersInfo() []point.User {
	rows, err := p.db.Query(`select id, login, role, active, name, surname, patronymic,
	tg_id from users order by login;`)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	var users []point.User
	for rows.Next() {
		var user point.User
		var sqlLogin, sqlRole, sqlName, sqlSurname, sqlPatr sql.NullString
		var sqlActive sql.NullBool
		var sqlTgID	sql.NullInt64
		err = rows.Scan(&user.ID, &sqlLogin, &sqlRole, &sqlActive, &sqlName,
		&sqlSurname, &sqlPatr, &sqlTgID)
		if err != nil {
			log.Println(err)
			return nil
		}
		if sqlActive.Valid {
			if sqlActive.Bool {
				user.Active = "Активный"
			} else {user.Active = "Деактивирован"}
		} else {
			user.Active = "Деактивирован"
		}
		if sqlTgID.Valid {
			user.TgID = sqlTgID.Int64
		} else {
			user.TgID = -1
		}
		if sqlLogin.Valid {user.Login = sqlLogin.String}
		if sqlRole.Valid {user.Role = sqlRole.String}
		if sqlName.Valid {user.Name = sqlName.String}
		if sqlSurname.Valid {user.Surname = sqlSurname.String}
		if sqlPatr.Valid {user.Patronymic = sqlPatr.String}
		users = append(users, user)
	}
	return users
}

func (p *PostgresDB) GetWorkersInfo() []point.User {
	rows, err := p.db.Query(`select id, login, role, active, name, surname, patronymic,
	tg_id from users where role = 'worker' order by login;`)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	var users []point.User
	for rows.Next() {
		var user point.User
		var sqlLogin, sqlRole, sqlName, sqlSurname, sqlPatr sql.NullString
		var sqlActive sql.NullBool
		var sqlTgID	sql.NullInt64
		err = rows.Scan(&user.ID, &sqlLogin, &sqlRole, &sqlActive, &sqlName,
		&sqlSurname, &sqlPatr, &sqlTgID)
		if err != nil {
			log.Println(err)
			return nil
		}
		if sqlActive.Valid {
			if sqlActive.Bool {
				user.Active = "Активный"
			} else {user.Active = "Деактивирован"}
		} else {
			user.Active = "Деактивирован"
		}
		if sqlTgID.Valid {
			user.TgID = sqlTgID.Int64
		} else {
			user.TgID = -1
		}
		if sqlLogin.Valid {user.Login = sqlLogin.String}
		if sqlRole.Valid {user.Role = sqlRole.String}
		if sqlName.Valid {user.Name = sqlName.String}
		if sqlSurname.Valid {user.Surname = sqlSurname.String}
		if sqlPatr.Valid {user.Patronymic = sqlPatr.String}
		users = append(users, user)
	}
	return users
}

func (p *PostgresDB) GetUserInfo(id int) (point.User, error) {
	row := p.db.QueryRow(`select id, login, role, active, name, surname, patronymic,
	tg_id from users where id = $1;`, id)
	var user point.User
	var sqlLogin, sqlRole, sqlName, sqlSurname, sqlPatr sql.NullString
	var sqlActive sql.NullBool
	var sqlTgID	sql.NullInt64
	err := row.Scan(&user.ID, &sqlLogin, &sqlRole, &sqlActive, &sqlName,
		&sqlSurname, &sqlPatr, &sqlTgID)
	if err != nil {
		log.Println(err)
		return user, err
	}

	if sqlActive.Valid {
		if sqlActive.Bool {
			user.Active = "Активный"
		} else {user.Active = "Деактивирован"}
	} else {
		user.Active = "Деактивирован"
	}
	if sqlTgID.Valid {
		user.TgID = sqlTgID.Int64
	} else {
		user.TgID = -1
	}
	if sqlLogin.Valid {user.Login = sqlLogin.String}
	if sqlRole.Valid {user.Role = sqlRole.String}
	if sqlName.Valid {user.Name = sqlName.String}
	if sqlSurname.Valid {user.Surname = sqlSurname.String}
	if sqlPatr.Valid {user.Patronymic = sqlPatr.String}
	
	return user, nil
}


func (p *PostgresDB) ChangeUserInfo(id int, name string, surname string,
	patronymic string, tgID string) error {
	user, err := p.GetUserInfo(id)
	if err != nil {
		log.Println(err)
		return err
	}

	if name == "" {name = user.Name}
	if surname == "" {surname = user.Surname}
	if patronymic == "" {patronymic = user.Patronymic}
	if tgID == "" {tgID = strconv.FormatInt(user.TgID, 10)}

	tx, err :=p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = tx.Exec(`update users set name = $1, surname = $2, patronymic = $3, tg_id = $4
	where id = $5`, name, surname, patronymic, tgID, id)
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

func (p *PostgresDB) NewUser(user point.NewUser) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	if user.TgID == "" {
		user.TgID = "-1"
	}
	_, err = tx.Exec(`insert into users values(default, $1, $2, $3, $4, $5, $6, $7, $8)`,
	user.Login, user.Password, user.Role, user.Active,
	user.Name, user.Surname, user.Patronymic, user.TgID)
	if err != nil {
		log.Println(err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Println(err2)
			return err
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *PostgresDB) ChangeUserAllInfo(id int, user point.NewUser) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	if user.Role != "current" {
		_, err = tx.Exec(`update users set role = $1 where id = $2`, user.Role, id)
		if err != nil {
			log.Println(err)
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err)
				return err
			}
			return err
		}
	}
	if user.Active != "current" {
		_, err = tx.Exec(`update users set active = $1 where id = $2`, user.Active, id)
		if err != nil {
			log.Println(err)
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err)
				return err
			}
			return err2
		}
	}
	_, err = tx.Exec(`update users set login = $1,
	name = $2, surname = $3, patronymic = $4, tg_id = $5 where id = $6;`,
	user.Login,
	user.Name, user.Surname, user.Patronymic, user.TgID, id)
	if err != nil {
		log.Println(err)
		err2 := tx.Rollback()
		if err2 != nil {
			log.Println(err)
			return err
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}