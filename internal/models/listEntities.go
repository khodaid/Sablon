package models

func GetEntity() []interface{} {
	return []interface{}{
		&User{},
		&Supplier{},
		&Role{},
		&UserRoleAdmin{},
	}

}
