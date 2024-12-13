package base

import (
	"log"

	"gorm.io/gorm"
)

func CreateEnumIfNotExists(db *gorm.DB, enumName string, values []string) error {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = ?)`
	if err := db.Raw(query, enumName).Scan(&exists).Error; err != nil {
		return err
	}

	if !exists {
		enumValues := "'" + values[0] + "'"
		for _, v := range values[1:] {
			enumValues += ", '" + v + "'"
		}

		createEnumQuery := `CREATE TYPE ` + enumName + ` AS ENUM (` + enumValues + `)`

		if err := db.Exec(createEnumQuery).Error; err != nil {
			return err
		}
		log.Printf("Enum type %s created successfully\n", enumName)
	}
	return nil
}
