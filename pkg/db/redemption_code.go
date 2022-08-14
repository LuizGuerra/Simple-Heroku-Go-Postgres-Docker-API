package db

import "github.com/go-pg/pg/v10"

type RedemptionCode struct {
	Code  string `json:"code"`
	Coins int16  `json:"coins"`
	Valid bool   `json:"bool"`
}

func GetCode(db *pg.DB, codeId string) (*RedemptionCode, error) {
	redemptionCode := &RedemptionCode{}
	err := db.Model(redemptionCode).
		Where("redemptionCode.code = ?", codeId).
		Select()

	return redemptionCode, err
}
