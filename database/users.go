package database

import (
	"map/business"
	"log"
)

func (p *PostgresDB) GetUsersInfo() (business.UsersInfo, error) {
	var result business.UsersInfo
	
	rows, err := p.db.Query(`select id, login, role, active,
	name, surname, patronymic, tg_id from users`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.UserInfo
		err := rows.Scan(&res.ID, &res.Login, &res.Role, &res.Active,
			&res.Name, &res.Surname, &res.Patronymic, &res.TgID)
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
	name, surname, patronymic, tg_id from users where id = $1`, id)
	err := row.Scan(&res.ID, &res.Login, &res.Role, &res.Active,
		&res.Name, &res.Surname, &res.Patronymic, &res.TgID)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}