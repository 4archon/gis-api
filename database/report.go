package database

import (
	"log"
	"map/business"
	"strconv"
	"time"

	"github.com/lib/pq"
)

func (p *PostgresDB) NewDeclineReport(userID int, report business.DeclineReport) (int, error) {
	if (len(report.Appoint) == 0) {
		report.Appoint = append(report.Appoint, 0)
	}
	targetAppoint := report.Appoint[len(report.Appoint) - 1]

	var primaryReson string = report.Reason
	if report.Reason == "Идет благоустройство - требуется забрать дуги" ||
	report.Reason == "Идет благоустройство - требуется демонтировать и забрать дуги" {
		if report.Yourself != nil && *report.Yourself {
			primaryReson = report.Reason
			report.Reason = "Идет благоустройство"
		}
	}
	
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if targetAppoint == 0 {
		row := tx.QueryRow(`insert into service values(default, $1, $2, $3, $4,
		$5, $6, $7, $8, $9) returning id`, report.PointID, pq.Array(userID), time.Now(), time.Now(),
		report.Comment, report.Reason, true, userID, true)
		err := row.Scan(&targetAppoint)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err2)
				return 0, err2
			}
			log.Println(err)
			return 0, err
		}
	} else {
		for _, i := range report.Appoint {
			var comment *string = nil
			if i == targetAppoint {
				comment = report.Comment
			}
			_, err := tx.Exec(`update service set execution_date = $1, comment = $2, sent_by = $3,
			without_task = $4, sent = $5, status = $6 where id = $7`, time.Now(), comment, userID,
			false, true, report.Reason, i)
			if err != nil {
				err2 := tx.Rollback()
				if err2 != nil {
					log.Println(err2)
					return 0, err2
				}
				log.Println(err)
				return 0, err
			}
		}
	}

	if primaryReson == "Идет благоустройство - требуется забрать дуги" &&
	report.Yourself != nil && *report.Yourself {
		var numberArc int
		row := tx.QueryRow(`select number_arc from points where id = $1`, report.PointID)
		err := row.Scan(&numberArc)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err2)
				return 0, err2
			}
			log.Println(err)
			return 0, err
		}
		_, err = tx.Exec(`insert into service_works values(default, $1, $2, $3, $4)`,
		targetAppoint, "done", "Вывоз дуг", numberArc)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err2)
				return 0, err2
			}
			log.Println(err)
			return 0, err
		}
	}

	if primaryReson == "Идет благоустройство - требуется демонтировать и забрать дуги" &&
	report.Yourself != nil && *report.Yourself {
		var numberArc int
		row := tx.QueryRow(`select number_arc from points where id = $1`, report.PointID)
		err := row.Scan(&numberArc)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err2)
				return 0, err2
			}
			log.Println(err)
			return 0, err
		}
		_, err = tx.Exec(`insert into service_works values(default, $1, $2, $3, $4)`,
		targetAppoint, "done", "Демонтаж", numberArc)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err2)
				return 0, err2
			}
			log.Println(err)
			return 0, err
		}
		_, err = tx.Exec(`insert into service_works values(default, $1, $2, $3, $4)`,
		targetAppoint, "done", "Вывоз дуг", numberArc)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err2)
				return 0, err2
			}
			log.Println(err)
			return 0, err
		}
	}

	if primaryReson == "Точка является дублем" {
		_, err = tx.Exec(`update points set active = $1, change_date = $2,
		comment = $3, status = $4 where id = $5`, false, time.Now(),
		"Точка явлется дублем точки - " + strconv.Itoa(report.Duplicate.Original),
		"Точка недоступна", report.PointID)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err2)
				return 0, err2
			}
			log.Println(err)
			return 0, err
		}
	}

	if primaryReson == "Невозможно установить дуги, необходимо деактивировать" {
		_, err = tx.Exec(`update points set active = $1, change_date = $2,
		status = $3 where id = $4`, false, time.Now(),
		"Точка недоступна", report.PointID)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Println(err2)
				return 0, err2
			}
			log.Println(err)
			return 0, err
		}
	}


	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return targetAppoint, nil
}