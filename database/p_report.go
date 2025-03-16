package database

import (
	"database/sql"
	"log"
	"strconv"
	"time"
	"map/point"
	// "strconv"
	// "github.com/lib/pq"
)

func (p *PostgresDB) GetPointIDFromReport(id int) (int, error) {
	row := p.db.QueryRow(`select point_id from report where id = $1`, id)
	var pointID int
	err := row.Scan(&pointID)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return pointID, nil
}


func (p *PostgresDB) CreateInspection(reportID int, checkup string,
	repairType string, comment string) (int, error) {
	pointID, err := p.GetPointIDFromReport(reportID)
	if err != nil {
		return -1, err
	}
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return -1, err
	}

	row := tx.QueryRow(`insert into inspection_log(point_id, execution_date) values($1, $2) returning id`,
				pointID, time.Now())
	var insLogID int
	err = row.Scan(&insLogID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return -1, err
	}


	row = tx.QueryRow(`insert into inspection_log_data(inspection_log_id, checkup, repair_type, comment)
	values($1, $2, $3, $4) returning id`,
	insLogID, checkup, repairType, comment)
	var insLogDataID int
	err = row.Scan(&insLogDataID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return -1, err
	}

	dest := "/media/inspection/" + strconv.Itoa(insLogDataID) + "/"
	_, err = tx.Exec(`update inspection_log_data set photo_left = $1, photo_right = $2,
	photo_front = $3, video = $4 where id = $5`, dest + "photo_left.jpeg",
	dest + "photo_right.jpeg", dest + "photo_front.jpeg", dest + "video.mov",
	insLogDataID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return -1, err
	}

	_, err = tx.Exec(`update report set inspection_log_id = $1 where id = $2`,
	insLogID, reportID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return insLogDataID, err
}

func (p *PostgresDB) GetInspection(inspectionID int) (point.InspectionReport, error) {
	row := p.db.QueryRow(`select inspection_log_id, checkup, repair_type,
	photo_left, photo_right, photo_front, video, comment
	from inspection_log_data where inspection_log_id = $1;`, inspectionID)
	var insReport point.InspectionReport
	var sqlComment sql.NullString
	err := row.Scan(&insReport.ID, &insReport.Checkup, &insReport.RepairType,
		&insReport.PhotoLeft, &insReport.PhotoRight, &insReport.PhotoFront,
		&insReport.Video, &sqlComment)
	if err != nil {
		log.Println(err)
		return insReport, err
	}
	if sqlComment.Valid {insReport.Comment = sqlComment.String}
	
	return insReport, nil
}

func (p *PostgresDB) DeleteInspection(reportID int) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(`update report set inspection_log_id = null where id = $1`, reportID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	
	return nil
}


func (p *PostgresDB) CreateService(reportID int) (int, error) {
	pointID, err := p.GetPointIDFromReport(reportID)
	if err != nil {
		return -1, err
	}
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return -1, err
	}

	row := tx.QueryRow(`insert into service_log(point_id, execution_date) values($1, $2) returning id`,
				pointID, time.Now())
	var serviceLogID int
	err = row.Scan(&serviceLogID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return serviceLogID, err
}


func (p *PostgresDB) CreateServiceLogData(serviceLogID int, serviceType string,
	subtype string, comment string) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return -1, err
	}

	row := tx.QueryRow(`insert into service_log_data(service_log_id, service_type, subtype,
	comment, action_arc) values($1, $2, $3, $4, $5) returning id`,
	serviceLogID, serviceType, subtype, comment, 1)
	var serviceLogDataID int
	err = row.Scan(&serviceLogDataID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return serviceLogDataID, err
}


func (p *PostgresDB) ApproveServiceReport(reportID int, serviceLogID int,
	serviceLogDataID int, extra bool) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	dest := "/media/service/" + strconv.Itoa(serviceLogDataID) + "/"
	if extra {
		_, err = tx.Exec(`update service_log_data set photo_before = $1, photo_left = $2,
		photo_right = $3, photo_front = $4, video = $5, photo_extra = $6 where id = $7`,
		dest + "photo_before.jpeg", dest + "photo_left.jpeg", dest + "photo_right.jpeg",
		dest + "photo_front.jpeg", dest + "video.mov", dest + "photo_extra.jpeg",
		serviceLogDataID)
	} else {
		_, err = tx.Exec(`update service_log_data set photo_before = $1, photo_left = $2,
		photo_right = $3, photo_front = $4, video = $5 where id = $6`,
		dest + "photo_before.jpeg", dest + "photo_left.jpeg", dest + "photo_right.jpeg",
		dest + "photo_front.jpeg", dest + "video.mov",
		serviceLogDataID)
	}
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`update report set service_log_id = $1 where id = $2`,
	serviceLogID, reportID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return  err
	}

	return nil
}


func (p *PostgresDB) GetService(serviceID int) ([]point.ServiceReport, error) {
	rows, err := p.db.Query(`select service_log_id, service_type, subtype, action_arc,
	photo_before, photo_left, photo_right, photo_front, video, photo_extra, comment
	from service_log_data where service_log_id = $1;`, serviceID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var serviceReports []point.ServiceReport
	counter := 0
	for rows.Next() {
		var service point.ServiceReport
		counter++
		service.Index = counter
		var sqlComment, sqlPhotoExtra sql.NullString
		err := rows.Scan(&service.ID, &service.ServiceType, &service.Subtype, &service.ActionArc,
		&service.PhotoBefore, &service.PhotoLeft, &service.PhotoRight, &service.PhotoFront,
		&service.Video, &sqlPhotoExtra, &sqlComment)
		if err != nil {
			log.Println(err)
			return serviceReports, err
		}
		if sqlComment.Valid {service.Comment = sqlComment.String}
		if sqlPhotoExtra.Valid {service.PhotoExtra = sqlPhotoExtra.String}

		serviceReports = append(serviceReports, service)
	}
	
	return serviceReports, nil
}


func (p *PostgresDB) DeleteService(reportID int) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(`update report set service_log_id = null where id = $1`, reportID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	
	return nil
}


func (p *PostgresDB) GetPointInfo(idLog int) (point.ChangeReport, error) {
	row := p.db.QueryRow(`select long, lat, point_address, district,
	number_arc, arc_type, carpet, change_date, comment
	from change_points_log where id = $1`, idLog)
	var change point.ChangeReport
	var sqlNumArc sql.NullInt32
	var sqlChangeDate sql.NullTime
	var sqlAddress, sqlDistrict, sqlArcType, sqlCarpet, sqlComment sql.NullString
	err := row.Scan(&change.Long, &change.Lat, &sqlAddress, &sqlDistrict,
	&sqlNumArc, &sqlArcType, &sqlCarpet, &sqlChangeDate, &sqlComment)
	if err != nil {
		log.Println(err)
		return change, err
	}
	if sqlAddress.Valid {change.PointAddress = sqlAddress.String}
	if sqlDistrict.Valid {change.District = sqlDistrict.String}
	if sqlNumArc.Valid {change.NumberArc = int(sqlNumArc.Int32)}
	if sqlArcType.Valid {change.ArcType = sqlArcType.String}
	if sqlCarpet.Valid {change.Carpet = sqlCarpet.String}
	if sqlChangeDate.Valid {change.ChangeDate = sqlChangeDate.Time}
	if sqlComment.Valid {change.Comment = sqlComment.String}

	return change, err
}


func (p *PostgresDB) GetPointFromReport(reportID int) (point.ChangeReport, error) {
	var change point.ChangeReport
	pointID, err := p.GetPointIDFromReport(reportID)
	if err != nil {
		log.Println(err)
		return change, err
	}

	row := p.db.QueryRow(`select change_id from points where id = $1`, pointID)
	var changePointID int
	err = row.Scan(&changePointID)
	if err != nil {
		log.Println(err)
		return change, err
	}

	change, err = p.GetPointInfo(changePointID)
	if err != nil {
		return change, err
	}
	
	return change, nil
}


func (p *PostgresDB) NewChangePoint(reportID int, change point.ChangeReport) error {
	pointID, err := p.GetPointIDFromReport(reportID)
	if err != nil {
		log.Println(err)
		return err
	}
	
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	row := tx.QueryRow(`insert into change_points_log values(default,
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10) returning id`,
			pointID, change.Long, change.Lat, change.PointAddress, change.District,
			change.NumberArc, change.ArcType, change.Carpet, change.ChangeDate, change.Comment)
	
	var changePointID int
	err = row.Scan(&changePointID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	_, err = tx.Exec(`update report set change_point_id = $1 where id = $2`,
	changePointID, reportID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	
	return nil
}


func (p *PostgresDB) DeleteChangePoint(reportID int) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(`update report set change_point_id = null where id = $1`, reportID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	
	return nil
}