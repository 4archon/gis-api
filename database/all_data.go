package database

import (
	"map/business"
	"log"
	"github.com/lib/pq"
)


func (p *PostgresDB) GetAllData() (business.AllData, error) {
	var result business.AllData

	rows1, err := p.db.Query(`select * from users order by id;`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows1.Close()
	for rows1.Next() {
		var res business.User
		err := rows1.Scan(&res.ID, &res.Login, &res.Password, &res.Role, &res.Active,
			&res.Name, &res.Surname, &res.Patronymic, &res.TgID, &res.Subgroup, &res.Trust)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.Users = append(result.Users, res)
	}

	rows2, err := p.db.Query(`select * from points order by id;`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var res business.AllDataPoint
		err := rows2.Scan(&res.ID, &res.Active, &res.Long, &res.Lat, &res.Address,
			&res.District, &res.NumberArc, &res.ArcType, &res.Carpet, &res.ChangeDate,
			&res.Comment, &res.Status, &res.Owner, &res.Operator, &res.ExternalID)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.Points = append(result.Points, res)
	}

	rows3, err := p.db.Query(`select * from points_log order by id;`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows3.Close()
	for rows3.Next() {
		var res business.AllDataLogPoint
		err := rows3.Scan(&res.ID, &res.PointID, &res.Active, &res.Long, &res.Lat, &res.Address,
			&res.District, &res.NumberArc, &res.ArcType, &res.Carpet, &res.ChangeDate,
			&res.Comment, &res.Status, &res.Owner, &res.Operator, &res.ExternalID)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.PointsLog = append(result.PointsLog, res)
	}

	rows4, err := p.db.Query(`select * from markings order by id;`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows4.Close()
	for rows4.Next() {
		var res business.AllDataMark
		err := rows4.Scan(&res.ID, &res.PointID, &res.Number, &res.Type, &res.Active)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.Marks = append(result.Marks, res)
	}

	rows5, err := p.db.Query(`select * from service order by id;`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows5.Close()
	for rows5.Next() {
		var res business.AllDataService
		err := rows5.Scan(&res.ID, &res.PointID, pq.Array(&res.UserID), &res.AppointmentDate,
		&res.ExecutionDate, &res.Comment, &res.Status, &res.Sent, &res.SentBy, &res.WithoutTask)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.Service = append(result.Service, res)
	}

	rows6, err := p.db.Query(`select * from service_works order by id;`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows6.Close()
	for rows6.Next() {
		var res business.AllDataWork
		err := rows6.Scan(&res.ID, &res.ServiceID, &res.Type, &res.Work, &res.Arc)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.Works = append(result.Works, res)
	}

	rows7, err := p.db.Query(`select * from media order by id;`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows7.Close()
	for rows7.Next() {
		var res business.Media
		err := rows7.Scan(&res.ID, &res.ServiceID, &res.MediaType, &res.MediaName)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.Media = append(result.Media, res)
	}

	rows8, err := p.db.Query(`select * from tasks order by id;`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows8.Close()
	for rows8.Next() {
		var res business.AllDataTask
		err := rows8.Scan(&res.ID, &res.PointID, &res.Type, &res.Comment, &res.ServiceID,
		&res.Customer, &res.EntryDate, &res.Deadline, &res.Done)
		if err != nil {
			log.Println(err)
			return result, err
		}

		result.Tasks = append(result.Tasks, res)
	}

	return result, nil
}