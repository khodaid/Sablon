package models

func GetListCreateEnum() [][]interface{} {
	return [][]interface{}{
		{
			"user_role_enum",
			UserRoles,
		},
	}
}
