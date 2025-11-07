package database

import (
	"map/business"
	"log"
)

func (p *PostgresDB) GetPointCurrentAppoint(id int) (business.PointAppoints, error) {
	var result business.PointAppoints

	rows, err := p.db.Query(`select id, appointment_date
	from service where sent is false and point_id = $1`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.PointAppoint
		err := rows.Scan(&res.ID, &res.AppointmentDate)
		if err != nil {
			log.Println(err)
			return result, err
		}
		result.Appoints = append(result.Appoints, res)
	}

	storage := make(map[int]*business.PointAppoint)
	for i, appoint := range result.Appoints {
		storage[appoint.ID] = &result.Appoints[i]
	}


	rows2, err := p.db.Query(`select s.id, u.id, login, subgroup, trust
	from service s, users u where s.sent is false and s.point_id = $1
	and u.id = any(user_id)`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var res business.PointAppointUser
		var serviceID int
		err := rows2.Scan(&serviceID, &res.ID, &res.Login, &res.Subgroup, &res.Trust)
		if err != nil {
			log.Println(err)
			return result, err
		}

		storage[serviceID].Users = append(storage[serviceID].Users, res)
	}

	return result, nil
}

func (p *PostgresDB) DeletePointAppoint(id int) (error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(`delete from service where id = $1`, id)
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