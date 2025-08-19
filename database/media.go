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
	where point_id = $1 order by execution_date desc, m.id desc`, id)
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

func (p *PostgresDB) NewMedia(media business.Media) (int, error) {
	var mediaID int

	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return mediaID, err
	}

	row := tx.QueryRow(`insert into media values(default, $1, $2, $3) returning id`,
	media.ServiceID, media.MediaType, media.MediaName)
	err = row.Scan(&mediaID)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Println(err2)
			return 0, err2
		}
		log.Println(err)
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return mediaID, err
	}

	return mediaID, nil
}