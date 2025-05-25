package database

import (
	"map/business"
	"log"
)

func (p *PostgresDB) GetPointMedia(id int) (business.PointMedias, error) {
	var result business.PointMedias
	result.ID = id

	rows, err := p.db.Query(`select m.id, media_type, media_name 
	from service s inner join media m on s.id = m.service_id
	where point_id = $1 order by execution_date desc`, id)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var res business.Media
		err := rows.Scan(&res.ID, &res.MediaType, &res.MediaName)
		if err != nil {
			log.Println(err)
			return result, err
		}
		result.Medias = append(result.Medias, res)
	}

	return result, nil
}