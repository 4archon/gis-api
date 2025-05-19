package database

import (
	"map/business"
	"log"
)

func (p *PostgresDB) GetPointsForAnalytics() ([]business.AnalyticsPoint, error) {
	var result []business.AnalyticsPoint
	
	rows, err := p.db.Query(`select id, active, long, lat, address,
	district, number_arc, arc_type, carpet, change_date, comment
	from points`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.AnalyticsPoint
		err := rows.Scan(&res.ID, &res.Active, &res.Long, &res.Lat, &res.Address,
			&res.District, &res.NumberArc, &res.ArcType, &res.Carpet, &res.ChangeDate, &res.Comment)
		if err != nil {
			log.Println(err)
			return result, err
		}

		res.Coordinates = append(res.Coordinates, res.Long, res.Lat)
		result = append(result, res)
	}
	return result, nil
}