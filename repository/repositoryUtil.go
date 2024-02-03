package repository

import (
	"TestErajaya/model"
	"gorm.io/gorm"
)

func convertUserParamToDBQuery(db *gorm.DB, userParam model.ListDataRequestStruct) (dbResult *gorm.DB) {
	dbResult = db.Order(userParam.Order)

	if userParam.Page > 0 && userParam.Limit > 0 {
		offset := (userParam.Page - 1) * userParam.Limit
		dbResult = dbResult.Limit(int(userParam.Limit)).Offset(int(offset))
	}
	return
}
