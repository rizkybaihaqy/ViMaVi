package models

import (
	"database/sql"

	"github.com/lib/pq"
)

// Vip model type
type Vip struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	CountryOfOrigin string   `json:"country_of_origin"`
	ETA             string   `json:"eta"`
	Photo           string   `json:"photo"`
	Arrived         bool     `json:"arrived"`
	Attributes      []string `json:"attributes"`
}

type VipModel struct {
	DB *sql.DB
}

// VipModel struct constructor
func NewVipModel(db *sql.DB) *VipModel {
	return &VipModel{DB: db}
}

// Get all vip
func (m VipModel) GetVips() ([]Vip, error) {
	var v []Vip

	q := `SELECT id, name, country_of_origin, eta, photo, arrived, attributes FROM vips`

	rows, err := m.DB.Query(q)
	if err != nil {
		return v, err
	}

	defer rows.Close()

	for rows.Next() {
		var vip Vip

		err = rows.Scan(
			&vip.ID,
			&vip.Name,
			&vip.CountryOfOrigin,
			&vip.ETA,
			&vip.Photo,
			&vip.Arrived,
			pq.Array(&vip.Attributes),
		)
		if err != nil {
			return v, err
		}

		v = append(v, vip)
	}

	return v, nil
}

// Get vip data by id from database
func (m VipModel) GetVip(id int64) (Vip, error) {
	var v Vip

	query := `SELECT id, name, country_of_origin, eta, photo, arrived, attributes FROM vips WHERE id = $1`
	row := m.DB.QueryRow(query, id)

	err := row.Scan(
		&v.ID,
		&v.Name,
		&v.CountryOfOrigin,
		&v.ETA,
		&v.Photo,
		&v.Arrived,
		pq.Array(&v.Attributes),
	)
	if err != nil {
		return v, err
	}

	return v, err
}

// Insert one vip to database
func (m VipModel) InsertVip(v Vip) error {
	query := `INSERT INTO
					vips (name, country_of_origin, eta, photo, arrived, attributes)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id`

	var id int64

	row := m.DB.QueryRow(
		query,
		v.Name,
		v.CountryOfOrigin,
		v.ETA,
		v.Photo,
		v.Arrived,
		pq.Array(v.Attributes),
	)

	err := row.Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

// Update one vip from database, returning affected rows
func (m VipModel) UpdateVip(id int64, v Vip) (int64, error) {
	q := `UPDATE vips
			SET name=$2, country_of_origin=$3, eta=$4, photo=$5, arrived=$6, attributes=$7
			WHERE id=$1`

	row, err := m.DB.Exec(
		q,
		id,
		v.Name,
		v.CountryOfOrigin,
		v.ETA,
		v.Photo,
		v.Arrived,
		pq.Array(v.Attributes),
	)
	if err != nil {
		return 0, err
	}

	rAffected, err := row.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rAffected, nil
}

// Delete one from database
func (m VipModel) DeleteVip(id int64) (int64, error) {
	sqlStatement := `DELETE FROM vips WHERE id=$1`

	res, err := m.DB.Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}

	rAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rAffected, nil
}

// Update one vip arrived status to true
func (m VipModel) ArrivedVip(id int64) (int64, error) {
	q := `UPDATE vips SET arrived=$2 WHERE id=$1`

	row, err := m.DB.Exec(q, id, true)
	if err != nil {
		return 0, err
	}

	rAffected, err := row.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rAffected, nil
}
