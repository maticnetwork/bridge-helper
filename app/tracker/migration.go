package tracker

import (
	"log"

	"gorm.io/gorm"
)

// Running automatic database migration, on application start up
func migrateDB(db *gorm.DB) {
	if err := db.AutoMigrate(&RootChain{}, &ChildChain{}); err != nil {
		log.Fatalln("[!] ", err)
	}
}
