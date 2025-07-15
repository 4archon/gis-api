package database

import (
	"map/business"
	"log"
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