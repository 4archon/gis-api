package database

import (
	"log"
	"map/business"
	"strconv"
	"time"

	"github.com/lib/pq"
)

func (p *PostgresDB) NewDeclineReport(userID int, report business.DeclineReport) (int, error) {
	reportTimeNow := time.Now()
	if (len(report.Appoint) == 0) {
		report.Appoint = append(report.Appoint, 0)
	}
	targetAppoint := report.Appoint[len(report.Appoint) - 1]

	var primaryReason string = report.Reason
	if report.Reason == "Идет благоустройство - требуется забрать дуги" ||
	report.Reason == "Идет благоустройство - требуется демонтировать и забрать дуги" {
		if report.Yourself != nil && *report.Yourself {
			primaryReason = report.Reason
			report.Reason = "Идет благоустройство"
		}
	}
	
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if targetAppoint == 0 {
		row := tx.QueryRow(`insert into service values(default, $1, $2,
		$3, $4,
		$5, $6, $7, $8, $9) returning id`, report.PointID, pq.Array([]int{userID}),
		reportTimeNow, reportTimeNow,
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
			without_task = $4, sent = $5, status = $6 where id = $7`, reportTimeNow, comment, userID,
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

	_, err = tx.Exec(`insert into tasks values(default, $1, $2, $3,
	$4, $5, $6, $7, $8)`, report.PointID, "Невозможно произвести работы", nil,
	targetAppoint, "Ultradop", reportTimeNow, nil, true)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Println(err2)
			return 0, err2
		}
		log.Println(err)
		return 0, err
	}

	if primaryReason == "Идет благоустройство - требуется забрать дуги" &&
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

	if primaryReason == "Идет благоустройство - требуется демонтировать и забрать дуги" &&
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

	if primaryReason == "Точка является дублем" {
		_, err = tx.Exec(`update points set active = $1, change_date = $2,
		comment = $3, status = $4 where id = $5`, false, reportTimeNow,
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

	if primaryReason == "Невозможно установить дуги, необходимо деактивировать" {
		_, err = tx.Exec(`update points set active = $1, change_date = $2,
		status = $3 where id = $4`, false, reportTimeNow,
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

	if primaryReason != "Точка является дублем" && 
	primaryReason != "Невозможно установить дуги, необходимо деактивировать"{
		_, err = tx.Exec(`update points set change_date = $1, status = $2
		where id = $3`, reportTimeNow, report.Reason, report.PointID)
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

	if report.Reason == "Идет благоустройство - требуется забрать дуги" ||
	report.Reason == "Идет благоустройство - требуется демонтировать и забрать дуги"{
		_, err = tx.Exec(`insert into tasks values(default, $1, $2, $3,
		$4, $5, $6, $7, $8)`, report.PointID, "Благоустройство - Временный демонтаж", nil,
		nil, "Ultradop", reportTimeNow, nil, false)
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


func (p *PostgresDB) NewInspectionReport(userID int, report business.InspectionReport) (int, error) {
	reportTimeNow := time.Now()
	if (len(report.Appoint) == 0) {
		report.Appoint = append(report.Appoint, 0)
	}
	targetAppoint := report.Appoint[len(report.Appoint) - 1]
	
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if targetAppoint == 0 {
		row := tx.QueryRow(`insert into service values(default, $1, $2,
		$3, $4,
		$5, $6, $7, $8, $9) returning id`, report.PointID, pq.Array([]int{userID}),
		reportTimeNow, reportTimeNow,
		report.Comment, "Точка доступна", true, userID, true)
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
			without_task = $4, sent = $5, status = $6 where id = $7`, reportTimeNow, comment, userID,
			false, true, "Точка доступна", i)
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

	for _, i := range report.Required {
		_, err = tx.Exec(`insert into service_works values(default, $1, $2, $3, $4)`,
			targetAppoint, "required", i.WorkType, i.Count)
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

	if report.PaintCount != nil {
		_, err = tx.Exec(`insert into service_works values(default, $1, $2, $3, $4)`,
			targetAppoint, "done", "Покраска", *report.PaintCount)
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

	for _, i := range report.Tasks {
		if (i.ID == 0) {
			_, err = tx.Exec(`insert into tasks values(default, $1, $2, $3,
			$4, $5, $6, $7, $8)`, report.PointID, i.Type, nil,
			targetAppoint, "Ultradop", reportTimeNow, nil, true)
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
			_, err = tx.Exec(`update tasks set service_id = $1, done = $2 where id = $3`,
			targetAppoint, true, i.ID)
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

	if report.Required != nil && report.Required[0].WorkType != "Работа не требуется" {
		_, err = tx.Exec(`insert into tasks values(default, $1, $2, $3,
			$4, $5, $6, $7, $8)`, report.PointID, "Произвести сервис", nil,
			nil, "Ultradop", reportTimeNow, nil, false)
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


func (p *PostgresDB) NewServiceReport(userID int, report business.ServiceReport) (int, error) {
	reportTimeNow := time.Now()
	if (len(report.Appoint) == 0) {
		report.Appoint = append(report.Appoint, 0)
	}
	targetAppoint := report.Appoint[len(report.Appoint) - 1]
	
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if report.Status == nil {
		row := tx.QueryRow(`select * from service where point_id = $1 and status is not null
		order by id desc`, report.PointID)
		err := row.Scan(&report.Status)
		if err != nil {
			val := "Точка доступна"
			report.Status = &val
		}
	}

	if targetAppoint == 0 {
		row := tx.QueryRow(`insert into service values(default, $1, $2,
		$3, $4,
		$5, $6, $7, $8, $9) returning id`, report.PointID, pq.Array([]int{userID}),
		reportTimeNow, reportTimeNow,
		report.Comment, report.Status, true, userID, true)
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
			without_task = $4, sent = $5, status = $6 where id = $7`, reportTimeNow, comment, userID,
			false, true, report.Status, i)
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

	_, err = tx.Exec(`update points set status = $1, change_date = $2 where id = $3`,
	report.Status, reportTimeNow, report.PointID)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Println(err2)
			return 0, err2
		}
		log.Println(err)
		return 0, err
	}

	_, err = tx.Exec(`update points set number_arc = $1, change_date = $2 where id = $3`,
	report.NumberArc, reportTimeNow, report.PointID)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Println(err2)
			return 0, err2
		}
		log.Println(err)
		return 0, err
	}

	if report.NewLocation != nil {
		_, err = tx.Exec(`update points set long = $1, lat = $2, change_date = $3 where id = $4`,
			report.NewLocation[0], report.NewLocation[1], reportTimeNow, report.PointID)
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

	if report.NewCarpet != nil {
		_, err = tx.Exec(`update points set carpet = $1, change_date = $2 where id = $3`,
			report.NewCarpet, reportTimeNow, report.PointID)
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

	if report.Required != nil && len(report.Required) > 0 {
		for _, i := range report.Required {
			_, err = tx.Exec(`insert into service_works values(default, $1, $2, $3, $4)`,
				targetAppoint, "required", i.WorkType, i.Count)
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

		_, err = tx.Exec(`insert into tasks values(default, $1, $2, $3,
			$4, $5, $6, $7, $8)`, report.PointID, "Произвести сервис", nil,
			nil, "Ultradop", reportTimeNow, nil, false)
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

	for _, i := range report.Done {
		if i.WorkType != "Демаркировка" {
			_, err = tx.Exec(`insert into service_works values(default, $1, $2, $3, $4)`,
				targetAppoint, "done", i.WorkType, i.Count)
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
			count := len(i.MarksID)
			_, err = tx.Exec(`insert into service_works values(default, $1, $2, $3, $4)`,
				targetAppoint, "done", i.WorkType, count)
			if err != nil {
				err2 := tx.Rollback()
				if err2 != nil {
					log.Println(err2)
					return 0, err2
				}
				log.Println(err)
				return 0, err
			}

			_, err = tx.Exec(`update markings set active = $1 where id = any($2)`,
				false, pq.Array(i.MarksID))
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

		if i.WorkType == "Нанесение разметки" {
			_, err = tx.Exec(`insert into markings values(default, $1, $2, $3, $4)`,
				report.PointID, i.Number, i.MarkType, true)
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

	for _, i := range report.Tasks {
		if (i.ID == 0) {
			_, err = tx.Exec(`insert into tasks values(default, $1, $2, $3,
			$4, $5, $6, $7, $8)`, report.PointID, i.Type, nil,
			targetAppoint, "Ultradop", reportTimeNow, nil, true)
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
			_, err = tx.Exec(`update tasks set service_id = $1, done = $2 where id = $3`,
			targetAppoint, true, i.ID)
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

		if i.Type == "Замена дуги на алюминиевую" {
			_, err = tx.Exec(`update points set arc_type = $1, change_date = $2 where id = $3`,
			"Алюминиевая", reportTimeNow, report.PointID)
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

		if i.Type == "Монтаж новой точки" {
			_, err = tx.Exec(`update points set active = $1, change_date = $2 where id = $3`,
			true, reportTimeNow, report.PointID)
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

		if i.Type == "Монтаж старой точки" {
			_, err = tx.Exec(`update points set active = $1, change_date = $2 where id = $3`,
			true, reportTimeNow, report.PointID)
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
	
	if *report.Status == "Точка недоступна" {
		_, err = tx.Exec(`update points set active = $1, change_date = $2 where id = $3`,
			false, reportTimeNow, report.PointID)
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