package models

import (
	"github.com/jmoiron/sqlx"

	log "github.com/Sirupsen/logrus"
)

type RegionRow struct {
	ID            string  `db:"id"`	
	Name          string  `db:"name"`	
	Code    int     `db:"code"`
	CityCode    int     `db:"city_code"`
	CityName    string     `db:"city_name"`
	CoverImage    string  `db:"cover_image"`
}

func (region *Region) Select(limit int, page int) ([]RegionRow, error) {

	r := []RegionRow{}
	err := region.db.Select(&r, "select t.*,c.name as city_name from town as t INNER JOIN city as c on t.city_code=c.code limit $1", limit)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}

func NewRegion(db *sqlx.DB) *Region {
	r := &Region{}
	r.db = db
	r.table = "town"
	r.hasID = true

	return r
}

type Region struct {
	Base
}
