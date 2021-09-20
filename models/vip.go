package models

import (
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
