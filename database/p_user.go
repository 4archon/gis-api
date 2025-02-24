package database

import (
	"log"
	"map/point"
	"database/sql"
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
	tg_id from users;`)
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