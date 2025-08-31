package database

import (
	"errors"
	"log"
	"map/business"

	"github.com/lib/pq"
)

func (p *PostgresDB) GetPointHistory(id int) (business.History, error) {
	var result business.History
	result.ID = id
	storage := make(map[int]*business.StoryPoint)
	storage_order := []int{}
	userLogins := make(map[string]string)

	rows3, err := p.db.Query(`select id, login from users`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows3.Close()
	for rows3.Next() {
		var userID string
		var login *string
		err := rows3.Scan(&userID, &login)
		if err != nil {
			log.Println(err)
			return result, err
		}
		if login == nil {
			userLogins[userID] = ""
		} else {
			userLogins[userID] = *login
		}
	}

	rows, err := p.db.Query(`select id, user_id, execution_date,
	comment, status, sent 
	from service where point_id = $1 order by execution_date desc`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.StoryPoint
		var storyID int
		err := rows.Scan(&storyID, pq.Array(&res.UserIDs), &res.Execution,
			&res.Comment, &res.Status, &res.Sent)
		if err != nil {
			log.Println(err)
			return result, err
		}
		res.ID = storyID
		result.StoryPoints = append(result.StoryPoints, res)
		storage_order = append(storage_order, storyID)
	}

	for j, i := range storage_order {
		storage[i] = &result.StoryPoints[j]
	}

	rows2, err := p.db.Query(`select s.id, type, work, arc 
	from service s inner join service_works w on s.id = w.service_id
	where s.point_id = $1`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var res business.Work
		var storyID int
		err := rows2.Scan(&storyID, &res.Type, &res.Work, &res.Arc)
		if err != nil {
			log.Println(err)
			return result, err
		}
		storage[storyID].Works = append(storage[storyID].Works, res)
	}

	rows4, err := p.db.Query(`select s.id, t.id, type, t.comment,
	t.customer, t.entry_date, t.deadline, t.done 
	from service s inner join tasks t on s.id = t.service_id
	where s.point_id = $1`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows4.Close()
	for rows4.Next() {
		var res business.Task
		var storyID int
		err := rows4.Scan(&storyID, &res.ID, &res.Type, &res.Comment,
		&res.Customer, &res.EntryDate, &res.Deadline, &res.Done)
		if err != nil {
			log.Println(err)
			return result, err
		}
		storage[storyID].Tasks = append(storage[storyID].Tasks, res)
	}

	rows5, err := p.db.Query(`select s.id, m.id, media_type, media_name
	from service s inner join media m on s.id = m.service_id
	where point_id = $1`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows5.Close()
	for rows5.Next() {
		var res business.Media
		var storyID int
		err := rows5.Scan(&storyID, &res.ID, &res.MediaType, &res.MediaName)
		if err != nil {
			log.Println(err)
			return result, err
		}
		storage[storyID].Medias = append(storage[storyID].Medias, res)
	}


	for i := range result.StoryPoints{
		for _, j := range result.StoryPoints[i].UserIDs{
			result.StoryPoints[i].UserLogins = append(result.StoryPoints[i].UserLogins, userLogins[j])
		}
	}

	return result, nil
}

func (p *PostgresDB) GetAllServices(numRows int, offset int) (business.AllServices, error) {
	var result business.AllServices
	if offset < 0 {
		return result, errors.New("offset is negative") 
	}
	storage := make(map[int]*business.ServiceStoryPoint)

	rows, err := p.db.Query(`select id,
	appointment_date, execution_date, comment, status, sent, without_task
	from service where sent is true order by execution_date desc, id desc
	limit $1 offset $2`, numRows, offset * numRows)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.ServiceStoryPoint
		err := rows.Scan(&res.ID, &res.Appoint, &res.Execution, &res.Comment,
			&res.Status, &res.Sent, &res.WithoutTasks)
		if err != nil {
			log.Println(err)
			return result, err
		}
		result.Services = append(result.Services, res)
	}

	for j, i := range result.Services {
		storage[i.ID] = &result.Services[j]
	}

	rows2, err := p.db.Query(`select s.id, s.sent_by, u.id, login, role, active,
	name, surname, patronymic, tg_id, subgroup, trust
	from service s, users u where u.id = any(user_id)`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var serviceID int
		var sentBy *int
		var res business.UserInfo
		err := rows2.Scan(&serviceID, &sentBy, &res.ID, &res.Login, &res.Role,
			&res.Active, &res.Name, &res.Surname, &res.Patronymic, &res.TgID,
			&res.Subgroup, &res.Trust)
		if err != nil {
			log.Println(err)
			return result, err
		}

		_, ok := storage[serviceID]; if ok {
			storage[serviceID].Users = append(storage[serviceID].Users, res)
		}

		if sentBy != nil && res.ID == *sentBy {
			_, ok := storage[serviceID]; if ok {
				storage[serviceID].SentBy = &res
			}
		}
	}

	rows3, err := p.db.Query(`select s.id, w.id, type, work, arc 
	from service s inner join service_works w on s.id = w.service_id`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows3.Close()
	for rows3.Next() {
		var res business.Work
		var serviceID int
		err := rows3.Scan(&serviceID, &res.ID, &res.Type, &res.Work, &res.Arc)
		if err != nil {
			log.Println(err)
			return result, err
		}
		_, ok := storage[serviceID]; if ok {
			storage[serviceID].Works = append(storage[serviceID].Works, res)
		}
	}

	rows4, err := p.db.Query(`select s.id, t.id, type, t.comment,
	t.customer, t.entry_date, t.deadline, t.done
	from service s inner join tasks t on s.id = t.service_id`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows4.Close()
	for rows4.Next() {
		var res business.Task
		var serviceID int
		err := rows4.Scan(&serviceID, &res.ID, &res.Type, &res.Comment,
		&res.Customer, &res.EntryDate, &res.Deadline, &res.Done)
		if err != nil {
			log.Println(err)
			return result, err
		}
		_, ok := storage[serviceID]; if ok {
			storage[serviceID].Tasks = append(storage[serviceID].Tasks, res)
		}
	}

	rows5, err := p.db.Query(`select service_id, id, media_type, media_name from media`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows5.Close()
	for rows5.Next() {
		var res business.Media
		var serviceID int
		err := rows5.Scan(&serviceID, &res.ID, &res.MediaType, &res.MediaName)
		if err != nil {
			log.Println(err)
			return result, err
		}
		_, ok := storage[serviceID]; if ok {
			storage[serviceID].Medias = append(storage[serviceID].Medias, res)
		}
	}

	rows6, err := p.db.Query(`select s.id, p.*
	from points p inner join service s on p.id = s.point_id;`)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows6.Close()
	for rows6.Next() {
		var res business.PurePoint
		var serviceID int
		err := rows6.Scan(&serviceID, &res.ID, &res.Active, &res.Long, &res.Lat,
		&res.Address, &res.District, &res.NumberArc, &res.ArcType, &res.Carpet,
		&res.ChangeDate, &res.Comment, &res.Status, &res.Owner, &res.Operator,
		&res.ExternalID)
		if err != nil {
			log.Println(err)
			return result, err
		}
		_, ok := storage[serviceID]; if ok {
			storage[serviceID].Point = res
		}
	}

	row7 := p.db.QueryRow(`select count(*) from service where sent is true`)
	var serviceCount int
	err = row7.Scan(&serviceCount)
	if err != nil {
		log.Println(err)
		return result, err
	}

	result.CurrentPage = offset + 1
	result.LastPage = serviceCount / numRows
	if serviceCount % numRows != 0 {
		result.LastPage += 1
	}

	return result, nil
}