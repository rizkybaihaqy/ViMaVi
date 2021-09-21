package models

import (
	"log"
	"vip-management-system-api/config"

	"github.com/lib/pq"
)

// Vip data type
type Vip struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	CountryOfOrigin string   `json:"country_of_origin"`
	ETA             string   `json:"eta"`
	Photo           string   `json:"photo"`
	Arrived         bool     `json:"arrived"`
	Attributes      []string `json:"attributes"`
}

func GetVips() ([]Vip, error) {
	db := config.CreateConnection()
	defer db.Close()

	var vips []Vip

	query := `SELECT id, name, country_of_origin, eta, photo, arrived, attributes FROM vips`

	rows, err := db.Query(query)
	if err != nil {
		return vips, err
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
			return vips, err
		}

		vips = append(vips, vip)
	}

	return vips, nil
}

// Get vip data by id from database
func GetVip(id int64) (Vip, error) {
	db := config.CreateConnection()
	defer db.Close()

	var vip Vip

	query := `SELECT id, name, country_of_origin, eta, photo, arrived, attributes FROM vips WHERE id = $1`
	row := db.QueryRow(query, id)

	err := row.Scan(
		&vip.ID,
		&vip.Name,
		&vip.CountryOfOrigin,
		&vip.ETA,
		&vip.Photo,
		&vip.Arrived,
		pq.Array(&vip.Attributes),
	)
	if err != nil {
		return vip, err
	}

	return vip, err
}

// Insert one vip to database
func InsertVip(vip Vip) error {
	db := config.CreateConnection()
	defer db.Close()

	query := `INSERT INTO
					vips (name, country_of_origin, eta, photo, arrived, attributes)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id`

	var id int64

	row := db.QueryRow(
		query,
		vip.Name,
		vip.CountryOfOrigin,
		vip.ETA,
		vip.Photo,
		vip.Arrived,
		pq.Array(vip.Attributes),
	)

	err := row.Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

// Update one vip from database, returning affected rows
func UpdateVip(id int64, v Vip) int64 {
	db := config.CreateConnection()
	defer db.Close()

	q := `UPDATE vips
			SET name=$2, country_of_origin=$3, eta=$4, photo=$5, arrived=$6, attributes=$7
			WHERE id=$1`

	row, err := db.Exec(
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
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rAffected, err := row.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rAffected
}
