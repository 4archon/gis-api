package database

import (
	"map/business"
	"log"
)

func (p *PostgresDB) GetDataForMain(id int) ([]business.MainPoint, error) {
	var result []business.MainPoint
	
	rows, err := p.db.Query(`select p.id, active, long, lat, address,
	district, number_arc, arc_type, carpet, change_date, p.comment,
	s.appointed
	from points p left join
	(select point_id, 't' as "appointed" from service where sent = 'f' and $1 = any(user_id)) s 
	on p.id = s.point_id`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.MainPoint
		err := rows.Scan(&res.ID, &res.Active, &res.Long, &res.Lat, &res.Address,
			&res.District, &res.NumberArc, &res.ArcType, &res.Carpet, &res.ChangeDate, &res.Comment,
			&res.Appointed)
		if err != nil {
			log.Println(err)
			return result, err
		}

		res.Coordinates = append(res.Coordinates, res.Long, res.Lat)
		result = append(result, res)
	}
	return result, nil
}