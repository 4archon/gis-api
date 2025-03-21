package database

import (
	"map/point"
	"log"
	"database/sql"
	"github.com/lib/pq"
)

func (p *PostgresDB) GetAllPoints() []point.Point {
	rows, err := p.db.Query(`select p.id, lat, long from points p inner join change_points_log l
	on p.change_id = l.id;`)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	var points []point.Point
	for rows.Next() {
		var point point.Point
		err = rows.Scan(&point.ID, &point.Lat, &point.Long)
		if err != nil {
			log.Println(err)
			return nil
		}
		points = append(points, point)
	}
	return points
}

func (p *PostgresDB) GetAllTaskPoints() []point.Point {
	rows, err := p.db.Query(`select p.id, lat, long
	from points p inner join change_points_log l
	on p.change_id = l.id inner join report r on p.id = r.point_id
	where r.verified is not true and r.active = 't';`)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	var points []point.Point
	for rows.Next() {
		var point point.Point
		err = rows.Scan(&point.ID, &point.Lat, &point.Long)
		if err != nil {
			log.Println(err)
			return nil
		}
		points = append(points, point)
	}
	return points
}

func (p *PostgresDB) GetUserTaskPoints(userID int) []point.Point {
	rows, err := p.db.Query(`select p.id, lat, long
	from points p inner join change_points_log l
	on p.change_id = l.id inner join report r on p.id = r.point_id
	where r.verified is not true and r.active = 't'
	and $1 = any(r.user_id);`, userID)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	var points []point.Point
	for rows.Next() {
		var point point.Point
		err = rows.Scan(&point.ID, &point.Lat, &point.Long)
		if err != nil {
			log.Println(err)
			return nil
		}
		points = append(points, point)
	}
	return points
}

func (p *PostgresDB) GetPointsDesc(pointsID []int) []point.PointDesc {
	targetID := pq.Array(pointsID)
	rows, err := p.db.Query(`select p.id, point_address, change_date, number_arc
	from points p inner join change_points_log l on p.change_id = l.id
	where p.id = any ($1);`, targetID)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	var pointsDesc []point.PointDesc
	for rows.Next() {
		var desc point.PointDesc
		var timeSql sql.NullTime
		desc.Img = ""
		err = rows.Scan(&desc.ID, &desc.Address, &timeSql, &desc.Amount)
		if err != nil {
			log.Println(err)
			return nil
		}
		if timeSql.Valid {
			desc.Date = timeSql.Time
		}
		pointsDesc = append(pointsDesc, desc)
	}
	return pointsDesc
}