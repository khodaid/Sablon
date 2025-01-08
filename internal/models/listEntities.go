package models

func GetEntity() []interface{} {
	return []interface{}{
		&User{},
		&Supplier{},
		&Store{},
		&Role{},
		&UserRoleAdmin{},
	}

}
