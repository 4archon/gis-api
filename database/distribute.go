package database

import (
	"log"
	"map/business"
	"time"

	"github.com/lib/pq"
)

func (p *PostgresDB) GetDataForDistribute() ([]business.DistibutePoint, error) {
	var result []business.DistibutePoint
	storage := make(map[int]*business.DistibutePoint)
	
	rows, err := p.db.Query(`select id, active, long, lat, address,
	district, number_arc, arc_type, carpet, change_date, comment,
	status, owner, operator, external_id
	from points`)
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

	for i, j := range result {
		storage[j.ID] = &result[i]
	}

	rows2, err := p.db.Query(`select s.point_id, w.id, type, work, arc
	from service s inner join service_works w on s.id = w.service_id
	inner join
	(select point_id, max(execution_date) as "max_date"
	from service s inner join service_works w on s.id = w.service_id
	where type = 'done' group by point_id) me
	on s.point_id = me.point_id
	where w.type = 'required' and
	s.execution_date >= coalesce(me.max_date, '01.01.2000')`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var res business.Work
		var pointID int
		err := rows2.Scan(&pointID, &res.ID, &res.Type, &res.Work, &res.Arc)
		if err != nil {
			log.Println(err)
			return result, err
		}

		storage[pointID].Works = append(storage[pointID].Works, res)
	}

	rows3, err := p.db.Query(`select point_id, id, type,
	comment, customer, entry_date, deadline
	from tasks where done is null or done is false`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows3.Close()
	for rows3.Next() {
		var res business.Task
		var pointID int
		err := rows3.Scan(&pointID, &res.ID, &res.Type,
			&res.Comment, &res.Customer, &res.EntryDate, &res.Deadline)
		if err != nil {
			log.Println(err)
			return result, err
		}

		storage[pointID].Tasks = append(storage[pointID].Tasks, res)
	}

	rows4, err := p.db.Query(`select point_id, id, number, type, active from markings
	where active is true`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows4.Close()
	for rows4.Next() {
		var res business.Mark
		var pointID int
		err := rows4.Scan(&pointID, &res.ID, &res.Number, &res.Type, &res.Active)
		if err != nil {
			log.Println(err)
			return result, err
		}

		storage[pointID].Marks = append(storage[pointID].Marks, res)
	}


	rows5, err := p.db.Query(`select point_id, u.id, subgroup from service s, users u
	where sent is false and u.id = any(user_id)`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows5.Close()
	for rows5.Next() {
		var res business.AppointUser
		var pointID int
		err := rows5.Scan(&pointID, &res.ID, &res.Subgroup)
		if err != nil {
			log.Println(err)
			return result, err
		}

		storage[pointID].Appoint = append(storage[pointID].Appoint, res)
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