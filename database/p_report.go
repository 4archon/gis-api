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