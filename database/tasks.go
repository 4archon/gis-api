package database

import (
	"map/business"
	"log"
)

func (p *PostgresDB) GetPointCurrentTasks(id int) (business.Tasks, error) {
	var result business.Tasks
	result.PointID = id

	rows, err := p.db.Query(`select id, type, comment
	from tasks where point_id = $1`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.Task
		err := rows.Scan(&res.ID, &res.Type, &res.Comment)
		if err != nil {
			log.Println(err)
			return result, err
		}
		result.Tasks = append(result.Tasks, res)
	}

	rows2, err := p.db.Query(`select type, work, arc
	from service s inner join service_works w on s.id = w.service_id
	where s.point_id = $1 and w.type = 'required' and 
	s.execution_date > (select max(execution_date)
	from service s inner join service_works w on s.id = w.service_id
	where type = 'done' and point_id = $1)`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var res business.Work
		err := rows2.Scan(&res.ID, &res.Type, &res.Work, &res.Arc)
		if err != nil {
			log.Println(err)
			return result, err
		}
		result.Works = append(result.Works, res)
	}

	return result, nil
}