package database

import (
	"database/sql"
	"log"
	"map/point"
	"strconv"
	"time"

	"github.com/lib/pq"
)

func (p *PostgresDB) AssignTasks(data point.TasksRequest) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	for _, i := range data.Points {
		if data.Deadline == "" {
			_, err = tx.Exec(`insert into report(point_id, user_id, appointment_date, 
			sent_worker, verified, active, deadline)
			values($1, $2, $3,
			$4, $5, $6, $7)`, i, pq.Array(data.Employees), time.Now(),
			false, nil, true, nil)
		} else {
			_, err = tx.Exec(`insert into report(point_id, user_id, appointment_date, 
			sent_worker, verified, active, deadline)
			values($1, $2, $3,
			$4, $5, $6, $7)`, i, pq.Array(data.Employees), time.Now(),
			false, nil, true, data.Deadline)
		}
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				log.Println(err)
				return err
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

func (p *PostgresDB) appliedUsers(users []int) ([]point.AppliedUser, error) {
	var appliedUsers []point.AppliedUser
	rows, err := p.db.Query(`select id, login from users where id = any ($1)`, pq.Array(users))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user point.AppliedUser
		var sqlLogin sql.NullString
		err = rows.Scan(&user.ID, &sqlLogin)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if sqlLogin.Valid {user.Login = sqlLogin.String}
		appliedUsers = append(appliedUsers, user)
	}
	return appliedUsers, nil
}

func (p *PostgresDB) GetTasksInfo() ([]point.Task, error) {
	rows, err := p.db.Query(`select r.id, r.user_id, p.id,
	r.change_point_id, r.service_log_id, r.inspection_log_id, r.point_active_id,
	c.point_address, c.lat, c.long, c.district, c.number_arc, c.arc_type, c.carpet,
	r.deadline, r.sent_worker, r.verified
	from report r inner join points p on r.point_id = p.id
	inner join change_points_log c on p.change_id = c.id
	where r.verified is not true and r.active = 't'
	order by r.sent_worker desc, r.appointment_date;`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var tasks []point.Task
	for rows.Next() {
		var task point.Task
		var sqlUsers []string
		var sqlChangeID, sqlServiceID, sqlInspectionID, sqlActiveID sql.NullInt32
		var sqlAddress, sqlDistrict, sqlArcType, sqlCarpet sql.NullString
		var sqlNumArc sql.NullInt32
		var sqlDeadline sql.NullTime
		var sqlVerified sql.NullBool
		err = rows.Scan(&task.TaskID, pq.Array(&sqlUsers), &task.PointID,
		&sqlChangeID, &sqlServiceID, &sqlInspectionID, &sqlActiveID,
		&sqlAddress, &task.Lat, &task.Long, &sqlDistrict, &sqlNumArc, &sqlArcType, &sqlCarpet,
		&sqlDeadline, &task.SentWorker, &sqlVerified)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if sqlChangeID.Valid {task.ChangeID = int(sqlChangeID.Int32)}
		if sqlServiceID.Valid {task.ServiceID = int(sqlServiceID.Int32)}
		if sqlInspectionID.Valid {task.InspectionID = int(sqlInspectionID.Int32)}
		if sqlActiveID.Valid {task.ActiveID = int(sqlActiveID.Int32)}
		if sqlAddress.Valid {task.Address = sqlAddress.String}
		if sqlDistrict.Valid {task.District = sqlDistrict.String}
		if sqlArcType.Valid {task.TypeArc = sqlArcType.String}
		if sqlCarpet.Valid {task.Carpet = sqlCarpet.String}
		if sqlNumArc.Valid {task.NumberArc = int(sqlNumArc.Int32)}
		if sqlDeadline.Valid {task.Deadline = sqlDeadline.Time}
		if sqlVerified.Valid {
			if sqlVerified.Bool {
				task.Verified = 1
			} else {
				task.Verified = 0
			}
		} else {
			task.Verified = -1
		}

		for _, i := range sqlUsers {
			id, err := strconv.Atoi(i)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			task.Users = append(task.Users, id)
		}

		task.UsersApplied, err = p.appliedUsers(task.Users)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	
	return tasks, nil
}

func (p *PostgresDB) GetUserTasksInfo(userID int) ([]point.Task, error) {
	rows, err := p.db.Query(`select r.id, r.user_id, p.id,
	r.change_point_id, r.service_log_id, r.inspection_log_id, r.point_active_id,
	c.point_address, c.lat, c.long, c.district, c.number_arc, c.arc_type, c.carpet,
	r.deadline, r.sent_worker, r.verified
	from report r inner join points p on r.point_id = p.id
	inner join change_points_log c on p.change_id = c.id
	where r.verified is not true and r.active = 't' and $1 = any(r.user_id)
	order by r.sent_worker, r.appointment_date;`, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var tasks []point.Task
	for rows.Next() {
		var task point.Task
		var sqlUsers []string
		var sqlChangeID, sqlServiceID, sqlInspectionID, sqlActiveID sql.NullInt32
		var sqlAddress, sqlDistrict, sqlArcType, sqlCarpet sql.NullString
		var sqlNumArc sql.NullInt32
		var sqlDeadline sql.NullTime
		var sqlVerified sql.NullBool
		err = rows.Scan(&task.TaskID, pq.Array(&sqlUsers), &task.PointID,
		&sqlChangeID, &sqlServiceID, &sqlInspectionID, &sqlActiveID,
		&sqlAddress, &task.Lat, &task.Long, &sqlDistrict, &sqlNumArc, &sqlArcType, &sqlCarpet,
		&sqlDeadline, &task.SentWorker, &sqlVerified)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if sqlChangeID.Valid {task.ChangeID = int(sqlChangeID.Int32)}
		if sqlServiceID.Valid {task.ServiceID = int(sqlServiceID.Int32)}
		if sqlInspectionID.Valid {task.InspectionID = int(sqlInspectionID.Int32)}
		if sqlActiveID.Valid {task.ActiveID = int(sqlActiveID.Int32)}
		if sqlAddress.Valid {task.Address = sqlAddress.String}
		if sqlDistrict.Valid {task.District = sqlDistrict.String}
		if sqlArcType.Valid {task.TypeArc = sqlArcType.String}
		if sqlCarpet.Valid {task.Carpet = sqlCarpet.String}
		if sqlNumArc.Valid {task.NumberArc = int(sqlNumArc.Int32)}
		if sqlDeadline.Valid {task.Deadline = sqlDeadline.Time}
		if sqlVerified.Valid {
			if sqlVerified.Bool {
				task.Verified = 1
			} else {
				task.Verified = 0
			}
		} else {
			task.Verified = -1
		}

		for _, i := range sqlUsers {
			id, err := strconv.Atoi(i)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			task.Users = append(task.Users, id)
		}

		task.UsersApplied, err = p.appliedUsers(task.Users)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	
	return tasks, nil
}


func (p *PostgresDB) DeleteDeactivation(reportID int) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(`update report set point_active_id = null where id = $1`, reportID)
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


func (p *PostgresDB) SendReport(reportID int) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(`update report set sent_worker = 't', verified = 'f',
	submission_date = $2
	where id = $1`, reportID, time.Now())
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

func (p *PostgresDB) DeclineReport(reportID int) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(`update report set sent_worker = 'f', submission_date = $2
	where id = $1`, reportID, nil)
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

func (p *PostgresDB) VerifyReport(reportID int) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(`update report set verified = 't'
	where id = $1`, reportID)
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