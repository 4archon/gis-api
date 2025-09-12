package database

import (
	"log"
	"map/business"
	"time"
	"github.com/lib/pq"
)


func (p *PostgresDB) GetGSheetBase() (business.GSheetBase, error) {
	var result business.GSheetBase

	rows, err := p.db.Query(`select * from points order by id;`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.AllDataPoint
		err := rows.Scan(&res.ID, &res.Active, &res.Long, &res.Lat, &res.Address,
			&res.District, &res.NumberArc, &res.ArcType, &res.Carpet, &res.ChangeDate,
			&res.Comment, &res.Status, &res.Owner, &res.Operator, &res.ExternalID)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.Points = append(result.Points, res)
	}

	return result, nil
}


func (p *PostgresDB) GetGSheetDoneWorks(start time.Time, end time.Time) (business.GSheetDoneWorks, error) {
	var result business.GSheetDoneWorks

	rows, err := p.db.Query(`select u.login,
	p.id, p.long, p.lat, p.address, p.owner,
	s.id, user_id, execution_date, sent_by, without_task,
	w.work, w.arc
	from service s inner join service_works w on s.id = w.service_id
	inner join points p on s.point_id = p.id
	inner join users u on s.sent_by = u.id
	where w.type = 'done' and work != 'Работа не требуется' and
	execution_date >= $1 and execution_date < $2
	order by execution_date`, start, end)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.GSheetWork
		err := rows.Scan(&res.Login,
			&res.PointID, &res.Long, &res.Lat, &res.Address, &res.Owner,
			&res.ServiceID, pq.Array(&res.UserID), &res.ExecutionDate,
			&res.SentBy, &res.WithoutTask,
			&res.Work, &res.Arc)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.Works = append(result.Works, res)
	}

	return result, nil
}

func (p *PostgresDB) GetGSheetDoneVisits(start time.Time, end time.Time) (business.GSheetDoneVisits, error) {
	var result business.GSheetDoneVisits

	rows, err := p.db.Query(`select
	distinct on(point_id, execution_date, s.comment, s.status, sent_by)
	u.login,
	p.id, p.long, p.lat, p.address, p.owner,
	s.id, user_id, execution_date, sent_by, without_task
	from service s inner join points p on s.point_id = p.id
	inner join users u on s.sent_by = u.id
	left join (select service_id as "id" from service_works) as "wt"
	on s.id = wt.id
	where execution_date >= $1 and execution_date < $2
	order by execution_date, point_id, sent_by, s.status, s.comment,
	wt.id`, start, end)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.GSheetVisit
		err := rows.Scan(&res.Login,
			&res.PointID, &res.Long, &res.Lat, &res.Address, &res.Owner,
			&res.ServiceID, pq.Array(&res.UserID), &res.ExecutionDate,
			&res.SentBy, &res.WithoutTask)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.Visits = append(result.Visits, res)
	}

	return result, nil
}