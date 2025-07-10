package database

import (
	"log"
	"map/business"
	"time"
)

func (p *PostgresDB) GetDataForMain(id int) ([]business.Point, error) {
	var result []business.Point
	storage := make(map[int]*business.Point)

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
		var res business.Point
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

	subgroup, err := p.GetUserSubgroup(id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	if subgroup == "service" {
		rows2, err := p.db.Query(`select point_id, min(deadline) from tasks
		where done is false and type != 'Проинспектировать'
		group by point_id`)
		if err != nil {
			log.Println(err)
			return result, err
		}
		defer rows2.Close()
		for rows2.Next() {
			var pointID int
			var deadline *time.Time
			err = rows2.Scan(&pointID, &deadline)
			if err != nil {
				log.Println(err)
				return result, err
			}
			if deadline != nil {storage[pointID].Deadline = deadline}
		}
	} else if subgroup == "inspection" {
		rows2, err := p.db.Query(`select point_id, min(deadline) from tasks
		where done is false and type = 'Проинспектировать'
		group by point_id`)
		if err != nil {
			log.Println(err)
			return result, err
		}
		defer rows2.Close()
		for rows2.Next() {
			var pointID int
			var deadline *time.Time
			err = rows2.Scan(&pointID, &deadline)
			if err != nil {
				log.Println(err)
				return result, err
			}
			if deadline != nil {storage[pointID].Deadline = deadline}
		}
	} else {
		rows2, err := p.db.Query(`select point_id, min(deadline) from tasks
		where done is false
		group by point_id`)
		if err != nil {
			log.Println(err)
			return result, err
		}
		defer rows2.Close()
		for rows2.Next() {
			var pointID int
			var deadline *time.Time
			err = rows2.Scan(&pointID, &deadline)
			if err != nil {
				log.Println(err)
				return result, err
			}
			if deadline != nil {storage[pointID].Deadline = deadline}
		}
	}

	rows3, err := p.db.Query(`select point_id, id from service
	where sent is false and $1 = any(user_id)`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows3.Close()
	for rows3.Next() {
		var pointID, serviceID int
		err := rows3.Scan(&pointID, &serviceID)
		if err != nil {
			log.Println(err)
			return result, err
		}

		storage[pointID].Appoint = append(storage[pointID].Appoint, serviceID)
	}

	return result, nil
}