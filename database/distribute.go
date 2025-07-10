package database

import (
	"log"
	"map/business"
	"time"

	"github.com/lib/pq"
)

func (p *PostgresDB) GetDataForDistribute() ([]business.DistibutePoint, error) {
	var result []business.DistibutePoint
	
	rows, err := p.db.Query(`select id, active, long, lat, address,
	district, number_arc, arc_type, carpet, change_date, comment,
	status, owner, operator, external_id
	from points p`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.DistibutePoint
		err := rows.Scan(&res.ID, &res.Active, &res.Long, &res.Lat, &res.Address,
			&res.District, &res.NumberArc, &res.ArcType, &res.Carpet, &res.ChangeDate, &res.Comment,
			&res.Status, &res.Owner, &res.Operator, &res.ExternalID)
		if err != nil {
			log.Println(err)
			return result, err
		}

		res.Coordinates = append(res.Coordinates, res.Long, res.Lat)
		result = append(result, res)
	}
	return result, nil
}


func (p *PostgresDB) NewTaskToPoints(data business.ApplyTask) (error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	for _, i := range data.Points {
		_, err := tx.Exec(`insert into tasks(point_id, type, comment,
		customer, done,
		deadline, entry_date) 
		values($1, $2, $3, $4, $5, $6, $7)`,
		i, data.Task, data.Comment,
		data.Customer, false,
		data.Deadline, time.Now())
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err2)
				return err2
			}
			log.Println(err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}


func (p *PostgresDB) AppointPointsToUsers(data business.Appoint) (error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	for _, i := range data.Points {
		_, err := tx.Exec(`insert into service(point_id, user_id,
		appointment_date, sent) values($1, $2, $3, $4)`,
		i, pq.Array(data.Users), time.Now(), false)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err2)
				return err2
			}
			log.Println(err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}