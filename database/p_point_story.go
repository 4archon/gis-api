package database

import (
	"database/sql"
	"log"
	"map/point"
	"strconv"

	"github.com/lib/pq"
)

func (p *PostgresDB) GetPointStory(id int) ([]point.StoryPoint, error) {
	rows, err := p.db.Query(`select id, user_id,
	change_point_id, service_log_id, inspection_log_id, point_active_id,
	deadline, appointment_date, submission_date,
	sent_worker, verified
	from report
	where active = 't' and point_id = $1
	order by submission_date desc, appointment_date desc`, id)
	var story []point.StoryPoint
	if err != nil {
		log.Println(err)
		return story, err
	}
	defer rows.Close()
	for rows.Next() {
		var point point.StoryPoint
		var sqlUsers []string
		var sqlChangeID, sqlServiceID, sqlInspectionID, sqlActiveID sql.NullInt32
		var sqlDeadline, sqlAppoint, sqlSubmit sql.NullTime
		var sqlVerified sql.NullBool
		err = rows.Scan(&point.TaskID, pq.Array(&sqlUsers),
		&sqlChangeID, &sqlServiceID, &sqlInspectionID, &sqlActiveID,
		&sqlDeadline, &sqlAppoint, &sqlSubmit,
		&point.SentWorker, &sqlVerified)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if sqlChangeID.Valid {point.ChangeID = int(sqlChangeID.Int32)}
		if sqlServiceID.Valid {point.ServiceID = int(sqlServiceID.Int32)}
		if sqlInspectionID.Valid {point.InspectionID = int(sqlInspectionID.Int32)}
		if sqlActiveID.Valid {point.ActiveID = int(sqlActiveID.Int32)}
		if sqlDeadline.Valid {point.Deadline = sqlDeadline.Time}
		if sqlAppoint.Valid {point.Appointment = sqlAppoint.Time}
		if sqlSubmit.Valid {point.Submission = sqlSubmit.Time}
		if sqlVerified.Valid {
			if sqlVerified.Bool {
				point.Verified = 1
			} else {
				point.Verified = 0
			}
		} else {
			point.Verified = -1
		}

		for _, i := range sqlUsers {
			idUser, err := strconv.Atoi(i)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			point.Users = append(point.Users, idUser)
		}

		point.UsersApplied, err = p.appliedUsers(point.Users)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		story = append(story, point)
	}
	
	return story, nil
}