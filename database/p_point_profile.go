package database

import (
	"database/sql"
	"log"
	"map/point"
)

func (p *PostgresDB) GetPointProfile(id int) (point.PointProfile, error) {
	row := p.db.QueryRow(`select a.point_status, a.change_date,
	c.long, c.lat, c.point_address, c.district,
	c.number_arc, c.arc_type, c.carpet, c.change_date,
	s.execution_date, i.execution_date, s.id, i.id
	from points p inner join point_active_log a on p.active_id = a.id
	inner join change_points_log c on p.change_id = c.id
	inner join service_log s on p.service_id = s.id
	inner join inspection_log i on p.inspection_id = i.id
	where p.id = $1;`, id)
	var profile point.PointProfile
	profile.ID = id
	var sqlAddress, sqlDistrict, sqlArcType, sqlCarpet sql.NullString
	var sqlNumArc sql.NullInt32
	var sqlActiveDate, sqlChangeDate, sqlServiceDate, sqlInspectionDate sql.NullTime
	var serviceID, inspectionID int
	err := row.Scan(&profile.Status, &sqlActiveDate,
	&profile.Long, &profile.Lat, &sqlAddress, &sqlDistrict,
	&sqlNumArc, &sqlArcType, &sqlCarpet, &sqlChangeDate,
	&sqlServiceDate, &sqlInspectionDate, &serviceID, &inspectionID)
	if err != nil {
		log.Println(err)
		return profile, err
	}
	if sqlActiveDate.Valid {profile.StatusLastChange = sqlActiveDate.Time}
	if sqlAddress.Valid {profile.Address = sqlAddress.String}
	if sqlDistrict.Valid {profile.District = sqlDistrict.String}
	if sqlNumArc.Valid {profile.NumberArc = int(sqlNumArc.Int32)}
	if sqlArcType.Valid {profile.ArcType = sqlArcType.String}
	if sqlCarpet.Valid {profile.Carpet = sqlCarpet.String}
	if sqlChangeDate.Valid {profile.PointLastChange = sqlChangeDate.Time}
	if sqlServiceDate.Valid {profile.ServiceLast = sqlServiceDate.Time}
	if sqlInspectionDate.Valid {profile.InspectionLast = sqlInspectionDate.Time}

	profile.Service, err = p.GetService(serviceID)
	if err != nil {
		return profile, err
	}
	profile.Inspection, err = p.GetInspection(inspectionID)
	if err != nil {
		return profile, err
	}

	return profile, nil
}