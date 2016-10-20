package models

import (
	"github.com/jmoiron/sqlx"

	log "github.com/Sirupsen/logrus"
)

type VenueRow struct {
	ID            string  `db:"id"`
	No            int     `db:"no"`
	Name          string  `db:"name"`
	Address       string  `db:"address"`
	Tel           string  `db:"tel"`
	Lat           float64 `db:"lat"`
	Long          float64 `db:"long"`
	City          string  `db:"city"`
	Town          string  `db:"town"`
	HasOwnWifi    bool    `db:"has_own_wifi"`
	HasWiSpotter  bool    `db:"has_wispotter"`
	HasTtnetWifi  bool    `db:"has_ttnet_wifi"`
	WifiRating    int     `db:"wifi_rating"`
	WifiAvgSpeed  float64 `db:"wifi_avg_speed"`
	PlugRating    int     `db:"plug_rating"`
	NoiseRating   int     `db:"noise_rating"`
	SeatCapacity  int     `db:"seat_capacity"`
	TableCapacity int     `db:"table_capacity"`
}

func (venue *Venue) Select(limit int,page int) ([]VenueRow, error) {

	v := []VenueRow{}
	err := venue.db.Select(&v, "SELECT * FROM venue limit $1",limit)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return v, nil
}

func NewVenue(db *sqlx.DB) *Venue {
	v := &Venue{}
	v.db = db
	v.table = "venue"
	v.hasID = true

	return v
}

type Venue struct {
	Base
}
