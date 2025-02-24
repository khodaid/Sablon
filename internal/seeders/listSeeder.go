package seeders

func ListSeeder() []interface{} {
	return []interface{}{
		SeedUsers(),
		SeedSuplier(),
		SeedRole(),
		SeedUserRoleAdmin(),
		SeedStore(),
		SeedUserStore(),
	}

}
