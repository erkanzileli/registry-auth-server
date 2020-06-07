package model

import "registry-auth/common"

func AutoMigrate() {
	common.DB.AutoMigrate(
		&Access{},
		&AccessType{},
		&Action{},
		&Repository{},
		&Tag{},
		&User{})
}
