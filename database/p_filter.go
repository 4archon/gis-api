package database

import (
	"database/sql"
	"log"
	"map/point"
	"time"
)

func (p *PostgresDB) GetFiltredPoints() ([]point.FilterPoint, error) {
	rows, err := p.db.Query(`select p.id, c.long, c.lat, a.point_status,
	s.execution_date, i.execution_date, idata.checkup,
	r1.max_time, r2.count
	from points p
	inner join change_points_log c on p.change_id = c.id
	inner join point_active_log a on p.active_id = a.id
	left join service_log s on p.service_id = s.id
	left join inspection_log i on p.inspection_id = i.id
	left join inspection_log_data idata on i.id = idata.inspection_log_id
	left join
	(select point_id, max(submission_date) as "max_time" from report group by point_id) r1
	on p.id = r1.point_id
	left join
	(select point_id, count(*) as "count" from report
	where verified is not true and active = 't' group by point_id) r2
	on p.id = r2.point_id
	order by id;`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var points []point.FilterPoint
	for rows.Next() {
		var point point.FilterPoint
		var inspectionTime, serviceTime, maxReportDate time.Time
		var checkup string
		var sqlInspectionTime, sqlServiceTime, sqlMaxReport sql.NullTime
		var sqlCheckup sql.NullString
		var sqlCount sql.NullInt32
		var count int
		err = rows.Scan(&point.ID, &point.Long, &point.Lat, &point.Active,
		&sqlServiceTime, &sqlInspectionTime, &sqlCheckup,
		&sqlMaxReport, &sqlCount)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if sqlServiceTime.Valid {serviceTime = sqlServiceTime.Time}
		if sqlInspectionTime.Valid {inspectionTime = sqlInspectionTime.Time}
		if sqlMaxReport.Valid {maxReportDate = sqlMaxReport.Time}
		if sqlCheckup.Valid {checkup = sqlCheckup.String}
		if sqlCount.Valid {count= int(sqlCount.Int32)}
		if inspectionTime.Compare(serviceTime) == 1 {
			if checkup != "Точка не требует ремонта" {
				point.Repair = true
			}
		}

		dur := time.Since(maxReportDate)
		var month time.Duration = time.Minute * 60 * 24 * 30 * 2
		if dur > month {
			point.LongTime = true
		}

		if count > 0 {
			point.Assigned = true
		}

		points = append(points, point)
	}

	return points, nil
}