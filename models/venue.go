package models

import (

	"github.com/jmoiron/sqlx"

	log "github.com/Sirupsen/logrus"
)


type VenueRow struct {
	ID      string `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
}

func (venue *Venue) GetOne() (*VenueRow,error){
	 	v := &VenueRow{}
	 	err := venue.db.Get(v, "SELECT name,address FROM venue limit 1")
	 	if err != nil {
	 		log.Println(err)
	 		return nil,err
	 	}
	 	return v,nil
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
