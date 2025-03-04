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
		_, err = tx.Exec(`insert into report(point_id, user_id, appointment_date, 
		sent_worker, verified, active) values($1, $2, $3,
		$4, $5, $6)`, i, pq.Array(data.Employees), time.Now(),
		false, false, false)
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

func (p *PostgresDB) stringifyUsers(users []int) (string, error) {
	res := ""
	rows, err := p.db.Query(`select id, login from users where id = any ($1)`, pq.Array(users))
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		var sqlLogin sql.NullString
		var id int
		err = rows.Scan(&id, &sqlLogin)
		if err != nil {
			log.Println(err)
			return "", err
		}
		res += strconv.Itoa(id)
		if sqlLogin.Valid {res += ":" +sqlLogin.String}
		res += ","
	}
	return res, nil
}

func (p *PostgresDB) GetTasksInfo() ([]point.Task, error) {
	rows, err := p.db.Query(`select r.id, r.user_id, p.id,
	r.change_point_id, r.service_log_id, r.inspection_log_id,
	c.point_address, c.lat, c.long, c.district, c.number_arc, c.arc_type, c.carpet
	from report r inner join points p on r.point_id = p.id
	inner join change_points_log c on p.change_id = c.id
	where r.verified = 'f';`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var tasks []point.Task
	for rows.Next() {
		var task point.Task
		var sqlUsers []string
		var sqlChangeID, sqlServiceID, sqlInspectionID sql.NullInt32
		var sqlAddress, sqlDistrict, sqlArcType, sqlCarpet sql.NullString
		var sqlNumArc sql.NullInt32
		err = rows.Scan(&task.TaskID, pq.Array(&sqlUsers), &task.PointID,
		&sqlChangeID, &sqlServiceID, &sqlInspectionID,
		&sqlAddress, &task.Lat, &task.Long, &sqlDistrict, &sqlNumArc, &sqlArcType, &sqlCarpet)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if sqlChangeID.Valid {task.ChangeID = int(sqlChangeID.Int32)}
		if sqlServiceID.Valid {task.ServiceID = int(sqlServiceID.Int32)}
		if sqlInspectionID.Valid {task.InspectionID = int(sqlInspectionID.Int32)}
		if sqlAddress.Valid {task.Address = sqlAddress.String}
		if sqlDistrict.Valid {task.District = sqlDistrict.String}
		if sqlArcType.Valid {task.TypeArc = sqlArcType.String}
		if sqlCarpet.Valid {task.Carpet = sqlCarpet.String}
		if sqlNumArc.Valid {task.NumberArc = int(sqlNumArc.Int32)}

		for _, i := range sqlUsers {
			id, err := strconv.Atoi(i)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			task.Users = append(task.Users, id)
		}

		task.UsersStr, err = p.stringifyUsers(task.Users)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	
	return tasks, nil
}