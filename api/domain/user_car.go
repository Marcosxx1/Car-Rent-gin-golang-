package domain

import "gorm.io/gorm"

type UserCar struct {
	gorm.Model
	CARID  string
	USERID string
}
