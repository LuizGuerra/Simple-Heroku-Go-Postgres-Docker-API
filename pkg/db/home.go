package db

import (
	"github.com/go-pg/pg/v10"
)

type Home struct {
	ID      int64  `json:"id"` // json values must match sql columns
	Price   int64  `json:"price"`
	AgentId int64  `json:"agent_id"`
	Agent   *Agent `pg:"rel:has-one" json:"agent"`
}

func GetHomes(pgDb *pg.DB) ([]*Home, error) {
	homes := make([]*Home, 0)
	err := pgDb.Model(&homes).
		Relation("Agent"). // se n tem relacao n precisa usar isso
		Select()

	return homes, err
}

func GetHome(pgDb *pg.DB, homeID string) (*Home, error) {
	queryResult := &Home{}
	err := pgDb.Model(queryResult).
		Relation("Agent"). // same as line 8
		Where("home.id = ?", homeID).
		Select()
	return queryResult, err
}

func CreateHome(pgDb *pg.DB, req *Home) (*Home, error) {
	_, err := pgDb.Model(req).Insert()
	if err != nil {
		return nil, err
	}
	queryResult := &Home{}
	err = pgDb.Model(queryResult).
		Relation("Agent"). // same as line 12
		Where("home.id = ?", req.ID).
		Select()
	return queryResult, err
}

func UpdateHome(pgDb *pg.DB, req *Home) (*Home, error) {
	_, err := pgDb.Model(req).WherePK().Update()
	if err != nil {
		return nil, err
	}
	queryResult := &Home{}
	err = pgDb.Model(queryResult).
		Relation("Agent"). // same as line 12
		Where("home.id = ?", req.ID).
		Select()

	return queryResult, err
}

func DeleteHome(pgDb *pg.DB, homeID int64) error {
	_, err := pgDb.Model(&Home{ID: homeID}).WherePK().Delete()
	return err
}

// with error if passed `id` doesn't exist
func DeleteHomeWEINE(pgDb *pg.DB, homeID int64) error {
	home := &Home{ID: homeID}
	err := pgDb.Model(home).
		Relation("Agent"). // same as line 12
		Where("home.id = ?", homeID).
		Select()
	if err != nil {
		return err
	}
	_, err = pgDb.Model(&Home{ID: homeID}).WherePK().Delete()
	return err
}
